package main

import (
	"fmt"
	"time"

	"github.com/multiverse-os/color"
	"github.com/multiverse-os/loading"
	"github.com/multiverse-os/loading/spinners/triangle"
)

const WAIT = 5

func main() {
	fmt.Println("Rainbow Spinner Example")
	fmt.Println("==============")
	rainbow := makeRainbow()
	rainbowDots := loading.Spinner(triangle.Animation).Message("Loading...").Speed(loading.Normal).Palette(rainbow).Start()
	time.Sleep(WAIT * time.Second)
	rainbowDots.Complete(color.Green("Loading Complete!"))

	fmt.Println("Custom Spinner Example")
	fmt.Println("==============")
	customSpinner()
}

func makeRainbow() []string {
	rainbow := []string{}
	// see https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
	ansiRainbow := []int{91, 208, 93, 32, 94, 34, 56}

	for _, code := range ansiRainbow {
		rainbow = append(rainbow, color.Sequence(code))
	}
	return rainbow
}

func customSpinner() {
	chars := []string{
		color.OliveBg(color.Black("♠")),
		color.OliveBg(color.Red("♥")),
		color.OliveBg(color.Green("♣")),
		color.OliveBg(color.Blue("♦")),
	}

	suits := loading.Spinner(chars).Message("Loading...").Speed(loading.Slow).Start()
	time.Sleep(WAIT * time.Second)
	suits.Complete(color.Green("Loading Complete!"))
}
