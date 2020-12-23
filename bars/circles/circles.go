package dots

import (
	color "github.com/multiverse-os/cli/terminal/ansi/color"
	loading "github.com/multiverse-os/cli/terminal/loading"
)

var Style = loading.BarStyle{
	Fill:     []string{color.Lime("⚫")},
	Unfilled: color.Gray("⚪"),
	Format:   " %s%s %0.2f%% %s",
}
