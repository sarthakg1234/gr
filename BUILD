load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

genrule(
    name = "testdata",
    srcs = glob(["testdata/**"]),
    outs = ["testdata.go"],
    cmd = """
$(location @com_github_go_bindata_go_bindata//go-bindata) -pkg main -o $@ go/src/grail.com/cmd/tidydata-client/...
""",
    tools = ["@com_github_go_bindata_go_bindata//go-bindata"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "testdata.go",
    ],
    importpath = "grail.com/cmd/tidydata-client",
    visibility = ["//visibility:private"],
    deps = [
        "//go/src/github.com/grailbio/v23/factories/grail:go_default_library",
        "//go/src/grail.com/tidy/dataframe:go_default_library",
        "//go/src/grail.com/tidy/vanadium/client:go_default_library",
        "//go/src/grail.com/tidy/vanadium/vdl/dataset:go_default_library",
        "@io_v//v23:go_default_library",
        "@io_v//v23/context:go_default_library",
        "@io_v//x/ref/lib/v23cmd:go_default_library",
        "@io_v_x_lib//cmdline:go_default_library",
    ],
)

go_binary(
    name = "tidydata-client",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
