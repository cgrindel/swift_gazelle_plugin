// swift-tools-version: 5.7

import PackageDescription

let package = Package(
    name: "grpc_example",
    dependencies: [
        .package(url: "https://github.com/apple/swift-protobuf.git", exact: "1.32.0")
    ]
)
