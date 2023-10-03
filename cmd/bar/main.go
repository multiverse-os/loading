package main

import (
	"fmt"
	"math/rand"
	"time"

	color "github.com/multiverse-os/ansi/color"

	loading "github.com/multiverse-os/loading"
	circles "github.com/multiverse-os/loading/bars/circles"
)

func randomWait() {
	time.Sleep(time.Duration(rand.Intn(1)+2) * time.Second)
}

func main() {
	fmt.Printf("Loading Bar Example\n")
	fmt.Printf("===================\n")

	fmt.Printf("Running 'circles' loading bar example:\n")

	animation := circles.Animation
	//animation.Format = " %s%s %0.2f% %s%s"
	//animation.Format = "%v%v%v"

	//firstFrame := animation.Fill[0]

	//firstFrameBytes := []rune(firstFrame)[0].Bytes()

	//fmt.Printf("%v len(%v)", firstFrameBytes, len(firstFrameBytes))

	//testFrame := "X"
	//testFrameBytes := rune(testFrame)[0].Bytes()
	//fmt.Printf("%v len(%v)", testFrameBytes, len(testFrameBytes))

	fmt.Sprintf("Format: %v\n\n\n", animation.Format)

	RunBarExample(animation)

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
	loadingBar := loading.NewBar(animation).ShowPercent(false)
	loadingBar.Start()

	for 0 < loadingBar.RemainingTicks() {
		randomWait()
		loadingBar.Increment(1.5)
	}
	loadingBar.Status(color.Green("Completed!")).End()
}
