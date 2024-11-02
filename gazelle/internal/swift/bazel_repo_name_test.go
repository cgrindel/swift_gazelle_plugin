package swift_test

import (
	"testing"

	"github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/swift"
	"github.com/stretchr/testify/assert"
)

func TestRepoNameFromIdentity(t *testing.T) {
	actual := swift.RepoNameFromIdentity("swift-argument-parser")
	assert.Equal(t, "swiftpkg_swift_argument_parser", actual)
}
