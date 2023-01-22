package main

import (
	"fmt"
	"math/rand"
	"time"

	color "github.com/multiverse-os/ansi/color"
	loading "github.com/multiverse-os/loading"
	bigcircles "github.com/multiverse-os/loading/bars/bigcircles"
	blocks "github.com/multiverse-os/loading/bars/blocks"
	circles "github.com/multiverse-os/loading/bars/circles"
	dots "github.com/multiverse-os/loading/bars/dots"
	rectangles "github.com/multiverse-os/loading/bars/rectangles"
	squares "github.com/multiverse-os/loading/bars/squares"
	thinblocks "github.com/multiverse-os/loading/bars/thinblocks"
)

func main() {
	fmt.Println("Loading Bar Example")
	fmt.Println("===================")

	fmt.Println("Running 'bigcircles' loading bar example:")
	RunBarExample(bigcircles.Style)

	fmt.Println("Running 'thinblocks' loading bar example:")
	RunBarExample(thinblocks.Style)

	fmt.Println("Running 'dots' loading bar example:")
	RunBarExample(dots.Style)

	fmt.Println("Running 'rectangles' loading bar example:")
	RunBarExample(rectangles.Style)

	fmt.Println("Running 'blocks' loading bar example:")
	RunBarExample(blocks.Style)

	fmt.Println("Running 'circle' loading bar example:")
	RunBarExample(circles.Style)

	fmt.Println("Running 'squares' loading bar example:")
	RunBarExample(squares.Style)

}

func RunBarExample(style loading.BarAnimation) {
	loadingBar := loading.Bar().Width(80).Style(style).Start()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(rand.Intn(135)+22) * time.Millisecond)
		if loadingBar.Increment(1) {
			break
		}
	}
	loadingBar.Status(color.Green("Completed!")).Complete()
}
