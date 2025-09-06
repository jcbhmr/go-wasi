//go:build wasm

/*
https://github.com/WebAssembly/WASI/blob/main/legacy/preview0/witx/typenames.witx

Type names used by low-level WASI interfaces.

Some content here is derived from https://github.com/NuxiNL/cloudabi.
*/
package wasi

import (
	"structs"

	"github.com/jcbhmr/go-wasi/w"
)

type Size uint32

// Non-negative file size or length of a region within a file.
type Filesize uint64

// Timestamp in nanoseconds.
type Timestamp uint64

// Identifiers for clocks.
type Clockid uint32

const (
	// The clock measuring real time. Time value zero corresponds with
	// 1970-01-01T00:00:00Z.
	ClockidRealtime Clockid = iota
	// The store-wide monotonic clock, which is defined as a clock measuring
	// real time, whose value cannot be adjusted and which cannot have negative
	// clock jumps. The epoch of this clock is undefined. The absolute time
	// value of this clock therefore has no meaning.
	ClockidMonotonic
	// The CPU-time clock associated with the current process.
	ClockidProcessCputimeId
	// The CPU-time clock associated with the current thread.
	ClockidThreadCputimeId
)

var clockidNames = [...]string{"realtime", "monotonic", "process_cputime_id"}

func (c Clockid) String() string {
	return clockidNames[c]
}

// Error codes returned by functions.
// Not all of these error codes are returned by the functions provided by this
// API; some are used in higher-level library layers, and others are provided
// merely for alignment with POSIX.
type Errno uint16

const (
	// No error occurred. System call completed successfully.
	ErrnoSuccess Errno = iota
	// Argument list too long.
	Errno2big
	// Permission denied.
	ErrnoAcces
	// Address in use.
	ErrnoAddrinuse
	// Address not available.
	ErrnoAddrnotavail
	// Address family not supported.
	ErrnoAfnosupport
	// Resource unavailable, or operation would block.
	ErrnoAgain
	// Connection already in progress.
	ErrnoAlready
	// Bad file descriptor.
	ErrnoBadf
	// Bad message.
	ErrnoBadmsg
	// Device or resource busy.
	ErrnoBusy
	// Operation canceled.
	ErrnoCanceled
	// No child processes.
	ErrnoChild
	// Connection aborted.
	ErrnoConnaborted
	// Connection refused.
	ErrnoConnrefused
	// Connection reset.
	ErrnoConnreset
	// Resource deadlock would occur.
	ErrnoDeadlk
	// Destination address required.
	ErrnoDestaddrreq
	// Mathematics argument out of domain of function.
	ErrnoDom
	// Reserved.
	ErrnoDquot
	// File exists.
	ErrnoExist
	// Bad address.
	ErrnoFault
	// File too large.
	ErrnoFbig
	// Host is unreachable.
	ErrnoHostunreach
	// Identifier removed.
	ErrnoIdrm
	// Illegal byte sequence.
	ErrnoIlseq
	// Operation in progress.
	ErrnoInprogress
	// Interrupted function.
	ErrnoIntr
	// Invalid argument.
	ErrnoInval
	// I/O error.
	ErrnoIo
	// Socket is connected.
	ErrnoIsconn
	// Is a directory.
	ErrnoIsdir
	// Too many levels of symbolic links.
	ErrnoLoop
	// File descriptor value too large.
	ErrnoMfile
	// Too many links.
	ErrnoMlink
	// Message too large.
	ErrnoMsgsize
	// Reserved.
	ErrnoMultihop
	// Filename too long.
	ErrnoNametoolong
	// Network is down.
	ErrnoNetdown
	// Connection aborted by network.
	ErrnoNetreset
	// Network unreachable.
	ErrnoNetunreach
	// Too many files open in system.
	ErrnoNfile
	// No buffer space available.
	ErrnoNobufs
	// No such device.
	ErrnoNodev
	// No such file or directory.
	ErrnoNoent
	// Executable file format error.
	ErrnoNoexec
	// No locks available.
	ErrnoNolck
	// Reserved.
	ErrnoNolink
	// Not enough space.
	ErrnoNomem
	// No message of the desired type.
	ErrnoNomsg
	// Protocol not available.
	ErrnoNoprotoopt
	// No space left on device.
	ErrnoNospc
	// Function not supported.
	ErrnoNosys
	// The socket is not connected.
	ErrnoNotconn
	// Not a directory or a symbolic link to a directory.
	ErrnoNotdir
	// Directory not empty.
	ErrnoNotempty
	// State not recoverable.
	ErrnoNotrecoverable
	// Not a socket.
	ErrnoNotsock
	// Not supported, or operation not supported on socket.
	ErrnoNotsup
	// Inappropriate I/O control operation.
	ErrnoNotty
	// No such device or address.
	ErrnoNxio
	// Value too large to be stored in data type.
	ErrnoOverflow
	// Previous owner died.
	ErrnoOwnerdead
	// Operation not permitted.
	ErrnoPerm
	// Broken pipe.
	ErrnoPipe
	// Protocol error.
	ErrnoProto
	// Protocol not supported.
	ErrnoProtonosupport
	// Protocol wrong type for socket.
	ErrnoPrototype
	// Result too large.
	ErrnoRange
	// Read-only file system.
	ErrnoRofs
	// Invalid seek.
	ErrnoSpipe
	// No such process.
	ErrnoSrch
	// Reserved.
	ErrnoStale
	// Connection timed out.
	ErrnoTimedout
	// Text file busy.
	ErrnoTxtbsy
	// Cross-device link.
	ErrnoXdev
	// Extension: Capabilities insufficient.
	ErrnoNotcapable
)

