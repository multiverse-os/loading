package dots

import (
	color "github.com/multiverse-os/cli/text/ansi/color"
	loading "github.com/multiverse-os/cli/text/loading"
)

// ⣿⣿⣿⣿⡟⣀⣀⣀⣀⣀⣀ 41%
var Style = loading.BarStyle{
	Fill: []string{
		color.Gray("⡀"),
		color.Gray("⡄"),
		color.Gray("⡆"),
		color.Gray("⡇"),
		color.Gray("⡏"),
		color.Silver("⡟"),
		color.Silver("⡿"),
		color.White("⣿"),
		color.Silver("⡿"),
		color.Silver("⡟"),
		color.Silver("⡟"),
		color.Gray("⡏"),
		color.Gray("⡇"),
		color.Gray("⡆"),
		color.Gray("⡄"),
		color.Gray("⡀"),
		color.Gray("⡀"),
		color.Gray("⡄"),
		color.Gray("⡆"),
		color.Gray("⡇"),
		color.Gray("⡏"),
		color.Silver("⡟"),
		color.Silver("⡿"),
		color.White("⣿"),
	},
	Unfilled: color.Black("⣀"),
	Format:   " %s%s %0.2f%% %s",
}
