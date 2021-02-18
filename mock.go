package mock

import (
	"fmt"
	"reflect"
)

func Patch(src, dest interface{}) error {
	p, err := newPatch(reflect.ValueOf(src), reflect.ValueOf(dest))
	if err != nil {
		return err
	}
	applyWithLock(p)
	return nil
}

func PatchMethod(tty reflect.Type, methodName string, dest interface{}) error {
	methodType, ok := tty.MethodByName(methodName)
	if !ok {
		return fmt.Errorf("can't find method: %s", methodName)
	}
	p, err := newPatch(methodType.Func, reflect.ValueOf(dest))
	if err != nil {
		return err
	}

	applyWithLock(p)
	return nil
}

func Unpatch(src interface{}) {
	undoWithLock(reflect.ValueOf(src).Pointer())
}

func UnpatchMethod(srcType reflect.Type, methodName string) error {
	methodType, ok := srcType.MethodByName(methodName)
	if !ok {
		return fmt.Errorf("can't find method: %s", methodName)
	}
	undoWithLock(methodType.Func.Pointer())
	return nil
}

func UnpatchAll() {
	undoAll()
}
