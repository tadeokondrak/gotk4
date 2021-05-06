// Code generated by girgen. DO NOT EDIT.

package gdkx

import (
	"github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 gtk4-x11
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
import "C"

func init() {
	glib.RegisterGValueMarshalers([]glib.TypeMarshaler{
		// Enums
		// Skipped X11DeviceType.

		// Objects/Classes
	})
}

type X11DeviceType int

const (
	X11DeviceTypeLogical  X11DeviceType = 0
	X11DeviceTypePhysical X11DeviceType = 1
	X11DeviceTypeFloating X11DeviceType = 2
)
