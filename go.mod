module github.com/cgrindel/swift_gazelle_plugin

go 1.23.3

// Workaround for inconsistent Go versions being used in rules_bazel_integration_test tests.
// toolchain go1.21.5

require (
	github.com/bazelbuild/bazel-gazelle v0.43.0
	github.com/bazelbuild/buildtools v0.0.0-20240918101019-be1c24cc9a44
	github.com/creasty/defaults v1.8.0
	github.com/deckarep/golang-set/v2 v2.8.0
	github.com/stretchr/testify v1.10.0
	golang.org/x/exp v0.0.0-20250531010427-b6e5de432a8b
	golang.org/x/text v0.26.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.25.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/tools/go/vcs v0.1.0-deprecated // indirect
)
