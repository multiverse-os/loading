package loading

import (
	"strconv"
)

const (
	escPrefix = "\u001B["
)

// TODO: Eventually use our ANSI library potentially

func render(value string) string { return (escPrefix + value) }

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
