package loading

import (
	"fmt"
)

// /////////////////////////////////////////////////////////////////////////////
// Attribute defines a single SGR Code
type sgr uint8

// Base attributes
const (
	//Reset     sgr = 0
	Bold      sgr = 1
	Faint     sgr = 2
	Italic    sgr = 3
	Underline sgr = 4
	Blink     sgr = 5
	//BlinkRapid sgr = 6
	Reverse sgr = 7
	//Concealed sgr = 8
	CrossedOut sgr = 9
	Overline   sgr = 53
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

// ////////////////////////////////////////////////////////////////////////////
// ANSI(Black, Bright, Background) => 0+60+40=100
func style(attrs ...sgr) (sum sgr) {
	for _, attr := range attrs {
		sum += attr
	}
	return sum
}

func color(cName, cType sgr) sgr { return cName + cType }

///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// NOTE
// This enables us to do, and it works!

type text struct {
	string
	attrs []sgr
}

func Text(str string) *text {
	return &text{
		string: str,
		//attrs:  []sgr{},
	}
}

//func (txt *text) attribute(attr sgr) *text {
//	txt.attrs = append(txt.attrs, attr)
//	return txt
//}

func (txt *text) sgr(attrs ...sgr) *text {
	//for _, attr := range txt.attrs {
	//txt.attribute(attr)
	txt.attrs = append(txt.attrs, attrs...)
	//}
	return txt
}

func (txt text) profile() profile {
	if len(txt.attrs) == 0 {
		return ASCII
	} else {
		return ANSI
	}
}

func (txt text) ansi() (str string) {
	// TODO: SEMICOLON MUST ONLY BE ADDED
	// IFFFFFF we are insertintg more than 1 ansi style, otherwise NOEN
	for index, attr := range txt.attrs {
		str += fmt.Sprintf("%v", attr)
		if index != 0 {
			str += ";"
		}
	}
	return str
}

func (txt text) String() string {
	if txt.profile() == ASCII {
		return txt.string
	}
	// TODO: This uses universal reset we want specific reset
	return fmt.Sprintf("%s%sm%s%sm", CSI, txt.ansi(), txt.string, CSI+RESET)
}

///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////

type profile uint8

const (
	ASCII profile = iota
	ANSI
)

//
//func (p profile) String() string {
//	switch p {
//	case ANSI:
//		return "ansi"
//	default:
//		return "ascii"
//	}
//}

const (
	RESET = "0"
	// Escape character
	//     \x1b[30m
	ESC = "\x1b"
	// Bell
	BEL = "\a"
	// Control Sequence Introducer
	CSI = ESC + "["
	// Operating System Command
	OSC = ESC + "]"
	// String Terminator
	ST = ESC + `\`
)

///////////////////////////////////////////////////////////////////////////////
