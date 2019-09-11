package dots

import (
	color "github.com/multiverse-os/cli/text/ansi/color"
	loading "github.com/multiverse-os/cli/text/loading"
)

var Style = loading.BarStyle{
	Fill:     []string{color.Lime("⚫")},
	Unfilled: color.Gray("⚪"),
	Format:   " %s%s %0.2f%% %s",
}
