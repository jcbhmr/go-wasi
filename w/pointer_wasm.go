package w

import "unsafe"

type Pointer[T any] uint32

func ToPointer[T any](v *T) Pointer[T] {
	return Pointer[T](uintptr(unsafe.Pointer(v)))
}

func (p Pointer[T]) Get() *T {
	return (*T)(unsafe.Pointer(uintptr(p)))
}
