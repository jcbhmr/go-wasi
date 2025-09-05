package w

import (
	"structs"
	"unsafe"
)

type Union[T any, U any, V any] struct {
	_ structs.HostLayout
	union[T, U, V]
}

type union[T any, U any, V any] struct {
	_    structs.HostLayout
	tag  T
	_    [0]V
	data U
}

func (u *union[T, U, V]) Tag() T {
	return u.tag
}

func (u *union[T, U, V]) Data() unsafe.Pointer {
	return unsafe.Pointer(&u.data)
}
