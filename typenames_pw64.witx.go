//go:build wasm && !tinygo

package wasi

import (
	"structs"

	"github.com/jcbhmr/bytecodealliance-go-modules/cm"
)

// A region of memory for scatter/gather reads.
type Iovec struct {
	// The address of the buffer to be filled.
	Buf *uint8
	// The length of the buffer to be filled.
	BufLen Size
}

type Iovec32 struct {
	_      structs.HostLayout
	Buf    uint32
	BufLen Size
}

func (i *Iovec) To32() *Iovec32 {
	return &Iovec32{
		Buf:    cm.PointerToU32(i.Buf),
		BufLen: i.BufLen,
	}
}

func (i *Iovec) From32(v *Iovec32) {
	i.Buf = cm.U32ToPointer[uint8](v.Buf)
	i.BufLen = v.BufLen
}

// A region of memory for scatter/gather writes.
type Ciovec struct {
	// The address of the buffer to be written.
	Buf *uint8
	// The length of the buffer to be written.
	BufLen Size
}

type Ciovec32 struct {
	_      structs.HostLayout
	Buf    uint32
	BufLen Size
}

func (i *Ciovec) To32() *Ciovec32 {
	return &Ciovec32{
		Buf:    cm.PointerToU32(i.Buf),
		BufLen: i.BufLen,
	}
}

func (i *Ciovec) From32(v *Ciovec32) {
	i.Buf = cm.U32ToPointer[uint8](v.Buf)
	i.BufLen = v.BufLen
}

