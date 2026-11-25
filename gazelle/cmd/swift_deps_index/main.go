// Command swift_deps_index writes the Swift dependency index that the Swift
// Gazelle plugin reads (-swift_dependency_index), from the pkg_info.json files
// that rules_swift_package_manager generates for each Swift package.
//
//	swift_deps_index create [-o <output>] [-swift_deps_info <manifest>] [<identity>=<pkg_info.json>...]
package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
