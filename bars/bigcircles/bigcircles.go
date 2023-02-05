package bigcircles

import (
	color "github.com/multiverse-os/ansi/color"
	loading "github.com/multiverse-os/loading"
)

// ⬤◯◯◯◯◯◯◯◯◯ 9%
// TODO: Should have half filled for 1.5/100 for more complex bars for example
// installation, transfer/download
var Animation = loading.BarAnimation{
	Fill:      []string{color.Silver("⬤ "), color.White("⬤ ")},
	Unfilled:  color.Gray("◯ "),
	RuneWidth: 2,
	Format:    " %s%s %0.2f %s",
}
