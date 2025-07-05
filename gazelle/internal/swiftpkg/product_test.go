package swiftpkg_test

import (
    "testing"

    "github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/spdump"
    "github.com/cgrindel/swift_gazelle_plugin/gazelle/internal/swiftpkg"
    "github.com/stretchr/testify/assert"
)

func TestNewProductFromManifestInfo(t *testing.T) {
    tests := []struct {
        name          string
        dumpProduct   *spdump.Product
        expectedError bool
        expectedType  swiftpkg.ProductType
    }{
        {
            name: "executable product",
            dumpProduct: &spdump.Product{
                Name:    "MyExecutable",
                Targets: []string{"MyExecutable"},
                Type:    spdump.ExecutableProductType,
            },
            expectedError: false,
            expectedType: swiftpkg.ProductType{
                Executable:   true,
                IsExecutable: true,
                IsLibrary:    false,
                IsMacro:      false,
                IsPlugin:     false,
                Library:      nil,
            },
        },
        {
            name: "library product",
            dumpProduct: &spdump.Product{
                Name:    "MyLibrary",
                Targets: []string{"MyLibrary"},
                Type:    spdump.LibraryProductType,
            },
            expectedError: false,
            expectedType: swiftpkg.ProductType{
                Executable:   false,
                IsExecutable: false,
                IsLibrary:    true,
                IsMacro:      false,
                IsPlugin:     false,
                Library: map[string]string{
                    "kind": "automatic",
                },
            },
        },
        {
            name: "plugin product",
            dumpProduct: &spdump.Product{
                Name:    "MyPlugin",
                Targets: []string{"MyPlugin"},
                Type:    spdump.PluginProductType,
            },
            expectedError: false,
            expectedType: swiftpkg.ProductType{
                Executable:   false,
                IsExecutable: false,
                IsLibrary:    false,
                IsMacro:      false,
                IsPlugin:     true,
                Library:      nil,
            },
        },
        {
            name: "unknown product type",
            dumpProduct: &spdump.Product{
                Name:    "Unknown",
                Targets: []string{"Unknown"},
                Type:    spdump.UnknownProductType,
            },
            expectedError: false,
            expectedType: swiftpkg.ProductType{
                Executable:   false,
                IsExecutable: false,
                IsLibrary:    false,
                IsMacro:      false,
                IsPlugin:     false,
                Library:      nil,
            },
        },
        {
            name: "invalid product type",
            dumpProduct: &spdump.Product{
                Name:    "Invalid",
                Targets: []string{"Invalid"},
                Type:    spdump.ProductType(999), // Invalid type
            },
            expectedError: true,
            expectedType:  swiftpkg.ProductType{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            product, err := swiftpkg.NewProductFromManifestInfo(tt.dumpProduct)

            if tt.expectedError {
                assert.Error(t, err)
                assert.Nil(t, product)
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, product)
                assert.Equal(t, tt.dumpProduct.Name, product.Name)
                assert.Equal(t, tt.dumpProduct.Targets, product.Targets)
                assert.Equal(t, tt.expectedType, product.Type)
            }
        })
    }
}

func TestProductType(t *testing.T) {
    t.Run("executable product type", func(t *testing.T) {
        pt := swiftpkg.ProductType{
            Executable:   true,
            IsExecutable: true,
        }
        assert.True(t, pt.Executable)
        assert.True(t, pt.IsExecutable)
        assert.False(t, pt.IsLibrary)
        assert.False(t, pt.IsPlugin)
    })

    t.Run("library product type", func(t *testing.T) {
        pt := swiftpkg.ProductType{
            IsLibrary: true,
            Library: map[string]string{
                "kind": "automatic",
            },
        }
        assert.False(t, pt.Executable)
        assert.False(t, pt.IsExecutable)
        assert.True(t, pt.IsLibrary)
        assert.False(t, pt.IsPlugin)
        assert.Equal(t, "automatic", pt.Library["kind"])
    })

    t.Run("plugin product type", func(t *testing.T) {
        pt := swiftpkg.ProductType{
            IsPlugin: true,
        }
        assert.False(t, pt.Executable)
        assert.False(t, pt.IsExecutable)
        assert.False(t, pt.IsLibrary)
        assert.True(t, pt.IsPlugin)
    })
}
