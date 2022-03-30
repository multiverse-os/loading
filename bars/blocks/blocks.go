package dots

import (
	color "github.com/multiverse-os/cli/terminal/ansi/color"
	loading "github.com/multiverse-os/cli/terminal/loading"
)

// █▓░░░░░░░░░░ 9%
var Animation = loading.BarAnimation{
	Fill:     []string{color.SkyBlue("▓"), color.SkyBlue("█")},
	Unfilled: color.Gray("░"),
	Format:   " %s%s %0.2f%% %s",
}
