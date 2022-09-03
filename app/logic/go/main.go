package main

import (
	"unsafe"

	"edgecomputing/logic"
)

//export t
func t(r, g, b int) bool {
	isWarm, _ := logic.Type(uint8(r), uint8(g), uint8(b))
	return isWarm
}

//export n
func n(r, g, b int) {
	s := logic.Name(uint8(r), uint8(g), uint8(b))
	name = *(*[]byte)(unsafe.Pointer(&s))
}

var name []byte

//export getNAddress
func getNAddress() *byte {
	return &name[0]
}

//export getNLen
func getNLen() int {
	return len(name)
}

func main() {}
