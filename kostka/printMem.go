package main

import (
	"fmt"
	"runtime"
)

const (
	BYTE = 1.0 << (10 * iota)
	KB
	MB
	GB
	TB
)

func printMem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("m.TotalAlloc:", m.TotalAlloc/MB, "MB")
}
