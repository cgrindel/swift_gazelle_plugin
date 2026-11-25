package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/swift"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func writeFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0755))
	require.NoError(t, os.WriteFile(path, []byte(content), 0644))
	return path
}

func readIndex(t *testing.T, path string) *swift.DependencyIndex {
	t.Helper()
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	di, err := swift.NewDependencyIndexFromJSON(data)
	require.NoError(t, err)
	return di
}

func TestCreate(t *testing.T) {
	dir := t.TempDir()
	fooInfo := writeFile(t, dir, "foo/pkg_info.json", `{"name": "Foo", "path": "/tmp/foo", "targets": [], "products": [], "dependencies": []}`)
	barInfo := writeFile(t, dir, "bar/pkg_info.json", `{"name": "Bar", "path": "/tmp/bar", "targets": [], "products": [], "dependencies": []}`)

	t.Run("identity pairs", func(t *testing.T) {
		out := filepath.Join(dir, "pairs.json")
		require.NoError(t, run([]string{"create", "-o", out, "foo=" + fooInfo, "bar=" + barInfo}))
		assert.ElementsMatch(t, []string{"foo", "bar"}, readIndex(t, out).DirectDepIdentities())
	})

	t.Run("swift_deps_info manifest", func(t *testing.T) {
		manifest := writeFile(t, dir, "swift_deps_info.json",
			`{"direct_dep_pkg_infos": {"foo": "`+fooInfo+`", "bar": "`+barInfo+`"}}`)
		out := filepath.Join(dir, "manifest.json")
		require.NoError(t, run([]string{"create", "-swift_deps_info", manifest, "-o", out}))
		assert.ElementsMatch(t, []string{"foo", "bar"}, readIndex(t, out).DirectDepIdentities())
	})

	t.Run("no packages", func(t *testing.T) {
		out := filepath.Join(dir, "empty.json")
		require.NoError(t, run([]string{"create", "-o", out}))
		assert.Empty(t, readIndex(t, out).DirectDepIdentities())
	})

	t.Run("bad pair", func(t *testing.T) {
		assert.Error(t, run([]string{"create", "foo"}))
	})

	t.Run("bad manifest", func(t *testing.T) {
		manifest := writeFile(t, dir, "bad.json", `not json`)
		assert.Error(t, run([]string{"create", "-swift_deps_info", manifest}))
	})

	t.Run("no command", func(t *testing.T) {
		assert.Error(t, run(nil))
	})
}
