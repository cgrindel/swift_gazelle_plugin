package swiftpkg_test

import (
	"testing"

	"github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/swiftpkg"
	"github.com/stretchr/testify/assert"
)

func TestNewModuleType(t *testing.T) {
	tests := []struct {
		name string
		str  string
		exp  swiftpkg.ModuleType
	}{
		{name: "swift", str: "SwiftTarget", exp: swiftpkg.SwiftModuleType},
		{name: "clang", str: "ClangTarget", exp: swiftpkg.ClangModuleType},
		{name: "unknown", str: "unknown", exp: swiftpkg.UnknownModuleType},
		{name: "unrecognized", str: "unrecognized", exp: swiftpkg.UnknownModuleType},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := swiftpkg.NewModuleType(test.str)
			assert.Equal(t, test.exp, actual)
		})
	}
}
