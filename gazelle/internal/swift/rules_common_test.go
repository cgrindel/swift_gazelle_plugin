package swift_test

import (
	"testing"

	"github.com/bazelbuild/bazel-gazelle/rule"
	"github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/swift"
	"github.com/stretchr/testify/assert"
)

func TestRulesForLibraryModule(t *testing.T) {
	tests := []struct {
		name              string
		defaultName       string
		defaultModuleName string
		srcs              []string
		swiftImports      []string
		shouldSetVis      bool
		swiftLibraryTags  []string
		expectedRuleKind  string
		expectedName      string
	}{
		{
			name:              "basic library module",
			defaultName:       "MyLibrary",
			defaultModuleName: "MyLibrary",
			srcs:              []string{"lib.swift"},
			swiftImports:      []string{"Foundation"},
			shouldSetVis:      true,
			swiftLibraryTags:  []string{"tag1"},
			expectedRuleKind:  "swift_library",
			expectedName:      "MyLibrary",
		},
		{
			name:              "library with no tags",
			defaultName:       "SimpleLib",
			defaultModuleName: "SimpleLib",
			srcs:              []string{"simple.swift"},
			swiftImports:      []string{},
			shouldSetVis:      false,
			swiftLibraryTags:  []string{},
			expectedRuleKind:  "swift_library",
			expectedName:      "SimpleLib",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildFile := &rule.File{
				Path: "test/BUILD.bazel",
				Pkg:  "test",
			}

			rules := swift.RulesForLibraryModule(
				tt.defaultName,
				tt.defaultModuleName,
				tt.srcs,
				tt.swiftImports,
				tt.shouldSetVis,
				tt.swiftLibraryTags,
				buildFile,
			)

			assert.Len(t, rules, 1)
			assert.Equal(t, tt.expectedRuleKind, rules[0].Kind())
			assert.Equal(t, tt.expectedName, rules[0].Name())
			assert.Equal(t, tt.srcs, rules[0].AttrStrings("srcs"))
		})
	}
}

func TestRulesForBinaryModule(t *testing.T) {
	tests := []struct {
		name              string
		defaultName       string
		defaultModuleName string
		srcs              []string
		swiftImports      []string
		shouldSetVis      bool
		expectedRuleKind  string
		expectedName      string
	}{
		{
			name:              "basic binary module",
			defaultName:       "MyBinary",
			defaultModuleName: "MyBinary",
			srcs:              []string{"main.swift"},
			swiftImports:      []string{"Foundation"},
			shouldSetVis:      true,
			expectedRuleKind:  "swift_binary",
			expectedName:      "MyBinary",
		},
		{
			name:              "binary with no visibility",
			defaultName:       "SimpleBin",
			defaultModuleName: "SimpleBin",
			srcs:              []string{"app.swift"},
			swiftImports:      []string{},
			shouldSetVis:      false,
			expectedRuleKind:  "swift_binary",
			expectedName:      "SimpleBin",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildFile := &rule.File{
				Path: "test/BUILD.bazel",
				Pkg:  "test",
			}

			rules := swift.RulesForBinaryModule(
				tt.defaultName,
				tt.defaultModuleName,
				tt.srcs,
				tt.swiftImports,
				tt.shouldSetVis,
				buildFile,
			)

			assert.Len(t, rules, 1)
			assert.Equal(t, tt.expectedRuleKind, rules[0].Kind())
			assert.Equal(t, tt.expectedName, rules[0].Name())
			assert.Equal(t, tt.srcs, rules[0].AttrStrings("srcs"))
		})
	}
}

func TestRulesForTestModule(t *testing.T) {
	tests := []struct {
		name              string
		defaultName       string
		defaultModuleName string
		srcs              []string
		swiftImports      []string
		shouldSetVis      bool
		expectedName      string
	}{
		{
			name:              "basic test module",
			defaultName:       "MyTest",
			defaultModuleName: "MyTest",
			srcs:              []string{"test.swift"},
			swiftImports:      []string{"XCTest"},
			shouldSetVis:      true,
			expectedName:      "MyTest",
		},
		{
			name:              "test with no visibility",
			defaultName:       "SimpleTest",
			defaultModuleName: "SimpleTest",
			srcs:              []string{"simple_test.swift"},
			swiftImports:      []string{},
			shouldSetVis:      false,
			expectedName:      "SimpleTest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildFile := &rule.File{
				Path: "test/BUILD.bazel",
				Pkg:  "test",
			}

			rules := swift.RulesForTestModule(
				tt.defaultName,
				tt.defaultModuleName,
				tt.srcs,
				tt.swiftImports,
				tt.shouldSetVis,
				buildFile,
			)

			assert.Len(t, rules, 1)
			assert.Equal(t, tt.expectedName, rules[0].Name())
			assert.Equal(t, tt.srcs, rules[0].AttrStrings("srcs"))
		})
	}
}
