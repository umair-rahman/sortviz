//go:build !windows

package main

// enableVT is a no-op on Unix-like systems, where terminals already support
// ANSI escape sequences out of the box.
func enableVT() {}
