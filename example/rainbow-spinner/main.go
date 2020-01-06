package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/multiverse-os/color"
	"github.com/multiverse-os/loading"
	"github.com/multiverse-os/loading/spinners/triangle"
)

func main() {
	fmt.Println("Rainbow Spinner Example")
	fmt.Println("==============")
	rainbow := makeRainbow()
	rainbowDots := loading.Spinner(triangle.Animation).Message("Loading...").Speed(loading.Normal).Palette(rainbow).Start()
	randomWait()
	rainbowDots.Complete(color.Green("Loading Complete!"))
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

func randomWait() {
	time.Sleep(time.Duration(rand.Intn(2)+2) * time.Second)
}
