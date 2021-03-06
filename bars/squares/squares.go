package squares

import (
	color "github.com/multiverse-os/cli/terminal/ansi/color"
	loading "github.com/multiverse-os/cli/terminal/loading"
)

// ■□□□□□□□□□□□ 9%
var Style = loading.BarStyle{
	Fill:     []string{color.White("■")},
	Unfilled: color.Gray("□"),
	Format:   " %s%s %0.2f%% %s",
}
