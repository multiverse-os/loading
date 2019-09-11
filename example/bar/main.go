package main

import (
	"fmt"
	"math/rand"
	"time"

	color "github.com/multiverse-os/cli/text/ansi/color"
	loading "github.com/multiverse-os/cli/text/loading"
	bigcircles "github.com/multiverse-os/cli/text/loading/bars/bigcircles"
	blocks "github.com/multiverse-os/cli/text/loading/bars/blocks"
	circles "github.com/multiverse-os/cli/text/loading/bars/circles"
	dots "github.com/multiverse-os/cli/text/loading/bars/dots"
	rectangles "github.com/multiverse-os/cli/text/loading/bars/rectangles"
	squares "github.com/multiverse-os/cli/text/loading/bars/squares"
	thinblocks "github.com/multiverse-os/cli/text/loading/bars/thinblocks"
)

func main() {
	fmt.Println("Loading Bar Example")
	fmt.Println("===================")

	fmt.Println("Running 'bigcircles' loading bar example:")
	RunLoadingBarExample(bigcircles.Style)

	fmt.Println("Running 'thinblocks' loading bar example:")
	RunLoadingBarExample(thinblocks.Style)

	fmt.Println("Running 'dots' loading bar example:")
	RunLoadingBarExample(dots.Style)

	fmt.Println("Running 'rectangles' loading bar example:")
	RunLoadingBarExample(rectangles.Style)

	fmt.Println("Running 'blocks' loading bar example:")
	RunLoadingBarExample(blocks.Style)

	fmt.Println("Running 'circle' loading bar example:")
	RunLoadingBarExample(circles.Style)

	fmt.Println("Running 'squares' loading bar example:")
	RunLoadingBarExample(squares.Style)

}

func RunLoadingBarExample(style loading.BarStyle) {
	loadingBar := loading.Bar().Width(80).Style(style).Start()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(rand.Intn(135)+22) * time.Millisecond)
		if loadingBar.Increment(1) {
			break
		}
	}
	loadingBar.Status(color.Green("Completed!")).Complete()
}
