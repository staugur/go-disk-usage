package du_test

import (
	"fmt"

	"tcw.im/go-disk-usage/du"
)

var KB = uint64(1024)

func ExampleDiskUsage() {
	usage := du.New(".")
	fmt.Println("Free:", usage.Free()/(KB*KB))
	fmt.Println("Available:", usage.Available()/(KB*KB))
	fmt.Println("Size:", usage.Size()/(KB*KB))
	fmt.Println("Used:", usage.Used()/(KB*KB))
	fmt.Println("Usage:", usage.Usage()*100, "%")
}
