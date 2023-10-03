package blocks

import (
	color "github.com/multiverse-os/ansi/color"
)

// █▓░░░░░░░░░░ 9%
var Animation = map[string][]string{
	"filling":  []string{color.SkyBlue("▓"), color.SkyBlue("█")},
	"unfilled": []string{color.Gray("░")},
}
