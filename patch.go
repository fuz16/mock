package mock

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/sanbsy/mock/sys"
)

type patch struct {
	src  uintptr
	dst  uintptr
	data []byte
}

// newPatch 创建一个patch
func newPatch(sty, dty reflect.Value) (*patch, error) {
	if sty.Kind() != reflect.Func {
		return nil, errors.New("invalid type: src")
	}

	if dty.Kind() != reflect.Func {
		return nil, errors.New("invalid type: dst")
	}

	if sty.Type() != sty.Type() {
		return nil, errors.New("type of src and dst must be same")
	}
	return &patch{
		src:  sty.Pointer(),
		dst:  (uintptr)(getPtr(dty)),
		data: nil,
	}, nil
}

// apply 应用patch
func (p *patch) apply() {
	jumpData := sys.JmpToFunctionValue(p.dst)
	originData := sys.LoadMemoryValue(p.src, len(jumpData))
	sys.WriteToMem(p.src, jumpData)
	p.data = originData
}

// undo 撤销patch
func (p *patch) undo() {
	if len(p.data) > 0 {
		sys.WriteToMem(p.src, p.data)
	}
}

// getPtr 获取到 reflect.Value 结构体私有变量ptr
func getPtr(v reflect.Value) unsafe.Pointer {
	return (*struct {
		_   uintptr
		ptr unsafe.Pointer
	})(unsafe.Pointer(&v)).ptr
}
