go-disk-usage
=============

[![Go Reference](https://pkg.go.dev/badge/pkg.tcw.im/go-disk-usage.svg)](https://pkg.go.dev/pkg.tcw.im/go-disk-usage)

Get disk usage information like how much space is available, free, and used.  

## Install

```bash
go get -u pkg.tcw.im/go-disk-usage
```

## Usage

```go
import "pkg.tcw.im/go-disk-usage/du"

// The first way
diskUsage := du.New("/path/to/one")

// The second way (Human-readable, recommend)
diskInfo := du.DiskInfo("/path/to/two")

// Or get disk usage directly
percent := du.DiskRate("/path/to")
```

## Compatibility

This works for Windows, MacOS, and Linux although there may some minor variability between what this library reports and what you get from `df`.  This library will maintain reverse compatability, any breaking changes will be made to a forked repository.