var errnoNames = [...]string{
    "success",
    "2big",
    "acces",
    "addrinuse",
    "addrnotavail",
    "afnosupport",
    "again",
    "already",
    "badf",
    "badmsg",
    "busy",
    "canceled",
    "child",
    "connaborted",
    "connrefused",
    "connreset",
    "deadlk",
    "destaddrreq",
    "dom",
    "dquot",
    "exist",
    "fault",
    "fbig",
    "hostunreach",
    "idrm",
    "ilseq",
    "inprogress",
    "intr",
    "inval",
    "io",
    "isconn",
    "isdir",
    "loop",
    "mfile",
    "mlink",
    "msgsize",
    "multihop",
    "nametoolong",
    "netdown",
    "netreset",
    "netunreach",
    "nfile",
    "nobufs",
    "nodev",
    "noent",
    "noexec",
    "nolck",
    "nolink",
    "nomem",
    "nomsg",
    "noprotoopt",
    "nospc",
    "nosys",
    "notconn",
    "notdir",
    "notempty",
    "notrecoverable",
    "notsock",
    "notsup",
    "notty",
    "nxio",
    "overflow",
    "ownerdead",
    "perm",
    "pipe",
    "proto",
    "protonosupport",
    "prototype",
    "range",
    "rofs",
    "spipe",
    "srch",
    "stale",
    "timedout",
    "txtbsy",
    "xdev",
    "notcapable",
}

func (e Errno) String() string {
	return errnoNames[e]
}

type Rights uint64

