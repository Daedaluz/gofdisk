package fdisk

// #cgo LDFLAGS: -lfdisk
// #include <stdlib.h>
// #include <libfdisk/libfdisk.h>
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func NewLabelItem() LabelItem {
	x := C.fdisk_new_labelitem()
	return goLabelItem(x)
}

func (l LabelItem) Reset() error {
	x, err := C.fdisk_reset_labelitem(cLabelItem(l))
	if x != 0 {
		return err
	}
	return nil
}

func (l LabelItem) GetName() string {
	x := C.fdisk_labelitem_get_name(cLabelItem(l))
	return C.GoString(x)
}

func (l LabelItem) GetID() int {
	x := C.fdisk_labelitem_get_id(cLabelItem(l))
	return int(x)
}

func (l LabelItem) GetDataU64() uint64 {
	o := C.uint64_t(0)
	// TODO: check return value?
	_ = C.fdisk_labelitem_get_data_u64(cLabelItem(l), &o)
	return uint64(o)
}

func (l LabelItem) GetDataString() string {
	var o *C.char
	// TODO: check return value?
	_ = C.fdisk_labelitem_get_data_string(cLabelItem(l), &o)
	return C.GoString(o)
}

func (l LabelItem) IsString() bool {
	x := C.fdisk_labelitem_is_string(cLabelItem(l))
	return x == 1
}

func (l LabelItem) IsNumber() bool {
	x := C.fdisk_labelitem_is_number(cLabelItem(l))
	return x == 1
}

var (
	ErrIDOutOfRange    = fmt.Errorf("id out of range")
	ErrUnsupportedItem = fmt.Errorf("unsupported item")
)

// TODO: perhaps rework call signature
func (c Context) GetDiskLabelItem(id int) (LabelItem, error) {
	res := NewLabelItem()
	x := C.int(0)
	var err error
	x, err = C.fdisk_get_disklabel_item(cCtx(c), C.int(id), cLabelItem(res))
	if x != 0 {
		res.Unref()
	}
	if x < 0 {
		return res, err
	}
	switch x {
	case 1:
		return res, ErrUnsupportedItem
	case 2:
		return res, ErrIDOutOfRange
	}
	return res, fmt.Errorf("unknown error")
}

func (c Context) GetDiskLabelID() (string, error) {
	var o *C.char
	var err error
	var x C.int
	x, err = C.fdisk_get_disklabel_id(cCtx(c), &o)
	if x != 0 {
		return "", err
	}
	oStr := C.GoString(o)
	C.free(unsafe.Pointer(o))
	return oStr, nil
}

func (c Context) SetDiskLabelID() error {
	x, err := C.fdisk_set_disklabel_id(cCtx(c))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) SetDiskLabelIDFromString(str string) error {
	arg := C.CString(str)
	x, err := C.fdisk_set_disklabel_id_from_string(cCtx(c), arg)
	C.free(unsafe.Pointer(arg))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) GetPartition(partNo uint64) (Partition, error) {
	part := NewPartition()
	cPart := cPartition(part)
	x, err := C.fdisk_get_partition(cCtx(c), C.size_t(partNo), &cPart)
	if x != 0 {
		part.Unref()
		return part, err
	}
	return part, nil
}

func (c Context) SetPartition(partNo uint, partition Partition) error {
	x, err := C.fdisk_set_partition(cCtx(c), C.size_t(partNo), cPartition(partition))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) AddPartition(partition Partition) (uint64, error) {
	var oPartNum C.size_t
	x, err := C.fdisk_add_partition(cCtx(c), cPartition(partition), &oPartNum)
	if x != 0 {
		return 0, err
	}
	return uint64(oPartNum), nil
}

func (c Context) DeletePartition(partNo uint64) error {
	x, err := C.fdisk_delete_partition(cCtx(c), C.size_t(partNo))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) DeleteAllPartitions() error {
	x, err := C.fdisk_delete_all_partitions(cCtx(c))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) WipePartition(partNo uint64, enable bool) error {
	cEnable := C.int(0)
	if enable {
		cEnable = 1
	}
	x, err := C.fdisk_wipe_partition(cCtx(c), C.size_t(partNo), cEnable)
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) SetPartitionType(partNo uint64, partType PartType) error {
	x, err := C.fdisk_set_partition_type(cCtx(c), C.size_t(partNo), cPartType(partType))
	if x != 0 {
		return err
	}
	return nil
}

