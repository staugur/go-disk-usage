// +build !windows

package du

import "syscall"

type DiskUsage struct {
	stat *syscall.Statfs_t
}

// Returns an object holding the disk usage of volumePath
// This function assumes volumePath is a valid path
func New(volumePath string) *DiskUsage {
	var stat syscall.Statfs_t
	syscall.Statfs(volumePath, &stat)
	return &DiskUsage{&stat}
}

// Total free bytes on file system
func (d *DiskUsage) Free() uint64 {
	return d.stat.Bfree * uint64(d.stat.Bsize)
}

// Total available bytes on file system to an unpriveleged user
func (d *DiskUsage) Available() uint64 {
	return d.stat.Bavail * uint64(d.stat.Bsize)
}

// Total size of the file system
func (d *DiskUsage) Size() uint64 {
	return d.stat.Blocks * uint64(d.stat.Bsize)
}

// Total bytes used in file system
func (d *DiskUsage) Used() uint64 {
	return d.Size() - d.Free()
}

// Percentage of use on the file system
func (d *DiskUsage) Usage() float32 {
	return float32(d.Used()) / float32(d.Size())
}

// DiskRate returns the usage rate of the disk where volumePath is located
func DiskRate(volumePath string) float32 {
	d := New(volumePath)
	return d.Usage()
}

const KB = uint64(1024)

// Human-readable disk information, in KB (uint64)
type DiskHumanReadable struct {
	Total     uint64
	Used      uint64
	Free      uint64
	Available uint64
	Percent   float32 // Decimal places, you might have to multiply by 100
}

// DiskInfo returns an instance of DiskHumanReadable (New it)
func DiskInfo(volumePath string) DiskHumanReadable {
	d := New(volumePath)
	return DiskHumanReadable{
		d.Size() / KB, d.Used() / KB, d.Free() / KB, d.Available() / KB,
		d.Usage(),
	}
}
