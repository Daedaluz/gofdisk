package fdisk

// #cgo LDFLAGS: -lfdisk
// #include <stdlib.h>
// #include <libfdisk/libfdisk.h>
import "C"

const (
	GPTFlagRequired      = 1
	GPTFlagNoBlock       = 2
	GPTFlagLegacyBoot    = 3
	GPTFLagGUIDSpedcific = 4
)

func (c Context) GPTIsHybrid() bool {
	x := C.fdisk_gpt_is_hybrid(cCtx(c))
	return x == 1
}

func (c Context) GPTSetNPartitions(nents uint32) error {
	x, err := C.fdisk_gpt_set_npartitions(cCtx(c), C.uint32_t(nents))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) GPTGetPartitionAttrs(partNo uint64) (uint64, error) {
	var out C.uint64_t
	x, err := C.fdisk_gpt_get_partition_attrs(cCtx(c), C.size_t(partNo), &out)
	if x != 0 {
		return 0, err
	}
	return uint64(out), nil
}

func (c Context) GPTSetPartitionAttrs(partNo, attrs uint64) error {
	x, err := C.fdisk_gpt_set_partition_attrs(cCtx(c), C.size_t(partNo), C.uint64_t(attrs))
	if x != 0 {
		return err
	}
	return nil
}

func (l Label) GPTDisableRelocation(disable bool) {
	cDisable := C.int(0)
	if disable {
		cDisable = 1
	}
	C.fdisk_gpt_disable_relocation(cLabel(l), cDisable)
}

func (l Label) GPTEnableMinimize(enable bool) {
	cEnable := C.int(0)
	if enable {
		cEnable = 1
	}
	C.fdisk_gpt_enable_minimize(cLabel(l), cEnable)
}
