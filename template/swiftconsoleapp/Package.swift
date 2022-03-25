// swift-tools-version: 5.6
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "Swiftconsoleapp",
    dependencies: [
        // Dependencies declare other packages that this package depends on.
        // .package(url: /* package url */, from: "1.0.0"),
        .package(path: "../../../abstractions/swift"),
    ],
    targets: [
        // Targets are the basic building blocks of a package. A target can define a module or a test suite.
        // Targets can depend on other targets in this package, and on products in packages this package depends on.
        .executableTarget(
            name: "Swiftconsoleapp",
            dependencies: [.product(name: "MicrosoftKiotaAbstractions", package: "swift")]),
        .testTarget(
            name: "SwiftconsoleappTests",
            dependencies: [.product(name: "MicrosoftKiotaAbstractions", package: "swift"), "Swiftconsoleapp"]),
    ]
)
