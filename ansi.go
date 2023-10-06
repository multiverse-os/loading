package loading

import (
	"fmt"
	"strconv"
)

type profile uint8

const (
	ASCII profile = iota
	ANSI
)

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

// TODO: Eventually use our ANSI library potentially

func render(value string) string { return (ESC + value) }

func EraseDisplay(code int) string { return render(strconv.Itoa(code) + "2J") }
func EraseLine(code int) string    { return render(strconv.Itoa(code) + "K") }

func HideCursor() string { return render("?25l") }
func ShowCursor() string { return render("?25h") }

func CursorUp(n int) string       { return render(strconv.Itoa(n) + "A") }
func CursorDown(n int) string     { return render(strconv.Itoa(n) + "B") }
func CursorForward(n int) string  { return render(strconv.Itoa(n) + "C") }
func CursorBackward(n int) string { return render(strconv.Itoa(n) + "D") }

func CursorStart(n int) string { return render(strconv.Itoa(n) + "G") }

func SaveCursorPosition() string    { return render("s") }
func RestoreCursorPosition() string { return render("u") }
func GetCursorPosition() string     { return render("6n") }

// ////////////////////////////////////////////////////////////////////////////
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

// 10	Primary (default) font
// 11–19	Alternative font	Select alternative font n − 10
// 20	Fraktur (Gothic)	Rarely supported
// 21	Doubly underlined; or: not bold	Double-underline per ECMA-48,[5]: 8.3.117  but instead disables bold intensity on several terminals, including in the Linux kernel's console before version 4.17.[44]
// 22	Normal intensity	Neither bold nor faint; color changes where intensity is implemented as such.
// 23	Neither italic, nor blackletter
// 24	Not underlined	Neither singly nor doubly underlined
// 25	Not blinking	Turn blinking off
// 26	Proportional spacing	ITU T.61 and T.416, not known to be used on terminals
// 27	Not reversed
// 28	Reveal	Not concealed
// 29	Not crossed out

// ///// This is simple 3-bit/4-bit color
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

// 39	Default foreground color	Implementation defined (according to standard)
// 49	Default background color	Implementation defined (according to standard)

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

////////////TODO not as important honestly, we can wait until its ANSI
// only lib again
//////////////////////////////////////////////
// 8-bit color (256 colors) is done with 38/48

//As 256-color lookup tables became common on graphic cards, escape sequences were added to select from a pre-defined set of 256 colors:[citation needed]
//
//ESC[38;5;⟨n⟩m Select foreground color      where n is a number from the table below
//ESC[48;5;⟨n⟩m Select background color
//  0-  7:  standard colors (as in ESC [ 30–37 m)
//  8- 15:  high intensity colors (as in ESC [ 90–97 m)
// 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//232-255:  grayscale from dark to light in 24 steps
//Grayscale colors
//232	233	234	235	236	237	238	239	240	241	242	243	244	245	246	247	248	249	250	251	252	253	254	255

////////////TODO LAST
//
//24-bit
//As "true color" graphic cards with 16 to 24 bits of color became common, applications began to support 24-bit colors. Terminal emulators supporting setting 24-bit foreground and background colors with escape sequences include Xterm,[29] KDE's Konsole,[50][51] and iTerm, as well as all libvte based terminals,[52] including GNOME Terminal.[citation needed]
//
//ESC[38;2;⟨r⟩;⟨g⟩;⟨b⟩ m Select RGB foreground color
//ESC[48;2;⟨r⟩;⟨g⟩;⟨b⟩ m Select RGB background color

// ////////////////////////////////////////////////////////////////////////////
// ANSI(Black, Bright, Background) => 0+60+40=100
func style(attrs ...sgr) (sum sgr) {
	// TODO: Here check for below 10, and if its below 10, only allow it, then
	// maybe have color only support values within sensible color ranges

	// so style will be below 10 only doing underline overline bold etc

	for _, attr := range attrs {
		// NOTE
		// No failure on != 0, we auto provide reset
		if attr != 0 {
			sum += attr
		}
	}
	return sum
}

//TODO
// or maybe decoration and color that check for their specific ranges, and they
// feed into style so you can put bold and black bckground by apssing each
// decoration() and color() through style()

// TODO
// and color will only support feasible color numbers

func decoration(dName sgr) sgr {
	if (dName <= 1 && dName <= 9) &&
		dName == 53 {
		return style(dName)
	}
	return 0
}

func color(cName, cType sgr) sgr {
	cStyle := cName + cType
	// Normal Foreground + Background
	if (30 <= cStyle && cStyle <= 50) ||
		(90 <= cStyle && cStyle <= 110) {
		return style(cName, cType)
	}
	return 0
}

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

func (txt *text) sgr(attrs ...sgr) *text {
	txt.attrs = append(txt.attrs, attrs...)
	return txt
}

func (txt *text) Color(cName sgr) *text {
	return txt.sgr(color(cName, Foreground))
}

func (txt *text) Background(cName sgr) *text {
	return txt.sgr(color(cName, Background))
}

func (txt text) profile() profile {
	if len(txt.attrs) == 0 {
		return ASCII
	} else {
		return ANSI
	}
}

func (txt text) ansi() (str string) {
	// NOTE
	// Semicolon is only used when combining multiple ANSI styles
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
