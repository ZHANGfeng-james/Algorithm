package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Print(isLittleEndian())
}

func isLittleEndian() bool {
	var i int32 = 0x01020304

	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	b := *pb
	return (b == 0x04)
}
