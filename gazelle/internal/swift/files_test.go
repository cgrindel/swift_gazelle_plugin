package swift_test

import (
	"testing"

	"github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/swift"
	"github.com/stretchr/testify/assert"
)

func TestFilterFiles(t *testing.T) {
	values := []string{"Foo.swift", "README.md", "Bar/Hello.swift", "Package.swift"}

	actual := swift.FilterFiles(values)
	expected := []string{"Foo.swift", "Bar/Hello.swift"}
	assert.Equal(t, expected, actual)
}
