package thinblocks

import (
	color "github.com/multiverse-os/ansi/color"
)

// ▉▉▋            41%
var Animation = map[string][]string{
	"fill":     []string{"▏", "▏", "▎", "▎", "▍", "▍", "▌", "▌", "▋", "▋", "▊", "▊", "▉", "▉"},
	"unfilled": []string{color.Gray(" ")},
}