const (
	// The right to invoke `fd_datasync`.
	//
	// If `rights::path_open` is set, includes the right to invoke
	// `path_open` with `fdflags::dsync`.
	RightsFdDatasync Rights = 1 << iota
	// The right to invoke `fd_read` and `sock_recv`.
	//
	// If `rights::fd_seek` is set, includes the right to invoke `fd_pread`.
	RightsFdRead
	// The right to invoke `fd_seek`. This flag implies `rights::fd_tell`.
	RightsFdSeek
	// The right to invoke `fd_fdstat_set_flags`.
	RightsFdFdstatSetFlags
	// The right to invoke `fd_sync`.
	//
	// If `rights::path_open` is set, includes the right to invoke
	// `path_open` with `fdflags::rsync` and `fdflags::dsync`.
	RightsFdSync
	// The right to invoke `fd_seek` in such a way that the file offset
	// remains unaltered (i.e., `whence::cur` with offset zero), or to
	// invoke `fd_tell`.
	RightsFdTell
	// The right to invoke `fd_write` and `sock_send`.
	// If `rights::fd_seek` is set, includes the right to invoke `fd_pwrite`.
	RightsFdWrite
	// The right to invoke `fd_advise`.
	RightsFdAdvise
	// The right to invoke `fd_allocate`.
	RightsFdAllocate
	// The right to invoke `path_create_directory`.
	RightsPathCreateDirectory
	// If `rights::path_open` is set, the right to invoke `path_open` with `oflags::creat`.
	RightsPathCreateFile
	// The right to invoke `path_link` with the file descriptor as the
	// source directory.
	RightsPathLinkSource
	// The right to invoke `path_link` with the file descriptor as the
	// target directory.
	RightsPathLinkTarget
	// The right to invoke `path_open`.
	RightsPathOpen
	// The right to invoke `fd_readdir`.
	RightsFdReaddir
	// The right to invoke `path_readlink`.
	RightsPathReadlink
	// The right to invoke `path_rename` with the file descriptor as the source directory.
	RightsPathRenameSource
	// The right to invoke `path_rename` with the file descriptor as the target directory.
	RightsPathRenameTarget
	// The right to invoke `path_filestat_get`.
	RightsPathFilestatGet
	// The right to change a file's size (there is no `path_filestat_set_size`).
	// If `rights::path_open` is set, includes the right to invoke `path_open` with `oflags::trunc`.
	RightsPathFilestatSetSize
	// The right to invoke `path_filestat_set_times`.
	RightsPathFilestatSetTimes
	// The right to invoke `fd_filestat_get`.
	RightsFdFilestatGet
	// The right to invoke `fd_filestat_set_size`.
	RightsFdFilestatSetSize
	// The right to invoke `fd_filestat_set_times`.
	RightsFdFilestatSetTimes
	// The right to invoke `path_symlink`.
	RightsPathSymlink
	// The right to invoke `path_remove_directory`.
	RightsPathRemoveDirectory
	// The right to invoke `path_unlink_file`.
	RightsPathUnlinkFile
	// If `rights::fd_read` is set, includes the right to invoke `poll_oneoff` to subscribe to `eventtype::fd_read`.
	// If `rights::fd_write` is set, includes the right to invoke `poll_oneoff` to subscribe to `eventtype::fd_write`.
	RightsPollFdReadwrite
	// The right to invoke `sock_shutdown`.
	RightsSockShutdown
)

// A file descriptor handle.
type Fd w.Handle

// A region of memory for scatter/gather reads.
type Iovec struct {
	_ structs.HostLayout
	// The address of the buffer to be filled.
	Buf w.Pointer32[uint8]
	// The length of the buffer to be filled.
	BufLen Size
}

// A region of memory for scatter/gather writes.
type Ciovec struct {
	_ structs.HostLayout
	// The address of the buffer to be written.
	Buf w.Pointer32[uint8]
	// The length of the buffer to be written.
	BufLen Size
}

type IovecArray w.List[Iovec]
type CiovecArray w.List[Ciovec]

// Relative offset within a file.
type Filedelta int64

// The position relative to which to set the offset of the file descriptor.
type Whence uint8

const (
	// Seek relative to current position.
	WhenceCur Whence = iota
	// Seek relative to end-of-file.
	WhenceEnd
	// Seek relative to start-of-file.
	WhenceSet
)

var whenceNames = [...]string{"cur", "end", "set"}

func (w Whence) String() string {
	return whenceNames[w]
}

// A reference to the offset of a directory entry.
type Dircookie uint64

// The type for the `dirent::d_namlen` field of `dirent` struct.
type Dirnamelen uint32

// File serial number that is unique within its file system.
type Inode uint64

// The type of a file descriptor or file.
type Filetype uint8

const (
	// The type of the file descriptor or file is unknown or is different from any of the other types specified.
	FiletypeUnknown Filetype = iota
	// The file descriptor or file refers to a block device inode.
	FiletypeBlockDevice
	// The file descriptor or file refers to a character device inode.
	FiletypeCharacterDevice
	// The file descriptor or file refers to a directory inode.
	FiletypeDirectory
	// The file descriptor or file refers to a regular file inode.
	FiletypeRegularFile
	// The file descriptor or file refers to a datagram socket.
	FiletypeSocketDgram
	// The file descriptor or file refers to a byte-stream socket.
	FiletypeSocketStream
	// The file refers to a symbolic link inode.
	FiletypeSymbolicLink
)

var filetypeNames = [...]string{
	"unknown",
	"block_device",
	"character_device",
	"directory",
	"regular_file",
	"socket_dgram",
	"socket_stream",
	"symbolic_link",
}

