package cimgui

// #cgo CPPFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #cgo amd64,darwin LDFLAGS: ${SRCDIR}/lib/macos/x64/cimgui.a -framework Cocoa -framework OpenGL -framework ApplicationServices -framework Carbon
// #cgo arm64,darwin LDFLAGS: ${SRCDIR}/lib/macos/arm64/cimgui.a -framework Cocoa -framework OpenGL -framework ApplicationServices -framework Carbon
// #include "cimgui_wrapper.h"
import "C"

func GetVersion() string {
	return C.GoString(C.GetVersion())
}