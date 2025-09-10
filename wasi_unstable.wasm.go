//go:build wasm

package wasi

type errno32 uint32
type advice32 uint32
type fdflags32 uint32
type fstflags32 uint32
type whence32 uint32
type oflags32 uint32
type signal32 uint32
type riflags32 uint32
type sdflags32 uint32
type siflags32 uint32

//go:wasmimport wasi_unstable args_get
//go:noescape
func wasmimport_args_get(argv *uint32, argvBuf *uint8) errno32

//go:wasmimport wasi_unstable args_sizes_get
//go:noescape
func wasmimport_args_sizes_get(t0 *Size, t1 *Size) errno32

//go:wasmimport wasi_unstable environ_get
//go:noescape
func wasmimport_environ_get(environ *uint32, environBuf *uint8) errno32

//go:wasmimport wasi_unstable environ_sizes_get
//go:noescape
func wasmimport_environ_sizes_get(t0 *Size, t1 *Size) errno32

//go:wasmimport wasi_unstable clock_res_get
//go:noescape
func wasmimport_clock_res_get(id Clockid, t *Timestamp) errno32

//go:wasmimport wasi_unstable clock_time_get
//go:noescape
func wasmimport_clock_time_get(id Clockid, precision Timestamp, t *Timestamp) errno32

//go:wasmimport wasi_unstable fd_advise
//go:noescape
func wasmimport_fd_advise(fd Fd, offset Filesize, len Filesize, advice advice32) errno32

//go:wasmimport wasi_unstable fd_allocate
//go:noescape
func wasmimport_fd_allocate(fd Fd, offset Filesize, len Filesize) errno32

//go:wasmimport wasi_unstable fd_close
//go:noescape
func wasmimport_fd_close(fd Fd) errno32

//go:wasmimport wasi_unstable fd_datasync
//go:noescape
func wasmimport_fd_datasync(fd Fd) errno32

//go:wasmimport wasi_unstable fd_fdstat_get
//go:noescape
func wasmimport_fd_fdstat_get(fd Fd, f *Fdstat) errno32

//go:wasmimport wasi_unstable fd_fdstat_set_flags
//go:noescape
func wasmimport_fd_fdstat_set_flags(fd Fd, flags fdflags32) errno32

//go:wasmimport wasi_unstable fd_fdstat_set_rights
//go:noescape
func wasmimport_fd_fdstat_set_rights(fd Fd, fsRightsBase Rights, fsRightsInheriting Rights) errno32

//go:wasmimport wasi_unstable fd_filestat_get
//go:noescape
func wasmimport_fd_filestat_get(fd Fd, f *Filestat) errno32

//go:wasmimport wasi_unstable fd_filestat_set_size
//go:noescape
func wasmimport_fd_filestat_set_size(fd Fd, size Filesize) errno32

//go:wasmimport wasi_unstable fd_filestat_set_times
//go:noescape
func wasmimport_fd_filestat_set_times(fd Fd, atim Timestamp, mtim Timestamp, fstFlags fstflags32) errno32

//go:wasmimport wasi_unstable fd_pread
//go:noescape
func wasmimport_fd_pread(fd Fd, iovsData *Iovec, iovsLen uint32, offset Filesize, s *Size) errno32

//go:wasmimport wasi_unstable fd_prestat_get
//go:noescape
func wasmimport_fd_prestat_get(fd Fd, p *Prestat) errno32

//go:wasmimport wasi_unstable fd_prestat_dir_name
//go:noescape
func wasmimport_fd_prestat_dir_name(fd Fd, path *uint8, pathLen Size) errno32

//go:wasmimport wasi_unstable fd_pwrite
//go:noescape
func wasmimport_fd_pwrite(fd Fd, iovsData *Ciovec, iovsLen uint32, offset Filesize, s *Size) errno32

//go:wasmimport wasi_unstable fd_read
//go:noescape
func wasmimport_fd_read(fd Fd, iovsData *Iovec, iovsLen uint32, s *Size) errno32

