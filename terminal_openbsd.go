// +build openbsd,!cgo

package terminal

import "syscall"

const ioctlReadTermios = syscall.TIOCGETA
const ioctlWriteTermios = syscall.TIOCSETA