func (f Filetype) String() string {
	return filetypeNames[f]
}

// A directory entry.
type Dirent struct {
	_ structs.HostLayout
	// The offset of the next directory entry stored in this directory.
	DNext Dircookie
	// The serial number of the file referred to by this directory entry.
	DIno Inode
	// The length of the name of the directory entry.
	DNamlen Dirnamelen
	// The type of the file referred to by this directory entry.
	DType Filetype
}

// File or memory access pattern advisory information.
type Advice uint8

const (
	// The application has no advice to give on its behavior with respect to the specified data.
	AdviceNormal Advice = iota
	// The application expects to access the specified data sequentially from lower offsets to higher offsets.
	AdviceSequential
	// The application expects to access the specified data in a random order.
	AdviceRandom
	// The application expects to access the specified data in the near future.
	AdviceWillneed
	// The application expects that it will not access the specified data in the near future.
	AdviceDontneed
	// The application expects to access the specified data once and then not reuse it thereafter.
	AdviceNoreuse
)

var adviceNames = [...]string{
	"normal",
	"sequential",
	"random",
	"willneed",
	"dontneed",
	"noreuse",
}

func (a Advice) String() string {
	return adviceNames[a]
}

// File descriptor flags.
type Fdflags uint16

const (
	// Append mode: Data written to the file is always appended to the file's end.
	FdflagsAppend Fdflags = 1 << iota
	// Write according to synchronized I/O data integrity completion. Only the data stored in the file is synchronized.
	FdflagsDsync
	// Non-blocking mode.
	FdflagsNonblock
	// Synchronized read I/O operations.
	FdflagsRsync
	// Write according to synchronized I/O file integrity completion. In
	// addition to synchronizing the data stored in the file, the implementation
	// may also synchronously update the file's metadata.
	FdflagsSync
)

// File descriptor attributes.
type Fdstat struct {
	_ structs.HostLayout
	// File type.
	FsFiletype Filetype
	// File descriptor flags.
	FsFlags Fdflags
	// Rights that apply to this file descriptor.
	FsRightsBase Rights
	// Maximum set of rights that may be installed on new file descriptors that
	// are created through this file descriptor, e.g., through `path_open`.
	FsRightsInheriting Rights
}

// Identifier for a device containing a file system. Can be used in combination
// with `inode` to uniquely identify a file or directory in the filesystem.
type Device uint64

// Which file time attributes to adjust.
type Fstflags uint16

const (
	// Adjust the last data access timestamp to the value stored in `filestat::atim`.
	FstflagsAtim Fstflags = 1 << iota
	// Adjust the last data access timestamp to the time of clock `clockid::realtime`.
	FstflagsAtimNow
	// Adjust the last data modification timestamp to the value stored in `filestat::mtim`.
	FstflagsMtim
	// Adjust the last data modification timestamp to the time of clock `clockid::realtime`.
	FstflagsMtimNow
)

// Flags determining the method of how paths are resolved.
type Lookupflags uint32

const (
	// As long as the resolved path corresponds to a symbolic link, it is expanded.
	LookupflagsSymlinkFollow Lookupflags = 1 << iota
)

// Open flags used by `path_open`.
type Oflags uint16

const (
	// Create file if it does not exist.
	OflagsCreat Oflags = 1 << iota
	// Fail if not a directory.
	OflagsDirectory
	// Fail if file already exists.
	OflagsExcl
	// Truncate file to size 0.
	OflagsTrunc
)

// Number of hard links to an inode.
type Linkcount uint32

// File attributes.
type Filestat struct {
	_ structs.HostLayout
	// Device ID of device containing the file.
	Dev Device
	// File serial number.
	Ino Inode
	// File type.
	Filetype Filetype
	// Number of hard links to the file.
	Nlink Linkcount
	// For regular files, the file size in bytes. For symbolic links, the length in bytes of the pathname contained in the symbolic link.
	Size Filesize
	// Last data access timestamp.
	Atim Timestamp
	// Last data modification timestamp.
	Mtim Timestamp
	// Last file status change timestamp.
	Ctim Timestamp
}

// User-provided value that may be attached to objects that is retained when
// extracted from the implementation.
type Userdata uint64

// Type of a subscription to an event or its occurrence.
type Eventtype uint8

