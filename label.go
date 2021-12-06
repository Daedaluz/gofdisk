package fdisk

// #cgo LDFLAGS: -lfdisk
// #include <stdlib.h>
// #include <libfdisk/libfdisk.h>
import "C"
import (
	"unsafe"
)

func (l Label) GetType() LabelType {
	x := C.fdisk_label_get_type(cLabel(l))
	return goLabelType(uint32(x))
}

func (l Label) GetName() string {
	x := C.fdisk_label_get_name(cLabel(l))
	return C.GoString(x)
}

func (l Label) RequireGeometry() bool {
	x := C.fdisk_label_require_geometry(cLabel(l))
	return x == 1
}

func (c Context) WriteDiskLabel() error {
	x, err := C.fdisk_write_disklabel(cCtx(c))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) VerifyDiskLabel() error {
	x, err := C.fdisk_verify_disklabel(cCtx(c))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) CreateDiskLabel(name string) error {
	cName := C.CString(name)
	x, err := C.fdisk_create_disklabel(cCtx(c), cName)
	C.free(unsafe.Pointer(cName))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) ListDiskLabel() int {
	x := C.fdisk_list_disklabel(cCtx(c))
	return int(x)
}

type DiskLabelLocation struct {
	Name   string
	Offset uint64
	Size   uint64
}

func (c Context) LocateDiskLabel(n int) (int, *DiskLabelLocation) {
	out := struct {
		Name   *C.char
		Offset C.uint64_t
		Size   C.size_t
	}{}

	x := C.fdisk_locate_disklabel(cCtx(c), C.int(n),
		&out.Name, &out.Offset, &out.Size)
	return int(x), &DiskLabelLocation{
		Name:   C.GoString(out.Name),
		Offset: uint64(out.Offset),
		Size:   uint64(out.Size),
	}
}

func (l Label) GetGeomRangeCylinders() (res int, min, max Sector) {
	x := C.fdisk_label_get_geomrange_cylinders(cLabel(l), (*C.ulong)(&min), (*C.ulong)(&max))
	res = int(x)
	return
}

func (l Label) GetGeomRangeHeads() (res int, min, max uint64) {
	cmin := C.uint(0)
	cmax := C.uint(0)
	x := C.fdisk_label_get_geomrange_heads(cLabel(l), &cmin, &cmax)
	min, max = uint64(cmin), uint64(cmax)
	res = int(x)
	return
}

func (l Label) GetGeomRangeSectors() (res int, min, max Sector) {
	x := C.fdisk_label_get_geomrange_sectors(cLabel(l), (*C.ulong)(&min), (*C.ulong)(&max))
	res = int(x)
	return
}
