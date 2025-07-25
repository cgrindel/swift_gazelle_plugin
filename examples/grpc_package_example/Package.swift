// swift-tools-version: 5.7

import PackageDescription

let package = Package(
    name: "grpc_package_example",
    dependencies: [
        // These are the versions used by rules_swift
        .package(url: "https://github.com/grpc/grpc-swift.git", exact: "1.26.1"),
        .package(url: "https://github.com/apple/swift-protobuf.git", exact: "1.30.0"),
        .package(url: "https://github.com/apple/swift-nio", exact: "2.85.0"),
    ]
)