const (
	// The time value of clock `subscription_clock::id` has
	// reached timestamp `subscription_clock::timeout`.
	EventtypeClock Eventtype = iota
	// File descriptor `subscription_fd_readwrite::file_descriptor` has data
	// available for reading. This event always triggers for regular files.
	EventtypeFdRead
	// File descriptor `subscription_fd_readwrite::file_descriptor` has capacity
	// available for writing. This event always triggers for regular files.
	EventtypeFdWrite
)

var eventtypeNames = [...]string{"clock", "fd_read", "fd_write"}

func (e Eventtype) String() string {
	return eventtypeNames[e]
}

// The state of the file descriptor subscribed to with
// `eventtype::fd_read` or `eventtype::fd_write`.
type Eventrwflags uint16

const (
	// The peer of this socket has closed or disconnected.
	EventrwflagsFdReadwriteHangup Eventrwflags = 1 << iota
)

// The contents of an `event` for the `eventtype::fd_read` and
// `eventtype::fd_write` variants
type EventFdReadwrite struct {
	_ structs.HostLayout
	// The number of bytes available for reading or writing.
	Nbytes Filesize
	// The state of the file descriptor.
	Flags Eventrwflags
}

// An event that occurred.
type Event struct {
	_ structs.HostLayout
	// User-provided value that got attached to `subscription::userdata`.
	Userdata Userdata
	// If non-zero, an error that occurred while processing the subscription request.
	Error Errno
	// The type of event that occured
	Type Eventtype
	// The contents of the event, if it is an `eventtype::fd_read` or
	// `eventtype::fd_write`. `eventtype::clock` events ignore this field.
	FdReadwrite EventFdReadwrite
}

// Flags determining how to interpret the timestamp provided in
// `subscription_clock::timeout`.
type Subclockflags uint16

const (
	// If set, treat the timestamp provided in
	// `subscription_clock::timeout` as an absolute timestamp of clock
	// `subscription_clock::id`. If clear, treat the timestamp
	// provided in `subscription_clock::timeout` relative to the
	// current time value of clock `subscription_clock::id`.
	SubclockflagsSubscriptionClockAbstime Subclockflags = 1 << iota
)

// The contents of a `subscription` when type is `eventtype::clock`.
type SubscriptionClock struct {
	_ structs.HostLayout
	// The user-defined unique identifier of the clock.
	Identifier Userdata
	// The clock against which to compare the timestamp.
	Id Clockid
	// The absolute or relative timestamp.
	Timeout Timestamp
	// The amount of time that the implementation may wait additionally
	// to coalesce with other events.
	Precision Timestamp
	// Flags specifying whether the timeout is absolute or relative
	Flags Subclockflags
}

// The contents of a `subscription` when the variant is
// `eventtype::fd_read` or `eventtype::fd_write`.
type SubscriptionFdReadwrite struct {
	_ structs.HostLayout
	// The file descriptor on which to wait for it to become ready for reading or writing.
	FileDescriptor Fd
}

// The contents of a `subscription`.
type SubscriptionU w.Union[Eventtype, SubscriptionClock, uint64]

func (s *SubscriptionU) SubscriptionClock() bool {
	return s.Tag() == 0
}
func (s *SubscriptionU) SubscriptionFdReadwrite() bool {
	return s.Tag() == 1
}
func (s *SubscriptionU) SubscriptionFdReadwrite2() bool {
	return s.Tag() == 2
}

// Subscription to an event.
type Subscription struct {
	_ structs.HostLayout
	// User-provided value that is attached to the subscription in the
	// implementation and returned through `event::userdata`.
	Userdata Userdata
	// The type of the event to which to subscribe.
	U SubscriptionU
}

// Exit code generated by a process when exiting.
type Exitcode uint32

// Signal condition.
type Signal uint8

