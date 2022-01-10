package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	var a uint
	var b int
	fmt.Println(runtime.GOARCH)
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
}
