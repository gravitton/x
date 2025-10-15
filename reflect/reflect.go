package reflect

import (
	"reflect"
)

// SizeOf returns the size of a type in bytes.
func SizeOf(tp reflect.Type) uintptr {
	size, align := tp.Size(), uintptr(tp.Align())
	return (size + (align - 1)) / align * align
}
