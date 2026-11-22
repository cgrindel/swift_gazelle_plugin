package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/swift"
	"github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/swiftpkg"
)

const dependencyIndexPerms = 0644

// swiftDepsInfo is the manifest rules_swift_package_manager's swift_deps_info
// target writes: each direct dependency's identity and the path to its
// pkg_info.json.
type swiftDepsInfo struct {
	DirectDepPkgInfos map[string]string `json:"direct_dep_pkg_infos"`
}

func run(args []string) error {
	if len(args) == 0 || args[0] != "create" {
		return errors.New("usage: swift_deps_index create [-o <output>] " +
			"[-swift_deps_info <manifest>] [<identity>=<pkg_info.json>...]")
	}
	return create(args[1:], os.Stdout)
}

func create(args []string, stdout io.Writer) error {
	fs := flag.NewFlagSet("create", flag.ContinueOnError)
	outputPath := fs.String("o", "", "The path where to write the Swift index file.")
	fs.StringVar(outputPath, "output", "", "The path where to write the Swift index file.")
	manifestPath := fs.String("swift_deps_info", "",
		"A swift_deps_info manifest from rules_swift_package_manager.")
	if err := fs.Parse(args); err != nil {
		return err
	}

	pkgInfoPaths := make(map[string]string)
	if *manifestPath != "" {
		if err := readManifest(*manifestPath, pkgInfoPaths); err != nil {
			return err
		}
	}
	for _, pair := range fs.Args() {
		identity, pkgInfoPath, err := parseIdentityPkgInfoPathPair(pair)
		if err != nil {
			return err
		}
		pkgInfoPaths[identity] = pkgInfoPath
	}

	diJSON, err := createIndex(pkgInfoPaths)
	if err != nil {
		return err
	}
	if *outputPath != "" {
		return os.WriteFile(*outputPath, diJSON, dependencyIndexPerms)
	}
	_, err = fmt.Fprintln(stdout, string(diJSON))
	return err
}

func readManifest(path string, pkgInfoPaths map[string]string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var info swiftDepsInfo
	if err := json.Unmarshal(data, &info); err != nil {
		return fmt.Errorf("failed to read swift_deps_info manifest %s: %w", path, err)
	}
	for identity, pkgInfoPath := range info.DirectDepPkgInfos {
		pkgInfoPaths[identity] = pkgInfoPath
	}
	return nil
}

func createIndex(pkgInfoPaths map[string]string) ([]byte, error) {
	identities := make([]string, 0, len(pkgInfoPaths))
	for identity := range pkgInfoPaths {
		identities = append(identities, identity)
	}
	sort.Strings(identities)

	di := swift.NewDependencyIndex()
	for _, identity := range identities {
		pkgInfoJSON, err := os.ReadFile(pkgInfoPaths[identity])
		if err != nil {
			return nil, err
		}
		pkgInfo, err := swiftpkg.NewPackageInfoFromJSON(pkgInfoJSON)
		if err != nil {
			return nil, err
		}
		if err := di.IndexBazelRepo(swift.NewBazelRepo(identity, pkgInfo, nil)); err != nil {
			return nil, err
		}
		di.AddDirectDependency(identity)
	}
	return di.JSON()
}

func parseIdentityPkgInfoPathPair(pair string) (string, string, error) {
	identity, pkgInfoPath, ok := strings.Cut(pair, "=")
	if !ok {
		return "", "", fmt.Errorf("invalid identity-package info path pair (%s)", pair)
	}
	return identity, pkgInfoPath, nil
}
