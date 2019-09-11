package loading

import (
	"errors"
	"syscall"
	"unsafe"
)

type terminalSize struct {
	Rows        uint16
	Columns     uint16
	PixelWidth  uint16
	PixelHeight uint16
}

func TerminalWidth() (int, error) {
	tSize := &terminalSize{}
	returnCode, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(tSize)),
	)
	if int(returnCode) == -1 {
		return 0, errors.New("[error] failed to determine terminal width: " + err.Error())
	}
	return int(tSize.Columns), nil
}
