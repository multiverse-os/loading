package squares

import (
	color "github.com/multiverse-os/color"
	loading "github.com/multiverse-os/loading"
)

// ■□□□□□□□□□□□ 9%
var Style = loading.BarStyle{
	Fill:     []string{color.White("■")},
	Unfilled: color.Gray("□"),
	Format:   " %s%s %0.2f%% %s",
}
