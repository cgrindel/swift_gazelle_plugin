# Swift Gazelle Plugin for Bazel

[![Build](https://github.com/cgrindel/swift_gazelle_plugin/actions/workflows/ci.yml/badge.svg?event=schedule)](https://github.com/cgrindel/swift_gazelle_plugin/actions/workflows/ci.yml)

This repository contains a [Gazelle plugin] used to generate [rules_swift] targets based upon your
Swift source code..

## Table of Contents

<!-- MARKDOWN TOC: BEGIN -->
* [Quickstart](#quickstart)
  * [1. Configure your `MODULE.bazel` to use swift_gazelle_plugin.](#1-configure-your-modulebazel-to-use-swift_gazelle_plugin)
  * [2. Add Gazelle targets to `BUILD.bazel` at the root of your workspace.](#2-add-gazelle-targets-to-buildbazel-at-the-root-of-your-workspace)
  * [3. Create or update Bazel build files for your project.](#3-create-or-update-bazel-build-files-for-your-project)
  * [4. Build and test your project.](#4-build-and-test-your-project)
  * [5. Check in `MODULE.bazel`.](#5-check-in-modulebazel)
  * [6. Start coding](#6-start-coding)
* [Tips and Tricks](#tips-and-tricks)
<!-- MARKDOWN TOC: END -->

## Quickstart

The following provides a quick introduction on how to set up and use the features in this
repository. These instructions assume that you are using [Bazel modules] to load your external
dependencies. If you are using Bazel's legacy external dependency management, we recommend using
[Bazel's hybrid mode], then follow the steps in this quickstart guide.

### 1. Configure your `MODULE.bazel` to use [swift_gazelle_plugin].

Add a dependency on `swift_gazelle_plugin`.

<!-- BEGIN MODULE SNIPPET -->

```python
bazel_dep(name = "swift_gazelle_plugin", version = "0.2.2")
```

<!-- END MODULE SNIPPET -->

### 2. Add Gazelle targets to `BUILD.bazel` at the root of your workspace.

Add the following to the `BUILD.bazel` file at the root of your workspace.

```bzl
load("@gazelle//:def.bzl", "gazelle", "gazelle_binary")

# This declaration builds a Gazelle binary that incorporates all of the Gazelle
# plugins for the languages that you use in your workspace. In this example, we
# are only listing the Gazelle plugin for Swift from swift_gazelle_plugin.
gazelle_binary(
    name = "gazelle_bin",
    languages = [
        "@swift_gazelle_plugin//gazelle",
    ],
)

# This target updates the Bazel build files for your project. Run this target
# whenever you add or remove source files from your project.
gazelle(
    name = "update_build_files",
    gazelle = ":gazelle_bin",
)
```

### 3. Create or update Bazel build files for your project.

Generate/update the Bazel build files for your project by running the following:

```sh
bazel run //:update_build_files
```

### 4. Build and test your project.

Build and test your project.

```sh
bazel test //...
```

### 5. Check in `MODULE.bazel`.

- The `MODULE.bazel` contains the declarations for your external dependencies.

### 6. Start coding

You are ready to start coding.

## Tips and Tricks

The following are a few tips to consider as you work with your repository:

- When you add or remove source files, run `bazel run //:update_build_files`. This will
  create/update the Bazel build files in your project. It is designed to be fast and unobtrusive.
- If things do not appear to be working properly, run the following:
  - `bazel run //:update_build_files`
- Do yourself a favor and create a Bazel target (e.g., `//:tidy`) that runs your repository
  maintenance targets (e.g., `//:update_build_files`, formatting utilities)
  in the proper order. If you are looking for an easy way to set this up, check out the
  [`//:tidy` declaration in this repository](BUILD.bazel) and the documentation for the [tidy] macro.

<!-- Links -->

[Gazelle plugin]: https://github.com/bazelbuild/bazel-gazelle/blob/master/extend.md
[rules_swift]: https://github.com/bazelbuild/rules_swift
[swift_gazelle_plugin]: https://github.com/cgrindel/swift_gazelle_plugin
[tidy]: https://github.com/cgrindel/bazel-starlib/blob/main/doc/bzltidy/rules_and_macros_overview.md#tidy
