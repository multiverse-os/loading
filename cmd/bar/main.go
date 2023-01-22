package main

import (
	"fmt"
	"math/rand"
	"time"

	loading "github.com/multiverse-os/loading"

	bigcircles "github.com/multiverse-os/loading/bars/bigcircles"
	blocks "github.com/multiverse-os/loading/bars/blocks"
	circles "github.com/multiverse-os/loading/bars/circles"
	dots "github.com/multiverse-os/loading/bars/dots"
	rectangles "github.com/multiverse-os/loading/bars/rectangles"
	squares "github.com/multiverse-os/loading/bars/squares"
	thinblocks "github.com/multiverse-os/loading/bars/thinblocks"

	color "github.com/multiverse-os/ansi/color"
)

func main() {
	fmt.Println("Loading Bar Example")
	fmt.Println("===================")

	fmt.Println("Running 'bigcircles' loading bar example:")
	RunBarExample(bigcircles.Animation)

	fmt.Println("Running 'thinblocks' loading bar example:")
	RunBarExample(thinblocks.Animation)

	fmt.Println("Running 'dots' loading bar example:")
	RunBarExample(dots.Animation)

	fmt.Println("Running 'rectangles' loading bar example:")
	RunBarExample(rectangles.Animation)

	fmt.Println("Running 'blocks' loading bar example:")
	RunBarExample(blocks.Animation)

	fmt.Println("Running 'circle' loading bar example:")
	RunBarExample(circles.Animation)

	fmt.Println("Running 'squares' loading bar example:")
	RunBarExample(squares.Animation)

}

func RunBarExample(animation loading.BarAnimation) {
	loadingBar := loading.NewBar(animation)
	loadingBar.Start()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(rand.Intn(135)+22) * time.Millisecond)
		if loadingBar.Increment(1) {
			break
		}
	}
	loadingBar.Status(color.Green("Completed!")).End()
}
