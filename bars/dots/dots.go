package dots

import (
	color "github.com/multiverse-os/ansi/color"
	loading "github.com/multiverse-os/loading"
)

// ⣿⣿⣿⣿⡟⣀⣀⣀⣀⣀⣀ 41%
var Animation = loading.BarAnimation{
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
	RuneWidth: 1,
	Unfilled:  color.Black("⣀"),
	Format:    " %s%s %0.2f %s",
}
