package blocks

import (
	color "github.com/multiverse-os/color"
	loading "github.com/multiverse-os/loading"
)

// █▓░░░░░░░░░░ 9%
var Style = loading.BarStyle{
	Fill:     []string{color.SkyBlue("▓"), color.SkyBlue("█")},
	Unfilled: color.Gray("░"),
	Format:   " %s%s %0.2f%% %s",
}
