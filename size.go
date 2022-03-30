package loading

import (
	"syscall"
	"unsafe"
)

type terminalSize struct {
	Rows        uint16
	Columns     uint16
	PixelWidth  uint16
	PixelHeight uint16
}

func TerminalWidth() int {
	tSize := &terminalSize{}
	returnCode, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(tSize)),
	)
	if int(returnCode) == -1 {
		panic("failed to determine terminal width: " + err.Error())
	}
	return int(tSize.Columns)
}
