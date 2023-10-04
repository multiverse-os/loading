package loading

import (
	"fmt"
)

// /////////////////////////////////////////////////////////////////////////////
// Attribute defines a single SGR Code
type sgr int

// Base attributes
const (
	Reset sgr = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

const (
	Black sgr = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	//Normal sgr = 0
	Bright sgr = 60
)

const (
	Foreground sgr = 30
	Background sgr = 40
	//HiForeground sgr = 90
	//HiBackground sgr = 100
)

// /////////////////////////////////////////////////////////////////////////////
// NOTE
// This enables us to do, and it works!
//
//	ANSI(Black, Bright, Background) => 0+60+40=100
func ANSI(attrs ...sgr) (sum sgr) {
	for _, attr := range attrs {
		sum += attr
	}
	return sum
}

func Color(cName, cType sgr) sgr {
	return cName + cType
}

func (attr sgr) Text(text string) string {
	return fmt.Sprintf("%v%s%v", attr, text, attr)
}

///////////////////////////////////////////////////////////////////////////////

const escape = "\x1b"
