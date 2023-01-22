package dots

import (
	color "github.com/multiverse-os/ansi/color"
	loading "github.com/multiverse-os/loading"
)

var Animation = loading.BarAnimation{
	Fill:     []string{color.Lime("⚫")},
	Unfilled: color.Gray("⚪"),
	Format:   " %s%s %0.2f%% %s",
}