func (l Label) GetFieldsIDs(ctx Context) ([]FieldTypes, error) {
	out := struct {
		Ids  *C.int
		Size C.size_t
	}{}
	x, err := C.fdisk_label_get_fields_ids(cLabel(l), cCtx(ctx), &out.Ids, &out.Size)
	if x != 0 {
		return nil, err
	}
	var arr []C.int
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&arr))
	hdr.Data = uintptr(unsafe.Pointer(out.Ids))
	hdr.Cap = int(out.Size)
	hdr.Len = int(out.Size)

	res := make([]FieldTypes, 0, int(out.Size))
	for _, n := range arr {
		res = append(res, FieldTypes(n))
	}
	return res, nil
}

func (l Label) GetFieldsIDsAll(ctx Context) ([]FieldTypes, error) {
	out := struct {
		Ids  *C.int
		Size C.size_t
	}{}
	x, err := C.fdisk_label_get_fields_ids_all(cLabel(l), cCtx(ctx), &out.Ids, &out.Size)
	if x != 0 {
		return nil, err
	}
	var arr []C.int
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&arr))
	hdr.Data = uintptr(unsafe.Pointer(out.Ids))
	hdr.Cap = int(out.Size)
	hdr.Len = int(out.Size)

	res := make([]FieldTypes, 0, int(out.Size))
	for _, n := range arr {
		res = append(res, FieldTypes(n))
	}
	return res, nil
}

func (l Label) GetField(ID FieldTypes) Field {
	x := C.fdisk_label_get_field(cLabel(l), C.int(ID))
	return goField(x)
}

func (l Label) GetFieldByName(name string) Field {
	cStr := C.CString(name)
	x := C.fdisk_label_get_field_by_name(cLabel(l), cStr)
	C.free(unsafe.Pointer(cStr))
	return goField(x)
}

func (l Label) SetChanged(changed bool) {
	cChanged := C.int(0)
	if changed {
		cChanged = 1
	}
	C.fdisk_label_set_changed(cLabel(l), cChanged)
}

func (l Label) IsChanged() bool {
	x := C.fdisk_label_is_changed(cLabel(l))
	return x == 1
}

func (l Label) SetDisabled(disabled bool) {
	cDisabled := C.int(0)
	if disabled {
		cDisabled = 1
	}
	C.fdisk_label_set_changed(cLabel(l), cDisabled)
}

func (l Label) IsDisabled() bool {
	x := C.fdisk_label_is_disabled(cLabel(l))
	return x == 1
}

func (c Context) IsPartitionUsed(partNo uint64) bool {
	x := C.fdisk_is_partition_used(cCtx(c), C.size_t(partNo))
	return x == 1
}

func (c Context) TogglePartitionFlag(partNo uint64, flagID uint64) error {
	//TODO: FlagID?
	x, err := C.fdisk_toggle_partition_flag(cCtx(c), C.size_t(partNo), C.ulong(flagID))
	if x != 0 {
		return err
	}
	return nil
}

func NewPartition() Partition {
	x := C.fdisk_new_partition()
	return goPartition(x)
}

func (p Partition) Reset() {
	C.fdisk_reset_partition(cPartition(p))
}

func (p Partition) IsFreeSpace() bool {
	x := C.fdisk_partition_is_freespace(cPartition(p))
	return x == 1
}

func (p Partition) SetStart(offset uint64) error {
	x, err := C.fdisk_partition_set_start(cPartition(p), C.uint64_t(offset))
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) UnsetStart() error {
	x, err := C.fdisk_partition_unset_start(cPartition(p))
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) GetStart() Sector {
	x := C.fdisk_partition_get_start(cPartition(p))
	return goSector(x)
}

func (p Partition) HasStart() bool {
	x := C.fdisk_partition_has_start(cPartition(p))
	return x == 1
}

func (p Partition) CompareStart(partition Partition) int {
	x := C.fdisk_partition_cmp_start(cPartition(p), cPartition(partition))
	return int(x)
}

func (p Partition) StartFollowDefault(enable bool) error {
	cEnable := C.int(0)
	if enable {
		cEnable = 1
	}
	x, err := C.fdisk_partition_start_follow_default(cPartition(p), cEnable)
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) StartIsDefault() bool {
	x := C.fdisk_partition_start_is_default(cPartition(p))
	return x == 1
}

func (p Partition) SetSize(size uint64) error {
	x, err := C.fdisk_partition_set_size(cPartition(p), C.uint64_t(size))
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) UnsetSize() error {
	x, err := C.fdisk_partition_unset_size(cPartition(p))
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) GetSize() Sector {
	x := C.fdisk_partition_get_size(cPartition(p))
	return goSector(x)
}

func (p Partition) HasSize() bool {
	x := C.fdisk_partition_has_size(cPartition(p))
	return x == 1
}

func (p Partition) SizeExplicit(enable bool) error {
	cEnable := C.int(0)
	if enable {
		cEnable = 1
	}
	x, err := C.fdisk_partition_size_explicit(cPartition(p), cEnable)
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) HasEnd() bool {
	x := C.fdisk_partition_has_end(cPartition(p))
	return x == 1
}

