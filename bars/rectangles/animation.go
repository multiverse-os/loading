package rectangles

import (
	color "github.com/multiverse-os/ansi/color"
)

// ▰▰▰▰▱▱▱▱▱▱ 41%
var Animation = map[string][]string{
	"fill":     []string{color.White("▰")},
	"unfilled": []string{color.Gray("▱")},
}
