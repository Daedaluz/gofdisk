package fdisk

// #cgo LDFLAGS: -lfdisk
// #include <stdlib.h>
// #include <libfdisk/libfdisk.h>
import "C"
import "unsafe"

func NewPartType() PartType {
	x := C.fdisk_new_parttype()
	return goPartType(x)
}

func (p PartType) SetName(name string) error {
	str := C.CString(name)
	x, err := C.fdisk_parttype_set_name(cPartType(p), str)
	C.free(unsafe.Pointer(str))
	if x != 0 {
		return err
	}
	return nil
}

func (p PartType) SetType(typ string) error {
	str := C.CString(typ)
	x, err := C.fdisk_parttype_set_typestr(cPartType(p), str)
	C.free(unsafe.Pointer(str))
	if x != 0{
		return err
	}
	return nil
}

func (p PartType) SetCode(code int) error {
	x, err := C.fdisk_parttype_set_code(cPartType(p), C.int(code))
	if x != 0{
		return err
	}
	return nil
}

func (l Label) GetNPartTypes() uint64 {
	x := C.fdisk_label_get_nparttypes(cLabel(l))
	return uint64(x)
}

func (l Label) GetPartType(n uint64) PartType {
	x := C.fdisk_label_get_parttype(cLabel(l), C.size_t(n))
	return goPartType(x)
}

func (l Label) HasCodePartTypes() bool {
	x := C.fdisk_label_has_code_parttypes(cLabel(l))
	return x == 1
}

func (l Label) GetPartTypeFromCode(code uint64) (PartType, error) {
	x, err := C.fdisk_label_get_parttype_from_code(cLabel(l), C.uint(code))
	if x == nil {
		return goPartType(x), err
	}
	return goPartType(x), nil
}

func (l Label) GetPartTypeFromString(str string) (PartType,error) {
	cstr := C.CString(str)
	x, err := C.fdisk_label_get_parttype_from_string(cLabel(l), cstr)
	C.free(unsafe.Pointer(cstr))
	if x == nil {
		return goPartType(x), err
	}
	return goPartType(x), nil
}

func NewUnknownPartType(code uint64, typeStr string) PartType {
	cstr := C.CString(typeStr)
	x := C.fdisk_new_unknown_parttype(C.uint(code), cstr)
	C.free(unsafe.Pointer(cstr))
	return goPartType(x)
}

func (p PartType) Copy() PartType {
	x := C.fdisk_copy_parttype(cPartType(p))
	return goPartType(x)
}

func (l Label) ParsePartType(str string) (PartType, error) {
	cstr := C.CString(str)
	x, err := C.fdisk_label_parse_parttype(cLabel(l), cstr)
	C.free(unsafe.Pointer(cstr))
	if x == nil {
		return goPartType(x), err
	}
	return goPartType(x), nil
}

func (l Label) AdvParsePartType(str string, flags PartTypeParserFlag) PartType {
	cstr := C.CString(str)
	x := C.fdisk_label_advparse_parttype(cLabel(l), cstr, C.int(flags))
	C.free(unsafe.Pointer(cstr))
	return goPartType(x)
}

func (p PartType) GetString() string {
	x := C.fdisk_parttype_get_string(cPartType(p))
	return C.GoString(x)
}

func (p PartType) GetCode() uint64 {
	x := C.fdisk_parttype_get_code(cPartType(p))
	return uint64(x)
}

func (p PartType) GetName() string {
	x := C.fdisk_parttype_get_name(cPartType(p))
	return C.GoString(x)
}

func (p PartType) IsUnknown() bool {
	x := C.fdisk_parttype_is_unknown(cPartType(p))
	return x == 1
}
