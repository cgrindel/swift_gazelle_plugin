package swiftpkg_test

import (
	"encoding/json"
	"testing"

	"github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/swiftpkg"
	"github.com/stretchr/testify/assert"
)

func TestTargetType_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name         string
		jsonValue    string
		expectedType swiftpkg.TargetType
	}{
		{
			name:         "executable type",
			jsonValue:    `"executable"`,
			expectedType: swiftpkg.ExecutableTargetType,
		},
		{
			name:         "test type",
			jsonValue:    `"test"`,
			expectedType: swiftpkg.TestTargetType,
		},
		{
			name:         "library type",
			jsonValue:    `"library"`,
			expectedType: swiftpkg.LibraryTargetType,
		},
		{
			name:         "regular type (alias for library)",
			jsonValue:    `"regular"`,
			expectedType: swiftpkg.LibraryTargetType,
		},
		{
			name:         "plugin type",
			jsonValue:    `"plugin"`,
			expectedType: swiftpkg.PluginTargetType,
		},
		{
			name:         "unknown type",
			jsonValue:    `"unknown"`,
			expectedType: swiftpkg.UnknownTargetType,
		},
		{
			name:         "invalid type",
			jsonValue:    `"invalid"`,
			expectedType: swiftpkg.UnknownTargetType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var targetType swiftpkg.TargetType
			err := json.Unmarshal([]byte(tt.jsonValue), &targetType)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedType, targetType)
		})
	}

	t.Run("invalid json", func(t *testing.T) {
		var targetType swiftpkg.TargetType
		err := json.Unmarshal([]byte(`123`), &targetType)
		assert.Error(t, err)
	})
}

func TestTargets_FindByName(t *testing.T) {
	targets := swiftpkg.Targets{
		&swiftpkg.Target{Name: "Target1", Path: "/path1"},
		&swiftpkg.Target{Name: "Target2", Path: "/path2"},
		&swiftpkg.Target{Name: "Target3", Path: "/path3"},
	}

	tests := []struct {
		name         string
		searchName   string
		expectedPath string
		shouldFind   bool
	}{
		{
			name:         "find existing target",
			searchName:   "Target2",
			expectedPath: "/path2",
			shouldFind:   true,
		},
		{
			name:         "target not found",
			searchName:   "NonExistent",
			expectedPath: "",
			shouldFind:   false,
		},
		{
			name:         "empty name",
			searchName:   "",
			expectedPath: "",
			shouldFind:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := targets.FindByName(tt.searchName)
			if tt.shouldFind {
				assert.NotNil(t, result)
				assert.Equal(t, tt.searchName, result.Name)
				assert.Equal(t, tt.expectedPath, result.Path)
			} else {
				assert.Nil(t, result)
			}
		})
	}
}

func TestTargets_FindByPath(t *testing.T) {
	targets := swiftpkg.Targets{
		&swiftpkg.Target{Name: "Target1", Path: "/path1"},
		&swiftpkg.Target{Name: "Target2", Path: "/path2"},
		&swiftpkg.Target{Name: "Target3", Path: "/path3"},
	}

	tests := []struct {
		name         string
		searchPath   string
		expectedName string
		shouldFind   bool
	}{
		{
			name:         "find existing target by path",
			searchPath:   "/path2",
			expectedName: "Target2",
			shouldFind:   true,
		},
		{
			name:         "path not found",
			searchPath:   "/nonexistent",
			expectedName: "",
			shouldFind:   false,
		},
		{
			name:         "empty path",
			searchPath:   "",
			expectedName: "",
			shouldFind:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := targets.FindByPath(tt.searchPath)
			if tt.shouldFind {
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedName, result.Name)
				assert.Equal(t, tt.searchPath, result.Path)
			} else {
				assert.Nil(t, result)
			}
		})
	}
}

func TestTarget(t *testing.T) {
	t.Run("basic target structure", func(t *testing.T) {
		target := &swiftpkg.Target{
			Name:               "TestTarget",
			C99name:            "TestTarget",
			Type:               swiftpkg.LibraryTargetType,
			ModuleType:         swiftpkg.SwiftModuleType,
			Path:               "/test/path",
			Sources:            []string{"file1.swift", "file2.swift"},
			Dependencies:       []*swiftpkg.TargetDependency{},
			ProductMemberships: []string{"TestProduct"},
		}

		assert.Equal(t, "TestTarget", target.Name)
		assert.Equal(t, "TestTarget", target.C99name)
		assert.Equal(t, swiftpkg.LibraryTargetType, target.Type)
		assert.Equal(t, swiftpkg.SwiftModuleType, target.ModuleType)
		assert.Equal(t, "/test/path", target.Path)
		assert.Equal(t, []string{"file1.swift", "file2.swift"}, target.Sources)
		assert.Equal(t, []string{"TestProduct"}, target.ProductMemberships)
	})
}

func TestTargets_EmptySlice(t *testing.T) {
	targets := swiftpkg.Targets{}

	t.Run("find by name in empty slice", func(t *testing.T) {
		result := targets.FindByName("AnyTarget")
		assert.Nil(t, result)
	})

	t.Run("find by path in empty slice", func(t *testing.T) {
		result := targets.FindByPath("/any/path")
		assert.Nil(t, result)
	})
}
