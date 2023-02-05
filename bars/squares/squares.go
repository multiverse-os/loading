package squares

import (
	color "github.com/multiverse-os/ansi/color"
	loading "github.com/multiverse-os/loading"
)

// ■□□□□□□□□□□□ 9%
var Animation = loading.BarAnimation{
	Fill:      []string{color.White("■")},
	Unfilled:  color.Gray("□"),
	RuneWidth: 1,
	Format:    " %s%s %0.2f%% %s",
}
