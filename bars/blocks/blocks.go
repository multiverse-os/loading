package dots

import (
	color "github.com/multiverse-os/cli/text/ansi/color"
	loading "github.com/multiverse-os/cli/text/loading"
)

// █▓░░░░░░░░░░ 9%
var Style = loading.BarStyle{
	Fill:     []string{color.SkyBlue("▓"), color.SkyBlue("█")},
	Unfilled: color.Gray("░"),
	Format:   " %s%s %0.2f%% %s",
}
