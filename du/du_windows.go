package du

import (
	"syscall"
	"unsafe"
)

type DiskUsage struct {
	freeBytes  int64
	totalBytes int64
	availBytes int64
}

// Returns an object holding the disk usage of volumePath
// This function assumes volumePath is a valid path
func New(volumePath string) *DiskUsage {

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	du := &DiskUsage{}

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&du.freeBytes)),
		uintptr(unsafe.Pointer(&du.totalBytes)),
		uintptr(unsafe.Pointer(&du.availBytes)))

	return du
}

// Total free bytes on file system
func (d *DiskUsage) Free() uint64 {
	return uint64(d.freeBytes)
}

// Total available bytes on file system to an unpriveleged user
func (d *DiskUsage) Available() uint64 {
	return uint64(d.availBytes)
}

// Total size of the file system
func (d *DiskUsage) Size() uint64 {
	return uint64(d.totalBytes)
}

// Total bytes used in file system
func (d *DiskUsage) Used() uint64 {
	return d.Size() - d.Free()
}

// Percentage of use on the file system
func (d *DiskUsage) Usage() float32 {
	return float32(d.Used()) / float32(d.Size())
}
