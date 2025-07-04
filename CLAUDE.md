# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this
repository.

## Overview

This repository contains a Gazelle plugin for the Swift programming language that automatically
generates Bazel BUILD files for Swift projects. The plugin is written in Go and integrates with
Bazel's Gazelle tool to analyze Swift source code and create appropriate build targets.

## Development Commands

### Core Development Commands

- `bazel test //...` - Run all tests in the project
- `bazel run //:tidy` - Run all maintenance tasks (formatting, updates, etc.)
- `bazel run //:update_files` - Quick update of source files and build files

### Allowed Tools

<!-- allowed_tools: Bash(bazel test //...), Bash(bazel run //:tidy), Bash(bazel run //:update_files) -->

### Testing

- `bazel test //examples:smoke_integration_tests` - Run smoke tests on examples
- `bazel test //examples:all_integration_tests` - Run all integration tests
- Examples have individual `do_test` scripts for testing specific scenarios

### Go Development

- `bazel run //:go_mod_tidy` - Clean up Go module dependencies
- `bazel run //:go_get_latest` - Update Go dependencies to latest versions

## Architecture

### Core Plugin Structure

- `/gazelle/` - Main Gazelle plugin implementation in Go
  - `lang.go` - Implements the Gazelle Language interface for Swift
  - `generate.go` - Core logic for generating BUILD files
  - `resolve.go` - Dependency resolution logic
  - `config.go` - Plugin configuration handling

### Internal Packages

- `/gazelle/internal/swift/` - Core Swift analysis and rule generation
  - Contains logic for Swift modules, packages, dependencies, and BUILD rule generation
- `/gazelle/internal/swiftpkg/` - Swift Package Manager integration
- `/gazelle/internal/swiftcfg/` - Configuration management
- `/gazelle/internal/spdump/` - Swift package description dumping
- `/gazelle/internal/spreso/` - Swift package resolution

### Examples and Testing

- `/examples/` - Working examples of different Swift project configurations
  - `simple/` - Basic Swift library and executable
  - `grpc_example/` - gRPC service implementation
  - `custom_swift_proto_compiler/` - Custom protobuf compiler setup
- `/bzlmod/` - Bzlmod-specific test workspace

### Build and CI

- `/ci/` - Continuous integration configuration and workflows
- `BUILD.bazel` (root) - Main build file with comprehensive tidy targets and test suites
- Uses Bazel's bzlmod for dependency management

## Key Concepts

### Swift Rule Generation

The plugin analyzes Swift source files and generates appropriate Bazel rules:

- `swift_library` for library code
- `swift_binary` for executables
- `swift_test` for test targets
- `swift_proto_library` for protocol buffer integration

### Integration Points

- Integrates with `rules_swift` for Swift compilation
- Supports `rules_swift_package_manager` for Swift Package Manager dependencies
- Works with both WORKSPACE and bzlmod dependency management

### Configuration

The plugin supports various configuration options through Gazelle directives in BUILD files and can
be customized for different Swift project structures.

## Code Style and Formatting

### Language-Specific Style Requirements

#### Go (Primary Language)

- Use 4-space indentation (not tabs)
- Wrap lines longer than 100 characters
- Use `gofmt` or `goimports` for formatting
- Follow standard Go conventions for naming and documentation
- Include comprehensive package documentation
- Use explicit error handling over silent failures

#### Starlark/Bazel (BUILD files and .bzl files)

- Use 4-space indentation
- Wrap lines longer than 80 characters
- Use `buildifier -mode=fix -lint=fix` for formatting
- Follow Bazel style guidelines

#### Markdown (Documentation)

- Use 2-space indentation for nested lists
- **MUST wrap lines at 100 characters maximum**
- Use `prettier` for formatting when possible
- Include descriptive comments for complex sections

#### Shell Scripts (Test and utility scripts)

- Always include shebang: `#!/usr/bin/env bash`
- Use `set -o errexit -o nounset -o pipefail` for error handling
- Use 2-space indentation
- Wrap lines longer than 80 characters
- Use `shfmt -i 2 -ci -bn` for formatting

#### Swift (Example projects and test cases)

- Use 2-space indentation
- Wrap lines longer than 100 characters
- Follow Swift naming conventions

#### YAML (CI/CD configuration)

- Use 2-space indentation
- Use `yamlfmt` for formatting when possible
- Follow YAML best practices for readability

### Style Enforcement Rules

#### Mandatory Pre-submission Checks

Before completing ANY code task, verify:

1. **Correct indentation** for the specific language
2. **Line length compliance** with the limits above
3. **Proper formatting tool** usage where specified
4. **Language-specific conventions** are followed

#### Critical Requirements

- **Go files**: Must pass `gofmt` validation
- **BUILD files**: Must be `buildifier` compatible
- **Markdown files**: MUST wrap at 100 characters (strictly enforced)
- **Shell scripts**: Must include shebang and error handling
- **All files**: No trailing whitespace

### Development Workflow Integration

#### Pre-submission Commands

1. `bazel run //:tidy` - Run all maintenance tasks (formatting, updates, etc.)
2. `bazel test //...` - Ensure all tests pass
3. Verify language-specific formatting compliance

#### Style Validation

When making changes:

- Check existing file style consistency
- Apply appropriate formatting tools
- Verify line length and indentation
- Confirm language-specific requirements

## Commit Message Guidelines

- Use conventional commit message format for this repository
