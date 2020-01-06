package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/multiverse-os/color"
	"github.com/multiverse-os/loading/spinners/dots"
	"github.com/multiverse-os/loading/spinners/grow"
)

var Rainbow = []string{color.LIGHT_RED, color.RED, color.LIGHT_RED,
	color.LIGHT_YELLOW, color.YELLOW, color.LIGHT_YELLOW, color.LIGHT_GREEN,
	color.GREEN, color.LIGHT_GREEN, color.LIGHT_CYAN, color.CYAN, color.LIGHT_CYAN,
	color.LIGHT_BLUE, color.BLUE, color.LIGHT_BLUE, color.LIGHT_MAGENTA,
	color.MAGENTA, color.LIGHT_MAGENTA}

var Blues = []string{color.LIGHT_BLUE, color.LIGHT_BLUE, color.BLUE, color.BLUE,
	color.BLUE, color.BLUE, color.BLUE, color.BLUE, color.BLUE, color.LIGHT_BLUE,
	color.LIGHT_BLUE, color.LIGHT_BLUE, color.WHITE, color.WHITE, color.LIGHT_CYAN,
	color.LIGHT_CYAN, color.LIGHT_CYAN, color.CYAN, color.LIGHT_CYAN,
	color.LIGHT_CYAN, color.LIGHT_CYAN}

func randomWait() {
	time.Sleep(time.Duration(rand.Intn(2)+2) * time.Second)
}

func main() {
	fmt.Println("Loader Example")
	fmt.Println("==============")
	RainbowSpinner()
	//MultiMessageSpinner()
}

// Options are defined using function chaining, and the available options
// are:
//   1) Speed(int), with aliases spinner.(Slowest,Slow,Normal,Fast,Fastest)
//   2) Messasge(string)
//   3) Palette([]string), expecting ansi color to cycle through

func RainbowSpinner() {
	rainbowDots := spinner.New(dots.Animation).Message("Loading...").Speed(spinner.Normal).Palette(Rainbow).Start()
	randomWait()
	rainbowDots.Complete(color.Green("Loading Complete!"))
}

func MultiMessageSpinner() {
	s := spinner.New(grow.Animation).Message("Loading...").Speed(spinner.Normal)
	s.Start()
	randomWait()
	randomWait()

	// Currently there are only six (6) available spinner options
	// any pull requests for additional spinners will be reviewed
	// and accepted if the code is consistent with the existing
	// library codebase. Pull requests welcome.
	// Loader Animation Options:
	//    (1) circle.Animation   ["◐","◓", "◑", "◒"]
	//    (2) clock.Animation    ["🕐","🕑","🕒","🕓",...]
	//    (3) dots.Animation     ["⠋","⠙","⠹","⠸",...]
	//    (4) lines.Animation    ["-","\","|","/"]
	//    (5) moon.Animation     ["🌑","🌒","🌓","🌔",...]
	//    (6) triangle.Animation ["◢","◣","◤","◥"]

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
