package swift_test

import (
    "testing"

    "github.com/bazelbuild/bazel-gazelle/language"
    "github.com/bazelbuild/bazel-gazelle/rule"
    "github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/swift"
    "github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/swiftpkg"
    "github.com/stretchr/testify/assert"
)

func TestRulesFromSrcs(t *testing.T) {
    t.Run("empty sources", func(t *testing.T) {
        args := language.GenerateArgs{
            Dir: "/tmp/test",
            File: &rule.File{
                Path: "/tmp/test/BUILD.bazel",
                Pkg:  "test",
            },
        }

        rules := swift.RulesFromSrcs(
            args,
            []string{},
            "Empty",
            "Empty",
            []string{},
        )

        // Even with empty sources, it creates a library rule
        assert.Len(t, rules, 1)
        assert.Equal(t, "swift_library", rules[0].Kind())
        assert.Equal(t, "Empty", rules[0].Name())
    })

    // Note: Testing with actual Swift files requires filesystem operations
    // which makes unit testing difficult. The core logic is tested through integration tests.
    t.Run("integration with actual files", func(t *testing.T) {
        t.Skip("Requires filesystem operations - tested through integration tests")
    })
}

func TestCollectSwiftInfo(t *testing.T) {
    tests := []struct {
        name            string
        fileInfos       []*swiftpkg.SwiftFileInfo
        expectedImports []string
        expectedModType swift.ModuleType
    }{
        {
            name: "library module",
            fileInfos: []*swiftpkg.SwiftFileInfo{
                {
                    Rel:          "lib.swift",
                    Imports:      []string{"Foundation"},
                    IsTest:       false,
                    ContainsMain: false,
                },
            },
            expectedImports: []string{"Foundation"},
            expectedModType: swift.LibraryModuleType,
        },
        {
            name: "binary module with main",
            fileInfos: []*swiftpkg.SwiftFileInfo{
                {
                    Rel:          "main.swift",
                    Imports:      []string{"Foundation"},
                    IsTest:       false,
                    ContainsMain: true,
                },
            },
            expectedImports: []string{"Foundation"},
            expectedModType: swift.BinaryModuleType,
        },
        {
            name: "test module",
            fileInfos: []*swiftpkg.SwiftFileInfo{
                {
                    Rel:          "test.swift",
                    Imports:      []string{"XCTest"},
                    IsTest:       true,
                    ContainsMain: false,
                },
            },
            expectedImports: []string{"XCTest"},
            expectedModType: swift.TestModuleType,
        },
        {
            name: "GUI module with main",
            fileInfos: []*swiftpkg.SwiftFileInfo{
                {
                    Rel:          "app.swift",
                    Imports:      []string{"SwiftUI"},
                    IsTest:       false,
                    ContainsMain: true,
                },
            },
            expectedImports: []string{"SwiftUI"},
            expectedModType: swift.LibraryModuleType, // GUI modules stay as library
        },
        {
            name: "multiple files with duplicate imports",
            fileInfos: []*swiftpkg.SwiftFileInfo{
                {
                    Rel:          "file1.swift",
                    Imports:      []string{"Foundation", "Combine"},
                    IsTest:       false,
                    ContainsMain: false,
                },
                {
                    Rel:          "file2.swift",
                    Imports:      []string{"Foundation", "CoreData"},
                    IsTest:       false,
                    ContainsMain: false,
                },
            },
            expectedImports: []string{"Combine", "CoreData", "Foundation"},
            expectedModType: swift.LibraryModuleType,
        },
        {
            name:            "empty file infos",
            fileInfos:       []*swiftpkg.SwiftFileInfo{},
            expectedImports: []string{},
            expectedModType: swift.LibraryModuleType,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            imports, moduleType := swift.CollectSwiftInfo(tt.fileInfos)
            assert.Equal(t, tt.expectedImports, imports)
            assert.Equal(t, tt.expectedModType, moduleType)
        })
    }
}
