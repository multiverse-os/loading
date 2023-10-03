package dots

import (
	color "github.com/multiverse-os/ansi/color"
)

// ⣿⣿⣿⣿⡟⣀⣀⣀⣀⣀⣀ 41%
var Animation = map[string][]string{
	"fill": []string{
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
	"unfilled": []string{color.Black("⣀")},
}
