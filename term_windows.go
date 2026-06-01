//go:build windows

package main

import (
	"syscall"
	"unsafe"
)

// enableVT turns on ANSI Virtual Terminal Processing on Windows consoles so
// our escape codes (colors, cursor control) render correctly. Uses only the
// standard-library syscall package — no external dependency.
func enableVT() {
	const enableVTProcessing = 0x0004
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getMode := kernel32.NewProc("GetConsoleMode")
	setMode := kernel32.NewProc("SetConsoleMode")

	h, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		return
	}
	var mode uint32
	if r, _, _ := getMode.Call(uintptr(h), uintptr(unsafe.Pointer(&mode))); r == 0 {
		return
	}
	setMode.Call(uintptr(h), uintptr(mode|enableVTProcessing))
}
