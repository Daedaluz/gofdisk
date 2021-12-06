package fdisk

// #cgo LDFLAGS: -lfdisk
// #include <stdlib.h>
// #include <libfdisk/libfdisk.h>
import "C"
import (
	"unsafe"
)

func NewContext() Context {
	return goCtx(C.fdisk_new_context())
}

func (c Context) NewNestedContext(name string) Context {
	cstr := C.CString(name)
	x := C.fdisk_new_nested_context(cCtx(c), cstr)
	C.free(unsafe.Pointer(cstr))
	return goCtx(x)
}

func (c Context) GetParent() Context {
	x := C.fdisk_get_parent(cCtx(c))
	return Context(unsafe.Pointer(x))
}

func (c Context) GetNPartitions() uint64 {
	x := C.fdisk_get_npartitions(cCtx(c))
	return uint64(x)
}

func (c Context) GetLabel(name string) Label {
	var cStr *C.char
	if name != "" {
		cStr = C.CString(name)
	}
	x := C.fdisk_get_label(cCtx(c), cStr)
	C.free(unsafe.Pointer(cStr))
	return goLabel(x)
}

func (c Context) NextLabel(label *Label) int {
	tmp := cLabel(*label)
	x := C.fdisk_next_label(cCtx(c), (**C.struct_fdisk_label)(&tmp))
	*label = goLabel(tmp)
	return int(x)
}

func (c Context) GetNLabels() uint64 {
	x := C.fdisk_get_nlabels(cCtx(c))
	return uint64(x)
}

func (c Context) HasLabel() bool {
	x := C.fdisk_has_label(cCtx(c))
	return x == 1
}

func (c Context) IsLabelType(label LabelType) bool {
	labelType := C.enum_fdisk_labeltype(label)
	x := C.fdisk_is_labeltype(cCtx(c), labelType)
	return x == 1
}

func (c Context) AssignDevice(name string, readonly bool) error {
	cname := C.CString(name)
	ro := C.int(0)
	if readonly {
		ro = 1
	}
	x, err := C.fdisk_assign_device(cCtx(c), cname, ro)
	C.free(unsafe.Pointer(cname))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) AssignDeviceByFD(fd uintptr, fname string, readonly bool) error {
	cname := C.CString(fname)
	ro := C.int(0)
	if readonly {
		ro = 1
	}
	x, err := C.fdisk_assign_device_by_fd(cCtx(c), C.int(fd), cname, ro)
	C.free(unsafe.Pointer(cname))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) DeAssignDevice(noSync bool) error {
	nsync := C.int(0)
	if noSync {
		nsync = 1
	}
	x, err := C.fdisk_deassign_device(cCtx(c), nsync)
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) ReAssaignDevice() error {
	x, err := C.fdisk_reassign_device(cCtx(c))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) IsReadOnly() bool {
	x := C.fdisk_is_readonly(cCtx(c))
	return x == 1
}

func (c Context) IsRegFile() bool {
	x := C.fdisk_is_regfile(cCtx(c))
	return x == 1
}

func (c Context) DeviceIsUsed() bool {
	x := C.fdisk_device_is_used(cCtx(c))
	return x == 1
}

func (c Context) DisableDialogs(disable bool) error {
	cDisable := C.int(0)
	if disable {
		cDisable = 1
	}
	x, err := C.fdisk_disable_dialogs(cCtx(c), cDisable)
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) HasDialogs() bool {
	x := C.fdisk_has_dialogs(cCtx(c))
	return x == 1
}

func (c Context) EnableDetails(enable bool) error {
	enabled := C.int(0)
	if enable {
		enabled = 1
	}
	x, err := C.fdisk_enable_details(cCtx(c), enabled)
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) IsDetails() bool {
	x := C.fdisk_is_details(cCtx(c))
	return x == 1
}

func (c Context) EnableListOnly(enable bool) error {
	enabled := C.int(0)
	if enable {
		enabled = 1
	}
	x, err := C.fdisk_enable_listonly(cCtx(c), enabled)
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) IsListOnly() bool {
	x := C.fdisk_is_listonly(cCtx(c))
	return x == 1
}

func (c Context) EnableWipe(enable bool) error {
	enabled := C.int(0)
	if enable {
		enabled = 1
	}
	x, err := C.fdisk_enable_wipe(cCtx(c), enabled)
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) HasWipe() bool {
	x := C.fdisk_has_wipe(cCtx(c))
	return x == 1
}

