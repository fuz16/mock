package sys

import (
	"reflect"
	"syscall"
	"unsafe"
)

// LoadMemoryValue 读取内存信息
func LoadMemoryValue(start uintptr, length int) []byte {
	f := rawMemoryAccess(start, length)
	original := make([]byte, len(f))
	copy(original, f)
	return original
}

// pageStart 指针所在内存页起始位置
func pageStart(ptr uintptr) uintptr {
	return ptr & ^(uintptr(syscall.Getpagesize() - 1))
}

// rawMemoryAccess 将内存指定位置开始的指定长度，设置为一个slice，便于读写
func rawMemoryAccess(p uintptr, length int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  length,
		Cap:  length,
	}))
}
