package loading

import (
	"fmt"
	"time"
)

//type Frame struct {
//	Rune string
//	Text
//}

// TODO: Consier complete message as assignable field in Spinner so it can be
// automatically used when End() is called which is important for further
// implementing the interface, palette will give us both rainbow-type effects
// but also allow us to do like light to dark to fill in bars.
type Spinner struct {
	animation      []string
	palette        []string
	message        string
	animationIndex int
	paletteIndex   int
	ticker         *time.Ticker
	speed          int
	end            chan bool
	bar            bool
}

func NewSpinner(animation []string) *Spinner {
	return &Spinner{
		animation: animation,
		speed:     Average,
		bar:       false,
		palette:   []string{""},
		ticker:    time.NewTicker(time.Millisecond * time.Duration(Average)),
		end:       make(chan bool),
	}
}

func (spinner *Spinner) hasPalette() bool { return len(spinner.palette) == 0 }

func (spinner *Spinner) LoadingBar(bar bool) *Spinner {
	spinner.bar = bar
	return spinner
}

func (spinner *Spinner) Animation(animation []string) *Spinner {
	return NewSpinner(animation)
}

func (spinner *Spinner) Start()  { go spinner.Animate() }
func (spinner *Spinner) Cancel() { spinner.Complete("") }
func (spinner *Spinner) End()    { spinner.Complete("") }

func (spinner *Spinner) Complete(message string) {
	if !spinner.bar {
		fmt.Print(EraseLine(2))
		fmt.Print(ShowCursor())
		fmt.Print(CursorStart(1))
		fmt.Printf("%v\n", message)
	}
	spinner.end <- true
}

func (spinner *Spinner) Animate() {
	for {
		select {
		case <-spinner.end:
			return
		case <-spinner.ticker.C:
			if !spinner.bar {
				fmt.Print(spinner.Frame() + spinner.message)
			} else {
				fmt.Print(spinner.Frame())
			}
		}
	}
}

func (spinner *Spinner) Message(message string) *Spinner {
	spinner.message = fmt.Sprintf("  %v", message)
	return spinner
}

func (spinner *Spinner) Speed(speed int) *Spinner {
	spinner.ticker = time.NewTicker(time.Millisecond * time.Duration(speed))
	return spinner
}

func (spinner *Spinner) Palette(palette []string) *Spinner {
	spinner.palette = palette
	return spinner
}

func (spinner Spinner) totalPalette() int { return len(spinner.animation) }
func (spinner Spinner) totalFrames() int  { return len(spinner.animation) }

func (spinner Spinner) firstPalette() string {
	if spinner.totalPalette() != 0 {
		return spinner.animation[0]
	}
	return ""
}

func (spinner Spinner) lastPalette() string {
	if spinner.totalPalette() != 0 {
		return spinner.animation[spinner.totalPalette()-1]
	}
	return ""
}

func (spinner Spinner) firstFrame() string {
	if spinner.totalFrames() != 0 {
		return spinner.animation[0]
	}
	return ""
}

func (spinner Spinner) lastFrame() string {
	if spinner.totalFrames() != 0 {
		return spinner.animation[spinner.totalFrames()-1]
	}
	return ""
}

// TODO
// Would like to be able to have a palette to easily pull colors from
// but also the palette concept is what colors we cycle through
//
// TODO
// Create a function to grab the frame for a given 0.1-0.9
//
//	This gives us the ability to animate frames to show progress in greater
//	resolution

// Two options here for fractional frame:
//   - I could take the value, and make it the new maximum, cycling up to
//     the fractional percentage value
//   - I could just slowly add parts of the animation to make it appear
//     like its representing 100% with greater resolution
func (spinner Spinner) FractionalFrameIndex(percentage int) int {
	return spinner.totalFrames() * (percentage / 100)
}

func (spinner *Spinner) Frame() string {
	if !spinner.bar {
		fmt.Print(HideCursor())
		fmt.Print(EraseLine(2))
		fmt.Print(CursorStart(1))
	}
	fmt.Printf("spinner.animation(%v)\n", spinner.animation)
	fmt.Printf("len(spinner.animation)(%v)\n", spinner.totalFrames())
	spinner.animationIndex = increment(spinner.animationIndex, spinner.totalFrames())
	if spinner.hasPalette() {
		spinner.paletteIndex = increment(spinner.paletteIndex, spinner.totalPalette())
		// TODO: For use with loading bar we can't do simple reset anymore, need
		// matching close to not ruin ANSI style on loading bar
		return spinner.palette[spinner.paletteIndex] + spinner.animation[spinner.animationIndex] + "\x1b[0m"
	} else {
		return spinner.animation[spinner.animationIndex]
	}
}

func (spinner *Spinner) Increment(skipFrames float64) bool {
	spinner.animationIndex = increment(spinner.animationIndex+int(skipFrames), spinner.totalFrames())
	return spinner.animationIndex != len(spinner.animation)
}

func increment(tick, max int) int {
	fmt.Printf("tick++(%v) max(%v)\n", tick, max)
	tick++
	if tick == max {
		return 0
	} else {
		return tick
	}
}
