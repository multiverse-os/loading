package main

import (
	"fmt"
	"math/rand"
	"time"

	color "github.com/multiverse-os/ansi/color"

	loading "github.com/multiverse-os/loading"
	bigcircles "github.com/multiverse-os/loading/bars/bigcircles"
	circles "github.com/multiverse-os/loading/bars/circles"
	dots "github.com/multiverse-os/loading/bars/dots"
)

func randomWait() {
	time.Sleep(time.Duration(rand.Intn(50)+20) * time.Millisecond)
}

func main() {
	fmt.Printf("Loading Bar Example\n")
	fmt.Printf("===================\n")

	fmt.Printf("Running 'circles' loading bar example:\n")

	RunBarExampleWithPercent(dots.Animation)
	RunBarExampleWithoutPercent(circles.Animation)
	RunBarExampleWithPercent(bigcircles.Animation)
}

func RunBarExampleWithoutPercent(animation map[string][]string) {
	loadingBar := loading.NewBar(animation)
	fmt.Printf("loadingBar(%v)\n", loadingBar)
	loadingBar.ShowPercent(false)
	loadingBar.Start()

	for 0 < loadingBar.RemainingTicks() {
		randomWait()
		loadingBar.Increment(1.5)
	}

	loadingBar.Status(color.Green("Completed!")).End()
}

func RunBarExampleWithPercent(animation map[string][]string) {
	loadingBar := loading.NewBar(animation)
	fmt.Printf("loadingBar(%v)\n", loadingBar)
	loadingBar.Start()

	for 0 < loadingBar.RemainingTicks() {
		randomWait()
		loadingBar.Increment(1.5)
	}

	loadingBar.Status(color.Green("Completed!")).End()
}