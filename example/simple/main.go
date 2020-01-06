package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/multiverse-os/loading"
	bar "github.com/multiverse-os/loading/bars/thinblocks"
)

func main() {
	fmt.Println("Loading Bar Example")
	fmt.Println("===================")
	fmt.Println("Simple example with one bar")

	loadingBar := loading.Bar().Width(80).Style(bar.Style).Start()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(rand.Intn(135)+22) * time.Millisecond)
		if loadingBar.Increment(1) {
			break
		}
	}
	loadingBar.Status("Completed!").Complete()
}