func (c Context) GetCollision() string {
	x := C.fdisk_get_collision(cCtx(c))
	return C.GoString(x)
}

func (c Context) IsPTCollision() bool {
	x := C.fdisk_is_ptcollision(cCtx(c))
	return x == 1
}

func (c Context) SetUnit(unit Unit) error {
	cunit := C.CString(string(unit))
	x, err := C.fdisk_set_unit(cCtx(c), cunit)
	C.free(unsafe.Pointer(cunit))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) GetUnit(plural bool) string {
	arg := C.int(1)
	if plural {
		arg = 0
	}
	x := C.fdisk_get_unit(cCtx(c), arg)
	return C.GoString(x)
}

func (c Context) UseCylinders() bool {
	x := C.fdisk_use_cylinders(cCtx(c))
	return x == 1
}

func (c Context) GetUnitsPerSector() uint64 {
	x := C.fdisk_get_units_per_sector(cCtx(c))
	return uint64(x)
}

func (c Context) GetOptimalIOSize() uint64 {
	x := C.fdisk_get_optimal_iosize(cCtx(c))
	return uint64(x)
}

func (c Context) GetMinimalIOSize() uint64 {
	x := C.fdisk_get_minimal_iosize(cCtx(c))
	return uint64(x)
}

func (c Context) GetPhySectorSize() uint64 {
	x := C.fdisk_get_physector_size(cCtx(c))
	return uint64(x)
}

func (c Context) GetSectorSize() uint64 {
	x := C.fdisk_get_sector_size(cCtx(c))
	return uint64(x)
}

func (c Context) GetAlignmentOffset() uint64 {
	x := C.fdisk_get_alignment_offset(cCtx(c))
	return uint64(x)
}

func (c Context) GetGrainSize() uint64 {
	x := C.fdisk_get_grain_size(cCtx(c))
	return uint64(x)
}

func (c Context) GetFirstLBA() Sector {
	x := C.fdisk_get_first_lba(cCtx(c))
	return Sector(x)
}

func (c Context) SetFirstLBA(sector Sector) Sector {
	sec := cSector(sector)
	x := C.fdisk_set_first_lba(cCtx(c), sec)
	return goSector(x)
}

func (c Context) GetLastLBA() Sector {
	x := C.fdisk_get_last_lba(cCtx(c))
	return goSector(x)
}

func (c Context) SetLastLBA(sector Sector) Sector {
	sec := cSector(sector)
	x := C.fdisk_set_last_lba(cCtx(c), sec)
	return goSector(x)
}

func (c Context) GetNSectors() Sector {
	x := C.fdisk_get_nsectors(cCtx(c))
	return Sector(x)
}

func (c Context) GetDevName() string {
	x := C.fdisk_get_devname(cCtx(c))
	return C.GoString(x)
}

func (c Context) GetDevFD() uintptr {
	x := C.fdisk_get_devfd(cCtx(c))
	return uintptr(x)
}

func (c Context) GetDevModel() string {
	x := C.fdisk_get_devmodel(cCtx(c))
	return C.GoString(x)
}

func (c Context) GetGeomHeads() uint64 {
	x := C.fdisk_get_geom_heads(cCtx(c))
	return uint64(x)
}

func (c Context) GetGeomSectors() Sector {
	x := C.fdisk_get_geom_sectors(cCtx(c))
	return Sector(x)
}

func (c Context) GetGeomCylinders() Sector {
	x := C.fdisk_get_geom_cylinders(cCtx(c))
	return Sector(x)
}

func (c Context) GetSizeUint() SizeUnit {
	x := C.fdisk_get_size_unit(cCtx(c))
	return SizeUnit(x)
}

func (c Context) SetSizeUnit(unit SizeUnit) error {
	x, err := C.fdisk_set_size_unit(cCtx(c), C.int(unit))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) HasProtectedBootBits() bool {
	x := C.fdisk_has_protected_bootbits(cCtx(c))
	return x == 1
}

func (c Context) EnableBootBitsProtection(enable bool) error {
	enabled := C.int(0)
	if enable {
		enabled = 1
	}
	x, err := C.fdisk_enable_bootbits_protection(cCtx(c), enabled)
	if x != 0 {
		return err
	}
	return nil
}