const (
	// No signal. Note that POSIX has special semantics for `kill(pid, 0)`,
	// so this value is reserved.
	SignalNone Signal = iota
	// Hangup.
	// Action: Terminates the process.
	SignalHup
	// Terminate interrupt signal.
	// Action: Terminates the process.
	SignalInt
	// Terminal quit signal.
	// Action: Terminates the process.
	SignalQuit
	// Illegal instruction.
	// Action: Terminates the process.
	SignalIll
	// Trace/breakpoint trap.
	// Action: Terminates the process.
	SignalTrap
	// Process abort signal.
	// Action: Terminates the process.
	SignalAbrt
	// Access to an undefined portion of a memory object.
	// Action: Terminates the process.
	SignalBus
	// Erroneous arithmetic operation.
	// Action: Terminates the process.
	SignalFpe
	// Kill.
	// Action: Terminates the process.
	SignalKill
	// User-defined signal 1.
	// Action: Terminates the process.
	SignalUsr1
	// Invalid memory reference.
	// Action: Terminates the process.
	SignalSegv
	// User-defined signal 2.
	// Action: Terminates the process.
	SignalUsr2
	// Write on a pipe with no one to read it.
	// Action: Ignored.
	SignalPipe
	// Alarm clock.
	// Action: Terminates the process.
	SignalAlrm
	// Termination signal.
	// Action: Terminates the process.
	SignalTerm
	// Child process terminated, stopped, or continued.
	// Action: Ignored.
	SignalChld
	// Continue executing, if stopped.
	// Action: Continues executing, if stopped.
	SignalCont
	// Stop executing.
	// Action: Stops executing.
	SignalStop
	// Terminal stop signal.
	// Action: Stops executing.
	SignalTstp
	// Background process attempting read.
	// Action: Stops executing.
	SignalTtin
	// Background process attempting write.
	// Action: Stops executing.
	SignalTtou
	// High bandwidth data is available at a socket.
	// Action: Ignored.
	SignalUrg
	// CPU time limit exceeded.
	// Action: Terminates the process.
	SignalXcpu
	// File size limit exceeded.
	// Action: Terminates the process.
	SignalXfsz
	// Virtual timer expired.
	// Action: Terminates the process.
	SignalVtalrm
	// Profiling timer expired.
	// Action: Terminates the process.
	SignalProf
	// Window changed.
	// Action: Ignored.
	SignalWinch
	// I/O possible.
	// Action: Terminates the process.
	SignalPoll
	// Power failure.
	// Action: Terminates the process.
	SignalPwr
	// Bad system call.
	// Action: Terminates the process.
	SignalSys
)

var signalNames = [...]string{
    "none",
    "hup",
    "int",
    "quit",
    "ill",
    "trap",
    "abrt",
    "bus",
    "fpe",
    "kill",
    "usr1",
    "segv",
    "usr2",
    "pipe",
    "alrm",
    "term",
    "chld",
    "cont",
    "stop",
    "tstp",
    "ttin",
    "ttou",
    "urg",
    "xcpu",
    "xfsz",
    "vtalrm",
    "prof",
    "winch",
    "poll",
    "pwr",
    "sys",
}

func (s Signal) String() string {
	return signalNames[s]
}

// Flags provided to `sock_recv`.
type Riflags uint16

const (
	// Returns the message without removing it from the socket's receive queue.
	RiflagsRecvPeek Riflags = 1 << iota
	// On byte-stream sockets, block until the full amount of data can be returned.
	RiflagsRecvWaitall
)

// Flags returned by `sock_recv`.
type Roflags uint16

const (
	// Returned by `sock_recv`: Message data has been truncated.
	RoflagsRecvDataTruncated Roflags = 1 << iota
)

// Flags provided to `sock_send`. As there are currently no flags
// defined, it must be set to zero.
type Siflags uint16

// Which channels on a socket to shut down.
type Sdflags uint8

const (
	// Disables further receive operations.
	SdflagsRd Sdflags = 1 << iota
	// Disables further send operations.
	SdflagsWr
)

// Identifiers for preopened capabilities.
type Preopentype uint8

const (
	// A pre-opened directory
	PreopentypeDir Preopentype = iota
)

var preopentypeNames = [...]string{"dir"}

func (p Preopentype) String() string {
	return preopentypeNames[p]
}

// The contents of a $prestat when type is `preopentype::dir`.
type PrestatDir struct {
	_ structs.HostLayout
	// Tye length of the directory name for use with `fd_prestat_dir_name`.
	PrNameLen Size
}

// Information about a pre-opened capability.
type Prestat w.Union[Preopentype, PrestatDir, uint32]

func (p *Prestat) PrestatDir() bool {
	return p.Tag() == 0
}
