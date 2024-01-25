# Declares that this directory is the root of a Bazel workspace.
# See https://docs.bazel.build/versions/main/build-ref.html#workspace
workspace(
    # How this workspace would be referenced with absolute labels from another workspace
    name = "go-commons",
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

## GO
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6734a719993b1ba4ebe9806e853864395a8d3968ad27f9dd759c196b3eb3abe8",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.45.1/rules_go-v0.45.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.45.1/rules_go-v0.45.1.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "32938bda16e6700063035479063d9d24c60eda8d79fd4739563f50d331cb3209",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

load("//:go_deps.bzl", "go_dependencies")

# gazelle:repository_macro go_deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.21.6")

gazelle_dependencies()

# Java toolchains configuration
http_archive(
    name = "rules_java",
    urls = [
        "https://github.com/bazelbuild/rules_java/releases/download/7.1.0/rules_java-7.1.0.tar.gz"
    ],
    sha256 = "b5ccdfdc5d0a43c5ba18a502dac6ec39add2d60cf4a30c7c15f99056bcb0ab64",
    strip_prefix = "rules_java-7.1.0",
)

# Remote JDK 11 toolchain - Windows
http_archive(
    name = "remotejdk11_win",
    urls = [
        "https://cdn.azul.com/zulu/bin/zulu11.52.13-ca-jdk11.0.13-win_x64.zip"
    ],
    sha256 = "c16c8338fda57412ac6dbacb85946245423b52d9bce7ed83fa25b3b5c2a0a6f6",
    build_file = "@rules_java//java:defs.bzl"
)

# Remote JDK 11 toolchain - Linux AArch64
http_archive(
    name = "remotejdk11_linux_aarch64",
    urls = [
        "https://cdn.azul.com/zulu/bin/zulu11.52.13-ca-jdk11.0.13-linux_aarch64.tar.gz"
    ],
    sha256 = "8ec22226e0fdfd8e7ac39f3e74dc8d5a3c6db3defef6b90ff688c6bdbf5defc4",
    build_file = "@rules_java//java:defs.bzl"
)

# Remote JDK 17 toolchain - Linux
http_archive(
    name = "remotejdk17_linux",
    urls = [
        "https://cdn.azul.com/zulu/bin/zulu17.34.19-ca-jdk17.0.2-linux_x64.tar.gz"
    ],
    sha256 = "12cb53d52d84b827d5ddf5b9edb8fd6fd9a2f648f23e7da4ee8ad2e4e3be8f20",
    build_file = "@rules_java//java:defs.bzl"
)

# Remote JDK 11 toolchain - Linux s390x
http_archive(
    name = "remotejdk11_linux_s390x",
    urls = [
        "https://cdn.azul.com/zulu/bin/zulu11.52.13-ca-jdk11.0.13-linux_s390x.tar.gz"
    ],
    sha256 = "1ecb503f6d837dabe9b4e140490a3eb0c9db1596f0ca70625c46ad92c249e833",
    build_file = "@rules_java//java:defs.bzl"
)
