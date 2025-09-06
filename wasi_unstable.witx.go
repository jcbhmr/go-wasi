//go:build wasm

/*
https://github.com/WebAssembly/WASI/blob/main/legacy/preview0/witx/wasi_unstable.witx

This API predated the convention of naming modules with a `wasi_unstable_`
prefix and a version number. It is preserved here for compatibility, but
we shouldn't follow this pattern in new APIs.

Some content here is derived from https://github.com/NuxiNL/cloudabi.
*/
package wasi

import (
	"github.com/jcbhmr/go-wasi/w"
)

// Read command-line argument data.
// The size of the array should match that returned by `args_sizes_get`.
// Each argument is expected to be `\0` terminated.
//
//go:nosplit
func ArgsGet(argv *w.Pointer32[uint8], argvBuf *uint8) (r w.Result[struct{}, Errno]) {
	r.Set(struct{}{}, Errno(wasmimport_args_get(argv, argvBuf)))
	return
}

// Return command-line argument data sizes.
//
// Returns: Returns the number of arguments and the size of the argument string
// data, or an error.
//
//go:nosplit
func ArgsSizesGet() (w.Tuple[Size, Size], Errno) {
	var t w.Tuple[Size, Size]
	e := Errno(wasmimport_args_sizes_get(&t.A, &t.B))
	return t, e
}

//go:nosplit
func EnvironGet(environ *w.Pointer32[uint8], environBuf *uint8) Errno {
	return Errno(wasmimport_environ_get(environ, environBuf))
}

//go:nosplit
func EnvironSizesGet() (w.Tuple[Size, Size], Errno) {
	var t w.Tuple[Size, Size]
	e := Errno(wasmimport_environ_sizes_get(&t.A, &t.B))
	return t, e
}

//go:nosplit
func ClockResGet(id Clockid) (Timestamp, Errno) {
	var t Timestamp
	e := Errno(wasmimport_clock_res_get(id, &t))
	return t, e
}

//go:nosplit
func ClockTimeGet(id Clockid, precision Timestamp) (Timestamp, Errno) {
	var t Timestamp
	e := Errno(wasmimport_clock_time_get(id, precision, &t))
	return t, e
}

//go:nosplit
func FdAdvise(fd Fd, offset Filesize, len Filesize, advice Advice) Errno {
	return Errno(wasmimport_fd_advise(fd, offset, len, advice32(advice)))
}

//go:nosplit
func FdAllocate(fd Fd, offset Filesize, len Filesize) Errno {
	return Errno(wasmimport_fd_allocate(fd, offset, len))
}

//go:nosplit
func FdClose(fd Fd) Errno {
	return Errno(wasmimport_fd_close(fd))
}

//go:nosplit
func FdDatasync(fd Fd) Errno {
	return Errno(wasmimport_fd_datasync(fd))
}

