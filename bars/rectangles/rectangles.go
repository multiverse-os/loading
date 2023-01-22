package rectangles

import (
	color "github.com/multiverse-os/ansi/color"
	loading "github.com/multiverse-os/loading"
)

// ▰▰▰▰▱▱▱▱▱▱ 41%
var Animation = loading.BarAnimation{
	Fill:     []string{color.White("▰")},
	Unfilled: color.Gray("▱"),
	Format:   " %s%s %0.2f%% %s",
}
