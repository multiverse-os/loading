package dots

import (
	color "github.com/multiverse-os/ansi/color"
	loading "github.com/multiverse-os/loading"
)

// █▓░░░░░░░░░░ 9%
var Animation = loading.BarAnimation{
	Fill:     []string{color.SkyBlue("▓"), color.SkyBlue("█")},
	Unfilled: color.Gray("░"),
	Format:   " %s%s %0.2f%% %s",
}