//go:wasmimport wasi_unstable fd_readdir
//go:noescape
func wasmimport_fd_readdir(fd Fd, buf *uint8, bufLen Size, cookie Dircookie, s *Size) errno32

//go:wasmimport wasi_unstable fd_renumber
//go:noescape
func wasmimport_fd_renumber(fd Fd, to Fd) errno32

//go:wasmimport wasi_unstable fd_seek
//go:noescape
func wasmimport_fd_seek(fd Fd, offset Filedelta, whence whence32, f *Filesize) errno32

//go:wasmimport wasi_unstable fd_sync
//go:noescape
func wasmimport_fd_sync(fd Fd) errno32

//go:wasmimport wasi_unstable fd_tell
//go:noescape
func wasmimport_fd_tell(fd Fd, f *Filesize) errno32

//go:wasmimport wasi_unstable fd_write
//go:noescape
func wasmimport_fd_write(fd Fd, iovsData *Ciovec, iovsLen uint32, s *Size) errno32

//go:wasmimport wasi_unstable path_create_directory
//go:noescape
func wasmimport_path_create_directory(fd Fd, path string) errno32

//go:wasmimport wasi_unstable path_filestat_get
//go:noescape
func wasmimport_path_filestat_get(fd Fd, flags Lookupflags, path string, f *Filestat) errno32

//go:wasmimport wasi_unstable path_filestat_set_times
//go:noescape
func wasmimport_path_filestat_set_times(fd Fd, flags Lookupflags, path string, atim Timestamp, mtim Timestamp, fstFlags fstflags32) errno32

//go:wasmimport wasi_unstable path_link
//go:noescape
func wasmimport_path_link(oldFd Fd, oldFlags Lookupflags, oldPath string, newFd Fd, newPath string) errno32

//go:wasmimport wasi_unstable path_open
//go:noescape
func wasmimport_path_open(fd Fd, dirflags Lookupflags, path string, oflags oflags32, fsRightsBase Rights, fsRightsInheriting Rights, fdFlags fdflags32, f *Fd) errno32

//go:wasmimport wasi_unstable path_readlink
//go:noescape
func wasmimport_path_readlink(fd Fd, path string, buf *uint8, bufLen Size, s *Size) errno32

//go:wasmimport wasi_unstable path_remove_directory
//go:noescape
func wasmimport_path_remove_directory(fd Fd, path string) errno32

//go:wasmimport wasi_unstable path_rename
//go:noescape
func wasmimport_path_rename(fd Fd, oldPath string, newFd Fd, newPath string) errno32

//go:wasmimport wasi_unstable path_symlink
//go:noescape
func wasmimport_path_symlink(oldPath string, fd Fd, newPath string) errno32

//go:wasmimport wasi_unstable path_unlink_file
//go:noescape
func wasmimport_path_unlink_file(fd Fd, path string) errno32

//go:wasmimport wasi_unstable poll_oneoff
//go:noescape
func wasmimport_poll_oneoff(in *Subscription, out *Event, nsubscriptions Size, s *Size) errno32

//go:wasmimport wasi_unstable proc_exit
//go:noescape
func wasmimport_proc_exit(rval Exitcode)

//go:wasmimport wasi_unstable proc_raise
//go:noescape
func wasmimport_proc_raise(sig signal32) errno32

//go:wasmimport wasi_unstable sched_yield
//go:noescape
func wasmimport_sched_yield() errno32

//go:wasmimport wasi_unstable random_get
//go:noescape
func wasmimport_random_get(buf *uint8, bufLen Size) errno32

//go:wasmimport wasi_unstable sock_recv
//go:noescape
func wasmimport_sock_recv(fd Fd, riDataData *Iovec, riDataLen uint32, riFlags riflags32, t0 *Size, t1 *Roflags) errno32

//go:wasmimport wasi_unstable sock_send
//go:noescape
func wasmimport_sock_send(fd Fd, siDataData *Ciovec, siDataLen uint32, siFlags siflags32, s *Size) errno32

//go:wasmimport wasi_unstable sock_shutdown
//go:noescape
func wasmimport_sock_shutdown(fd Fd, how sdflags32) errno32
