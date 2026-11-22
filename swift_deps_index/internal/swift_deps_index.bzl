"""Implementation for the `swift_deps_index` rule."""

# The output groups rules_swift_package_manager's `swift_deps_info` target
# provides: its manifest (identity to pkg_info.json path), and the
# pkg_info.json files it names.
_MANIFEST_OUTPUT_GROUP = "swift_deps_info_manifest"
_PKG_INFOS_OUTPUT_GROUP = "swift_deps_info_pkg_infos"

def _swift_deps_index_impl(ctx):
    out = ctx.actions.declare_file("{}.json".format(ctx.label.name))
    args = ctx.actions.args()
    args.add("create")
    args.add("-o", out)
    inputs = []

    if ctx.attr.swift_deps_info:
        output_groups = ctx.attr.swift_deps_info[OutputGroupInfo]
        if not hasattr(output_groups, _MANIFEST_OUTPUT_GROUP):
            fail("{} doesn't provide a swift_deps_info manifest; use the \
`@swift_deps_info` repository from rules_swift_package_manager.".format(
                ctx.attr.swift_deps_info.label,
            ))
        manifests = getattr(output_groups, _MANIFEST_OUTPUT_GROUP).to_list()
        if len(manifests) != 1:
            fail("Expected one swift_deps_info manifest, got {}.".format(len(manifests)))
        args.add("-swift_deps_info", manifests[0])
        inputs.append(manifests[0])
        if hasattr(output_groups, _PKG_INFOS_OUTPUT_GROUP):
            inputs.extend(getattr(output_groups, _PKG_INFOS_OUTPUT_GROUP).to_list())

    for (idx, identity) in enumerate(ctx.attr.direct_dep_pkg_infos.values()):
        pi_file = ctx.files.direct_dep_pkg_infos[idx]
        args.add("{}={}".format(identity, pi_file.path))
    inputs.extend(ctx.files.direct_dep_pkg_infos)

    ctx.actions.run(
        outputs = [out],
        inputs = inputs,
        executable = ctx.executable._swift_deps_index_tool,
        arguments = [args],
        mnemonic = "SwiftDepsIndex",
        progress_message = "Indexing Swift dependencies for %{label}",
    )

    return [DefaultInfo(files = depset([out]))]

swift_deps_index = rule(
    implementation = _swift_deps_index_impl,
    attrs = {
        "direct_dep_pkg_infos": attr.label_keyed_string_dict(
            allow_files = True,
            doc = """\
A `dict` where the key is the label for a Swift package's `pkg_info.json` file \
and the value is the Swift package's identity value.\
""",
        ),
        "swift_deps_info": attr.label(
            doc = """\
The `swift_deps_info` target from rules_swift_package_manager (usually \
`@swift_deps_info`), which lists the direct Swift package dependencies and \
their `pkg_info.json` files.\
""",
        ),
        "_swift_deps_index_tool": attr.label(
            executable = True,
            cfg = "exec",
            default = Label("//gazelle/cmd/swift_deps_index"),
        ),
    },
    doc = """\
Generates the Swift dependencies index that the Swift Gazelle plugin reads \
(`-swift_dependency_index`).

```starlark
load("@swift_gazelle_plugin//swift_deps_index:defs.bzl", "swift_deps_index")

swift_deps_index(
    name = "swift_deps_index",
    swift_deps_info = "@swift_deps_info",
)

gazelle(
    name = "update_build_files",
    data = [":swift_deps_index"],
    extra_args = ["-swift_dependency_index=$(location :swift_deps_index)"],
    gazelle = ":gazelle_bin",
)
```
""",
)
