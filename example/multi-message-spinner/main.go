package main

import (
	"fmt"
	"math/rand"
	"time"

	loading "github.com/multiverse-os/loading"
	lines "github.com/multiverse-os/loading/spinners/lines"
	moon "github.com/multiverse-os/loading/spinners/moon"
)

func main() {
	fmt.Println("Loader Example")
	fmt.Println("==============")
	MultiMessageSpinner()
	fmt.Println("==============")
	TimedSpinner()
	fmt.Println("==============")
}

func MultiMessageSpinner() {
	s := loading.Spinner(moon.Animation).Message("Loading...").Speed(loading.Normal)
	s.Start()
	randomWait()
	randomWait()

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

func TimedSpinner() {
	s := loading.Spinner(lines.Animation).Message("Loading...").Speed(loading.Normal)
	s.Start()
	start := time.Now()
	timer := time.NewTimer(15 * time.Second)
	tick := time.Tick(15 * time.Millisecond)
	for {
		select {
		case <-tick:
			s.Message(fmt.Sprintf("Loading for: %v", time.Since(start)))
		}
		if time.Since(start) > 15*time.Second {
			break
		}
	}
	<-timer.C
	timer.Stop()
	s.Complete(fmt.Sprintf("Completed after %v", time.Since(start)))
}

func randomWait() {
	time.Sleep(time.Duration(rand.Intn(2)+2) * time.Second)
}