func (p Partition) GetEnd() Sector {
	x := C.fdisk_partition_get_end(cPartition(p))
	return goSector(x)
}

func (p Partition) SetPartNo(partNo uint64) error {
	x, err := C.fdisk_partition_set_partno(cPartition(p), C.size_t(partNo))
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) UnsetPartNo() error {
	x, err := C.fdisk_partition_unset_partno(cPartition(p))
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) GetPartNo() uint64 {
	x := C.fdisk_partition_get_partno(cPartition(p))
	return uint64(x)
}

func (p Partition) HasPartNo() bool {
	x := C.fdisk_partition_has_partno(cPartition(p))
	return x == 1
}

func (p Partition) ComparePartNo(partition Partition) int {
	x := C.fdisk_partition_cmp_partno(cPartition(p), cPartition(partition))
	return int(x)
}

func (p Partition) PartNoFollowDefault(enable bool) error {
	cEnable := C.int(0)
	if enable {
		cEnable = 1
	}
	x, err := C.fdisk_partition_partno_follow_default(cPartition(p), cEnable)
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) SetType(typ PartType) error {
	x, err := C.fdisk_partition_set_type(cPartition(p), cPartType(typ))
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) GetType() PartType {
	x := C.fdisk_partition_get_type(cPartition(p))
	return goPartType(x)
}

func (p Partition) SetName(name string) error {
	cStr := C.CString(name)
	x, err := C.fdisk_partition_set_name(cPartition(p), cStr)
	C.free(unsafe.Pointer(cStr))
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) GetName() string {
	x := C.fdisk_partition_get_name(cPartition(p))
	return C.GoString(x)
}

func (p Partition) SetUUID(uuid string) error {
	cUUID := C.CString(uuid)
	x, err := C.fdisk_partition_set_uuid(cPartition(p), cUUID)
	C.free(unsafe.Pointer(cUUID))
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) SetAttrs(attrs string) error {
	cAttrs := C.CString(attrs)
	x, err := C.fdisk_partition_set_attrs(cPartition(p), cAttrs)
	C.free(unsafe.Pointer(cAttrs))
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) GetUUID() string {
	x := C.fdisk_partition_get_uuid(cPartition(p))
	return C.GoString(x)
}

func (p Partition) GetAttrs() string {
	x := C.fdisk_partition_get_attrs(cPartition(p))
	return C.GoString(x)
}

func (p Partition) IsNested() bool {
	x := C.fdisk_partition_is_nested(cPartition(p))
	return x == 1
}

func (p Partition) IsContainer() bool {
	x := C.fdisk_partition_is_container(cPartition(p))
	return x == 1
}

func (p Partition) GetParent() (uint64, error) {
	var res C.uint64_t
	x, err := C.fdisk_partition_get_parent(cPartition(p), &res)
	if x != 0 {
		return 0, err
	}
	return uint64(res), nil
}

func (p Partition) IsUsed() bool {
	x := C.fdisk_partition_is_used(cPartition(p))
	return x == 1
}

func (p Partition) IsBootable() bool {
	x := C.fdisk_partition_is_bootable(cPartition(p))
	return x == 1
}

func (p Partition) IsWholeDisk() bool {
	x := C.fdisk_partition_is_wholedisk(cPartition(p))
	return x == 1
}

func (p Partition) ToString(ctx Context, field FieldTypes) (string, error) {
	var data *C.char
	x, err := C.fdisk_partition_to_string(cPartition(p), cCtx(ctx), C.int(field), &data)
	res := C.GoString(data)
	C.free(unsafe.Pointer(data))
	if x != 0 {
		return "", err
	}
	return res, nil
}

func (p Partition) NextPartNo(ctx Context) (uint64, error) {
	var partNo C.size_t
	x, err := C.fdisk_partition_next_partno(cPartition(p), cCtx(ctx), &partNo)
	if x != 0 {
		return 0, err
	}
	return uint64(partNo), nil
}

func (p Partition) EndFollowDefault(enable bool) error {
	cEnable := C.int(0)
	if enable {
		cEnable = 1
	}
	x, err := C.fdisk_partition_end_follow_default(cPartition(p), cEnable)
	if x != 0 {
		return err
	}
	return nil
}

func (p Partition) EndIsDefault() bool {
	x := C.fdisk_partition_end_is_default(cPartition(p))
	return x == 1
}

func (c Context) ReorderPartitions() error {
	x, err := C.fdisk_reorder_partitions(cCtx(c))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) PartitionHasWipe(partition Partition) bool {
	x := C.fdisk_partition_has_wipe(cCtx(c), cPartition(partition))
	return x == 1
}
