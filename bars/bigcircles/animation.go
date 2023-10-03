package bigcircles

import (
	color "github.com/multiverse-os/ansi/color"
)

// ⬤◯◯◯◯◯◯◯◯◯ 9%
// TODO: Should have half filled for 1.5/100 for more complex bars for example
// installation, transfer/download
var Animation = map[string][]string{
	"filling":  []string{color.Silver("⬤ "), color.White("⬤ ")},
	"unfilled": []string{color.Gray("◯ ")},
}
