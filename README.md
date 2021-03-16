go-disk-usage
=============

[![Go Reference](https://pkg.go.dev/badge/tcw.im/go-disk-usage.svg)](https://pkg.go.dev/tcw.im/go-disk-usage)

Get disk usage information like how much space is available, free, and used.  

## Install

```bash
go get -u tcw.im/go-disk-usage
```

## Usage

```go
import "tcw.im/go-disk-usage/du"
usage := du.New("/path/to")
```

## Compatibility

This works for Windows, MacOS, and Linux although there may some minor variability between what this library reports and what you get from `df`.  This library will maintain reverse compatability, any breaking changes will be made to a forked repository.
