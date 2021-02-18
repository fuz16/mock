package sys

import (
	"syscall"
)

// setMemPageAccess 设置数据页访问权限
func setMemPageAccess(addr uintptr, length int, prot int) {
	pageSize := syscall.Getpagesize()
	for p := pageStart(addr); p < addr+uintptr(length); p += uintptr(pageSize) {
		page := rawMemoryAccess(p, pageSize)
		err := syscall.Mprotect(page, prot)
		if err != nil {
			panic(err)
		}
	}
}

// WriteToMem 将一个 byte slice 写入到内存指定位置
func WriteToMem(location uintptr, data []byte) {
	f := rawMemoryAccess(location, len(data))

	setMemPageAccess(location, len(data), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
	copy(f, data[:])
	setMemPageAccess(location, len(data), syscall.PROT_READ|syscall.PROT_EXEC)
}
