package w

import (
	"structs"
	"unsafe"
)

type List[T any] struct {
	_ structs.HostLayout
	list[T]
}

type list[T any] struct {
	_    structs.HostLayout
	data Pointer[T]
	len  int32
}

func (l list[T]) Data() *T {
	return l.data.Get()
}

func (l list[T]) Len() int {
	return int(l.len)
}

func (l list[T]) Slice() []T {
	return unsafe.Slice(l.Data(), l.Len())
}
