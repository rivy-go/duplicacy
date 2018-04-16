package xattr

import (
	"syscall"
	"unsafe"
)

func getxattr(path string, name string, value *byte, size int) (int, error) {
	if r0, _, e1 := syscall.Syscall6(syscall.SYS_GETXATTR,
		uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
		uintptr(unsafe.Pointer(syscall.StringBytePtr(name))),
		uintptr(unsafe.Pointer(value)), uintptr(size), 0, 0); e1 != syscall.Errno(0) {
		if e1 == syscall.ENODATA {
			return 0, e1
		}
		return -1, e1
	} else {
		return int(r0), nil
	}
}

func listxattr(path string, namebuf *byte, size int) (int, error) {
	if r0, _, e1 := syscall.Syscall(syscall.SYS_LISTXATTR,
		uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
		uintptr(unsafe.Pointer(namebuf)),
		uintptr(size)); e1 != syscall.Errno(0) {
		return -1, e1
	} else {
		return int(r0), nil
	}
}

func setxattr(path string, name string, value *byte, size int) (err error) {
	_, _, e1 := syscall.Syscall6(syscall.SYS_SETXATTR,
		uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
		uintptr(unsafe.Pointer(syscall.StringBytePtr(name))),
		uintptr(unsafe.Pointer(value)), uintptr(size), 0, 0)
	return e1
}

func removexattr(path string, name string) (err error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_REMOVEXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(syscall.StringBytePtr(name))), 0)
	return e1
}
