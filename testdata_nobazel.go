//go:build !bazel
// +build !bazel

package main

// Asset is a  dummy replacement for when the bindata is not loaded.
func Asset(string) ([]byte, error) {
	panic("use bazel to compile this binary")
}
