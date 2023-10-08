package main

import (
	"fmt"
	"math/rand"
	"time"

	loading "github.com/multiverse-os/loading"
	bigcircles "github.com/multiverse-os/loading/bars/bigcircles"
	circles "github.com/multiverse-os/loading/bars/circles"
	dots "github.com/multiverse-os/loading/bars/dots"
)

func randomWait() {
	time.Sleep(time.Duration(rand.Intn(1)+3) * time.Second)
}

func main() {
	fmt.Printf("Loading Bar Example\n")
	fmt.Printf("===================\n")
	fmt.Printf("Running 'dots', 'circles', and 'bigcircles' loading bar example:\n")

	RunBarExampleWithPercent(dots.Animation)
	RunBarExampleWithoutPercent(circles.Animation)
	RunBarExampleWithPercent(bigcircles.Animation)
}

func RunBarExampleWithoutPercent(animation []string) {
	loadingBar := loading.NewBar(animation)
	loadingBar.ShowPercent(false)

	loadingBar.Start()
	for 0 < loadingBar.RemainingTicks() {
		randomWait()
		loadingBar.Increment(1)
	}

	loadingBar.Status(loading.Text("Completed!", loading.Green).String()).End()
}

func RunBarExampleWithPercent(animation []string) {
	loadingBar := loading.NewBar(animation)

	loadingBar.Start()
	for 0 < loadingBar.RemainingTicks() {
		randomWait()
		loadingBar.Increment(1.5)
	}

	loadingBar.Status(loading.Text("Completed!").Color(loading.Green).String()).End()
}
