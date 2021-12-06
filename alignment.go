package fdisk

// #cgo LDFLAGS: -lfdisk
// #include <stdlib.h>
// #include <libfdisk/libfdisk.h>
import "C"

func (c Context) AlignLBA(lba Sector, direction Direction) Sector {
	x := C.fdisk_align_lba(cCtx(c), C.fdisk_sector_t(lba), C.int(direction))
	return Sector(x)
}

func (c Context) AlignLBAInRange(lba, start, stop Sector) Sector {
	x := C.fdisk_align_lba_in_range(cCtx(c), C.fdisk_sector_t(lba),
		C.fdisk_sector_t(start), C.fdisk_sector_t(stop))
	return Sector(x)
}

func (c Context) LBAIsPhyAligned(sector Sector) bool {
	x := C.fdisk_lba_is_phy_aligned(cCtx(c), cSector(sector))
	return x == 1
}

func (c Context) OverrideGeometry(cylinders, heads, sectors uint64) error {
	x, err := C.fdisk_override_geometry(cCtx(c), C.uint(cylinders), C.uint(heads), C.uint(sectors))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) SaveUserGeometry(cylinders, heads, sectors uint64) error {
	x, err := C.fdisk_save_user_geometry(cCtx(c), C.uint(cylinders), C.uint(heads), C.uint(sectors))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) SaveUserSectorSize(phy, log uint64) error {
	x, err := C.fdisk_save_user_sector_size(cCtx(c), C.uint(phy), C.uint(log))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) SaveUserGrain(grain uint64) error {
	x, err := C.fdisk_save_user_grain(cCtx(c), C.ulong(grain))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) HasUserDeviceProperties() bool {
	x := C.fdisk_has_user_device_properties(cCtx(c))
	return x == 1
}

func (c Context) ResetAlignment() error {
	x, err := C.fdisk_reset_alignment(cCtx(c))
	if x != 0 {
		return err
	}
	return nil
}
func (c Context) ResetDeviceProperties() error {
	x, err := C.fdisk_reset_device_properties(cCtx(c))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) RereadPartitionTable() error {
	x, err := C.fdisk_reread_partition_table(cCtx(c))
	if x != 0 {
		return err
	}
	return nil
}

func (c Context) RereadChanges(org Table) error {
	x, err := C.fdisk_reread_changes(cCtx(c), cTable(org))
	if x != 0 {
		return err
	}
	return nil
}
