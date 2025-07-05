package swift_test

import (
    "testing"

    "github.com/bazelbuild/bazel-gazelle/config"
    "github.com/bazelbuild/bazel-gazelle/rule"
    "github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/swift"
    "github.com/stretchr/testify/assert"
)

func TestImports(t *testing.T) {
    tests := []struct {
        name     string
        rules    []*rule.Rule
        expected []any
    }{
        {
            name:     "empty rules",
            rules:    []*rule.Rule{},
            expected: []any{},
        },
        {
            name: "single rule with import",
            rules: []*rule.Rule{
                func() *rule.Rule {
                    r := rule.NewRule("swift_library", "test")
                    r.SetPrivateAttr(config.GazelleImportsKey, "Foundation")
                    return r
                }(),
            },
            expected: []any{"Foundation"},
        },
        {
            name: "multiple rules with imports",
            rules: []*rule.Rule{
                func() *rule.Rule {
                    r := rule.NewRule("swift_library", "test1")
                    r.SetPrivateAttr(config.GazelleImportsKey, "Foundation")
                    return r
                }(),
                func() *rule.Rule {
                    r := rule.NewRule("swift_library", "test2")
                    r.SetPrivateAttr(config.GazelleImportsKey, "Combine")
                    return r
                }(),
            },
            expected: []any{"Foundation", "Combine"},
        },
        {
            name: "rule with nil import",
            rules: []*rule.Rule{
                func() *rule.Rule {
                    r := rule.NewRule("swift_library", "test")
                    // No import set, so it should be nil
                    return r
                }(),
            },
            expected: []any{nil},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            actual := swift.Imports(tt.rules)
            assert.Equal(t, tt.expected, actual)
        })
    }
}
