"""
Utilities for proto compiler rules.
"""

load(
    "@bazel_skylib//lib:dicts.bzl",
    "dicts",
)
load(
    "@rules_swift//proto:proto.bzl",
    "swift_proto_compiler",
)

# NOTE: The ProtoPathModuleMappings option is set internally for all plugins.
# This is used to inform the plugins which Swift module the generated code for each plugin is located in.
PROTO_PLUGIN_OPTION_ALLOWLIST = [
    "FileNaming",
    "Visibility",
]
PROTO_PLUGIN_OPTIONS = {
    "Visibility": "Public",
}
GRPC_VARIANT_SERVER = "Server"
GRPC_VARIANT_CLIENT = "Client"
GRPC_VARIANT_TEST_CLIENT = "TestClient"
GRPC_VARIANTS = [
    GRPC_VARIANT_SERVER,
    GRPC_VARIANT_CLIENT,
    GRPC_VARIANT_TEST_CLIENT,
]
GRPC_PLUGIN_OPTION_ALLOWLIST = PROTO_PLUGIN_OPTION_ALLOWLIST + [
    "KeepMethodCasing",
    "ExtraModuleImports",
    "GRPCModuleName",
    "SwiftProtobufModuleName",
] + GRPC_VARIANTS

def make_grpc_swift_proto_compiler(
        name,
        variants,
        plugin_options = PROTO_PLUGIN_OPTIONS):
    """Generates a GRPC swift_proto_compiler target for the given variants.

    Args:
        name: The name of the generated swift proto compiler target.
        variants: The list of variants the compiler should generate.
        plugin_options: Additional options to pass to the plugin.
    """

    # Merge the plugin options to include the variants:
    merged_plugin_options = dicts.add(
        plugin_options,
        {variant: "false" for variant in GRPC_VARIANTS},
    )
    for variant in variants:
        merged_plugin_options[variant] = "true"

    swift_proto_compiler(
        name = name,
        protoc = "@com_google_protobuf//:protoc",
        plugin = "@swiftpkg_grpc_swift//:protoc-gen-grpc-swift",
        plugin_name = name.removesuffix("_proto"),
        plugin_option_allowlist = GRPC_PLUGIN_OPTION_ALLOWLIST,
        plugin_options = merged_plugin_options,
        suffixes = [".grpc.swift"],
        deps = [
            "@swiftpkg_swift_protobuf//:SwiftProtobuf",
            "@swiftpkg_grpc_swift//:GRPC",
        ],
        visibility = ["//visibility:public"],
    )
