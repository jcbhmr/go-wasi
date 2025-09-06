package w

import "unsafe"

type Pointer32[T any] uint32

func ToPointer32[T any](v *T) Pointer32[T] {
	return Pointer32[T](uintptr(unsafe.Pointer(v)))
}

func (p Pointer32[T]) Get() *T {
	return (*T)(unsafe.Pointer(uintptr(p)))
}
