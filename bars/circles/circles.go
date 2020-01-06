package circles

import (
	color "github.com/multiverse-os/color"
	loading "github.com/multiverse-os/loading"
)

var Style = loading.BarStyle{
	Fill:     []string{color.Lime("⚫")},
	Unfilled: color.Gray("⚪"),
	Format:   " %s%s %0.2f%% %s",
}
