package fdisk

// #cgo LDFLAGS: -lfdisk
// #include <stdlib.h>
// #include <libfdisk/libfdisk.h>
import "C"

func NewTable() Table {
	x := C.fdisk_new_table()
	return goTable(x)
}

func (t Table) Reset() error {
	x, err := C.fdisk_reset_table(cTable(t))
	if x != 0 {
		return err
	}
	return nil
}

func (t Table) GetNEnts() uint64 {
	x := C.fdisk_table_get_nents(cTable(t))
	return uint64(x)
}

func (t Table) IsEmpty() bool {
	x := C.fdisk_table_is_empty(cTable(t))
	return x == 1
}

func (t Table) AddPartition(partition Partition) error {
	x, err := C.fdisk_table_add_partition(cTable(t), cPartition(partition))
	if x != 0 {
		return err
	}
	return nil
}

func (t Table) RemovePartition(partition Partition) error {
	x, err := C.fdisk_table_remove_partition(cTable(t), cPartition(partition))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) GetPartitions() (Table, error) {
	t := NewTable()
	cT := cTable(t)
	x, err := C.fdisk_get_partitions(cCtx(c), &cT)
	if x != 0 {
		t.Unref()
		return goTable(cT), err
	}
	return goTable(cT), err
}

func (c Context) GetFreeSpaces() (Table, error) {
	t := NewTable()
	cT := cTable(t)
	x, err := C.fdisk_get_freespaces(cCtx(c), &cT)
	if x != 0 {
		t.Unref()
		return goTable(cT), err
	}
	return goTable(cT), err
}

func (t Table) WrongOrder() bool {
	x := C.fdisk_table_wrong_order(cTable(t))
	return x == 1
}

// TODO: Implement!
//func (t Table) SortPartitions() error {
//}

// TODO: Implement!
//func (t Table) NextPartition(iter Iter) (Partition, error) {
//}

func (t Table) GetPartition(tableEntry uint64) (Partition, error) {
	x, err := C.fdisk_table_get_partition(cTable(t), C.size_t(tableEntry))
	if x == nil {
		return goPartition(x), err
	}
	return goPartition(x), nil
}

func (t Table) GetPartitionByPartNo(partNo uint64) (Partition, error) {
	x, err := C.fdisk_table_get_partition_by_partno(cTable(t), C.size_t(partNo))
	if x == nil {
		return goPartition(x), err
	}
	return goPartition(x), nil
}

func (c Context) ApplyTable(table Table) error {
	x, err := C.fdisk_apply_table(cCtx(c), cTable(table))
	if x != 0 {
		return err
	}
	return nil
}
