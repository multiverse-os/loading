package loading

import (
	width "golang.org/x/text/width"
)

func RuneWidth(runeString string) uint {
	runeProperties, _ := width.Lookup([]byte(runeString))
	switch runeProperties.Kind() {
	case width.EastAsianWide, width.EastAsianFullwidth:
		return 2
	case width.EastAsianAmbiguous, width.EastAsianNarrow, width.EastAsianHalfwidth:
		return 1
	default: // Neutral
		return 1
	}
}
