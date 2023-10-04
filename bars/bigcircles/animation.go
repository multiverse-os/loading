package bigcircles

import (
	color "github.com/multiverse-os/ansi/color"
)

// ⬤◯◯◯◯◯◯◯◯◯ 9%
var Animation = map[string][]string{
	"filling":  []string{color.Silver("⬤ "), color.White("⬤ ")},
	"unfilled": []string{color.Gray("◯ ")},
}
