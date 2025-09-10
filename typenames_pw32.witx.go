//go:build wasm && tinygo

package wasi

import (
	"structs"
)

// A region of memory for scatter/gather reads.
type Iovec struct {
	_ structs.HostLayout
	// The address of the buffer to be filled.
	Buf *uint8
	// The length of the buffer to be filled.
	BufLen Size
}

// A region of memory for scatter/gather writes.
type Ciovec struct {
	_ structs.HostLayout
	// The address of the buffer to be written.
	Buf *uint8
	// The length of the buffer to be written.
	BufLen Size
}
