package fdisk

// #cgo LDFLAGS: -lfdisk
// #include <stdlib.h>
// #include <libfdisk/libfdisk.h>
import "C"

func (f Field) GetID() int {
	x := C.fdisk_field_get_id(cField(f))
	return int(x)
}

func (f Field) GetName() string {
	x := C.fdisk_field_get_name(cField(f))
	return C.GoString(x)
}

func (f Field) GetWidth() float64 {
	x := C.fdisk_field_get_width(cField(f))
	return float64(x)
}

func (f Field) IsNumber() bool {
	x := C.fdisk_field_is_number(cField(f))
	return x == 1
}
