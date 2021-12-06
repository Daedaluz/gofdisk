package fdisk

// #cgo LDFLAGS: -lfdisk
// #include <stdlib.h>
// #include <libfdisk/libfdisk.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// Context c representation
type Context uintptr

func cCtx(ctx Context) *C.struct_fdisk_context {
	return (*C.struct_fdisk_context)(unsafe.Pointer(ctx))
}

func goCtx(ctx *C.struct_fdisk_context) Context {
	return Context(unsafe.Pointer(ctx))
}

func (c Context) Ref() {
	C.fdisk_ref_context(cCtx(c))
}

func (c Context) UnRef() {
	C.fdisk_unref_context(cCtx(c))
}

// Label c representation
type Label uintptr

func cLabel(label Label) *C.struct_fdisk_label {
	return (*C.struct_fdisk_label)(unsafe.Pointer(label))
}

func goLabel(label *C.struct_fdisk_label) Label {
	return Label(unsafe.Pointer(label))
}

// PartType c representation
type PartType uintptr

func cPartType(partType PartType) *C.struct_fdisk_parttype {
	return (*C.struct_fdisk_parttype)(unsafe.Pointer(partType))
}

func goPartType(partType *C.struct_fdisk_parttype) PartType {
	return PartType(unsafe.Pointer(partType))
}

func (p PartType) Ref() {
	C.fdisk_ref_parttype(cPartType(p))
}

func (p PartType) UnRef() {
	C.fdisk_unref_parttype(cPartType(p))
}

// Partition c representation
type Partition uintptr

func cPartition(partition Partition) *C.struct_fdisk_partition {
	return (*C.struct_fdisk_partition)(unsafe.Pointer(partition))
}

func goPartition(partition *C.struct_fdisk_partition) Partition {
	return Partition(unsafe.Pointer(partition))
}

func (p Partition) Ref() {
	C.fdisk_ref_partition(cPartition(p))
}

func (p Partition) Unref() {
	C.fdisk_unref_partition(cPartition(p))
}

// Iter c representation
type Iter uintptr

func cIter(iter Iter) *C.struct_fdisk_iter {
	return (*C.struct_fdisk_iter)(unsafe.Pointer(iter))
}

func goIter(iter *C.struct_fdisk_iter) Iter {
	return Iter(unsafe.Pointer(iter))
}

// Table c representation
type Table uintptr

func cTable(table Table) *C.struct_fdisk_table {
	return (*C.struct_fdisk_table)(unsafe.Pointer(table))
}

func goTable(table *C.struct_fdisk_table) Table {
	return Table(unsafe.Pointer(table))
}

func (t Table) Ref() {
	C.fdisk_ref_table(cTable(t))
}

func (t Table) Unref() {
	C.fdisk_unref_table(cTable(t))
}

// Field c representation
type Field uintptr

func cField(field Field) *C.struct_fdisk_field {
	return (*C.struct_fdisk_field)(unsafe.Pointer(field))
}

func goField(field *C.struct_fdisk_field) Field {
	return Field(unsafe.Pointer(field))
}

// Script c representation
type Script uintptr

func cScript(script Script) *C.struct_fdisk_script {
	return (*C.struct_fdisk_script)(unsafe.Pointer(script))
}

func goScript(script *C.struct_fdisk_script) Script {
	return Script(unsafe.Pointer(script))
}

// Sector c representation
type Sector uint64

func cSector(sector Sector) C.fdisk_sector_t {
	return C.fdisk_sector_t(sector)
}

func goSector(sector C.fdisk_sector_t) Sector {
	return Sector(sector)
}

// LabelType c representation
type LabelType uint32

const (
	LabelDOS = LabelType(1 << (iota + 1))
	LabelSUN
	LabelSGI
	LabelBSD
	LabelGPT
)

func (l LabelType) String() string {
	switch l {
	case LabelDOS:
		return "dos"
	case LabelSUN:
		return "sun"
	case LabelSGI:
		return "sgi"
	case LabelBSD:
		return "bsd"
	case LabelGPT:
		return "gpt"
	}
	return fmt.Sprintf("unknown(%x)", int(l))
}

func cLabelType(labelType LabelType) C.enum_fdisk_labeltype {
	return C.enum_fdisk_labeltype(labelType)
}

func goLabelType(labelType C.enum_fdisk_labeltype) LabelType {
	return LabelType(labelType)
}

// LabelItem c representation
type LabelItem uintptr

func cLabelItem(item LabelItem) *C.struct_fdisk_labelitem {
	return (*C.struct_fdisk_labelitem)(unsafe.Pointer(item))
}

func goLabelItem(item *C.struct_fdisk_labelitem) LabelItem {
	return LabelItem(unsafe.Pointer(item))
}

func (l LabelItem) Ref() {
	C.fdisk_ref_labelitem(cLabelItem(l))
}

func (l LabelItem) Unref() {
	C.fdisk_unref_labelitem(cLabelItem(l))
}

// AskType c representation
type AskType int

const (
	AskNone = iota
	AskNumber
	AskOffset
	AskWarn
	AskWarnX
	AskInfo
	AskYesNo
	AskString
	AskMenu
)
