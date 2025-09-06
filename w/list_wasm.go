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
	data Pointer32[T]
	len  uint32
}

func (l list[T]) Data() *T {
	return l.data.Get()
}

func (l list[T]) Len() int {
	return int(l.len)
}

func (l list[T]) Slice() []T {
	return unsafe.Slice(l.data.Get(), l.len)
}

type AnyList[T any] interface {
	~struct {
		_ structs.HostLayout
		list[T]
	}
}

func (l *list[T]) Lift(data Pointer32[T], len uint32) {
	l.data = data
	l.len = len
}

func (l list[T]) Lower() (Pointer32[T], uint32) {
	return l.data, uint32(l.len)
}
