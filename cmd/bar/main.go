package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/multiverse-os/ansi/color"
	loading "github.com/multiverse-os/loading"
	"github.com/multiverse-os/loading/bars/dots"
)

func randomWait() {
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
}

func main() {
	fmt.Printf("Loading Bar Example\n")
	fmt.Printf("===================\n")

	fmt.Printf("Running 'dots' loading bar example:\n")
	RunBarExample(dots.Animation)

	//fmt.Printf("Running 'bigcircles' loading bar example:\n")
	//RunBarExample(bigcircles.Animation)

	//fmt.Printf("Running 'thinblocks' loading bar example:\n")
	//RunBarExample(thinblocks.Animation)

	//fmt.Printf("Running 'rectangles' loading bar example:\n")
	//RunBarExample(rectangles.Animation)

	//fmt.Printf("Running 'blocks' loading bar example:\n")
	//RunBarExample(blocks.Animation)

	//fmt.Printf("Running 'circle' loading bar example:\n")
	//RunBarExample(circles.Animation)

	//fmt.Printf("Running 'squares' loading bar example:\n")
	//RunBarExample(squares.Animation)

}

func RunBarExample(animation loading.BarAnimation) {
	//loadingBar := loading.NewBar(animation)
	loadingBar := loading.NewBar(animation)
	loadingBar.Start()

	for 0 < loadingBar.RemainingTicks() {
		randomWait()
		loadingBar.Increment(2)
	}
	loadingBar.Status(color.Green("Completed!")).End()
}
