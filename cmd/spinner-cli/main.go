package main

import (
	"fmt"
	"math/rand"
	"time"

	loading "github.com/multiverse-os/loading"
	moon "github.com/multiverse-os/loading/spinners/moon"
)

func randomWait() {
	time.Sleep(time.Duration(rand.Intn(2)+2) * time.Second)
}

func main() {
	fmt.Println("Loader Example")
	fmt.Println("==============")
	MultiMessageSpinner()
}

// Options are defined using function chaining, and the available options
// are:
//   1) Speed(int), with aliases spinner.(Slowest,Slow,Normal,Fast,Fastest)
//   2) Messasge(string)
//   3) Palette([]string), expecting ansi colors to cycle through

func MultiMessageSpinner() {
	s := loading.NewSpinner(moon.Animation).Message("Loading...").Speed(loading.Average)
	s.Start()
	randomWait()
	randomWait()

	// Currently there are only six (6) available spinner options
	// any pull requests for additional spinners will be reviewed
	// and accepted if the code is consistent with the existing
	// library codebase. Pull requests welcome.
	// Loader Animation Options:
	//    (1) circle.Animation   ["◐","◓", "◑", "◒"]
	//    (2) dots.Animation     ["⠋","⠙","⠹","⠸",...]
	//    (3) lines.Animation    ["-","\","|","/"]
	//    (4) moon.Animation     ["🌑","🌒","🌓","🌔",...]
	//    (5) triangle.Animation ["◢","◣","◤","◥"]

	// To provide a more interactive UI, messages, speed and even palette can
	// be updated while the spinner animation is active to provide updates
	// which lets the user know the program is not frozen.
	s.Message("Water, Dirt & Grass")
	randomWait()
	s.Message("Trees, Debris & Hideouts")
	randomWait()
	s.Message("Wildlife, Werewolves & Bandits")
	randomWait()
	s.Message("Sounds of wildlife & trees waving in the wind")
	randomWait()
	s.Message("Hiding treasure in the haunted woods...")
	randomWait()
	s.Complete("Completed")

}
