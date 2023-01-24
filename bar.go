package loading

import (
	"fmt"
	"strings"
	"time"
)

type Bar struct {
	length        int
	progress      float64
	animationTick int
	ticker        *time.Ticker
	end           chan bool
	increment     chan bool
	animation     BarAnimation
}

type BarAnimation struct {
	AnimationSpeed int
	Status         string
	HidePercent    bool
	Fill           []string
	Unfilled       string
	Format         string
}

func (bar *Bar) TerminalWidth() *Bar {
	return bar.Length(TerminalWidth())
}

func NewBar(animation BarAnimation) *Bar {
	bar := &Bar{
		length:        80,
		animationTick: 0,
		ticker:        time.NewTicker(time.Millisecond * time.Duration(Fastest)),
		end:           make(chan bool),
		increment:     make(chan bool),
		animation:     animation,
	}
	// TODO: This breaks it and has it repeat instead of clear
	//bar.TerminalWidth()
	terminalWidth := TerminalWidth()
	fmt.Printf("terminal width: %v", terminalWidth)
	return bar
}

func (bar *Bar) HidePercent(hide bool) *Bar {
	bar.animation.HidePercent = hide
	return bar
}

func (bar *Bar) Animation(animation BarAnimation) *Bar {
	if 0 < len(animation.Status) {
		bar.animation.Status = animation.Status
	}
	if 0 < len(animation.Format) {
		bar.animation.Format = animation.Format
	}
	if 0 < len(animation.Fill) {
		bar.animation.Fill = animation.Fill
	}
	// TODO: This is redudant if we are tracking progress
	if 0 < len(animation.Unfilled) {
		bar.animation.Unfilled = animation.Unfilled
	}
	bar.animation = animation
	return bar
}

func (bar *Bar) Length(length int) *Bar {
	bar.length = length
	return bar
}

func (bar *Bar) Fill(characters ...string) *Bar {
	bar.animation.Fill = characters
	return bar
}

func (bar *Bar) Unfilled(character string) *Bar {
	bar.animation.Unfilled = character
	return bar
}

func (bar *Bar) Status(message string) *Bar {
	bar.animation.Status = message
	return bar
}

func (bar *Bar) Start() {
	go bar.Animate()
}

//func (bar Bar) incrementSize() float64 { return float64(bar.length) / 100.00 }

func (bar Bar) RemainingTicks() int {
	if float64(bar.length) <= bar.progress {
		bar.progress = float64(bar.length)
		return 0
	} else {
		return bar.length - int(bar.progress)
	}
}

func (bar Bar) percent() float64 {
	return (bar.progress / float64(bar.length)) * 100
}

// TODO: Incredibly overly complicated just to have a filled version of the bar;
// it should be the filled in version x the length
func (bar Bar) filled() string {
	// TODO: Seems overly complex
	fill := strings.Repeat(bar.animation.Fill[len(bar.animation.Fill)-1], int(bar.progress))
	if bar.progress == float64(bar.length) {
		return fill
	} else if len(bar.animation.Fill) > 1 {
		if (len(bar.animation.Fill) - 1) == bar.animationTick {
			bar.animationTick = 0
		}
		fill += bar.animation.Fill[bar.animationTick]
		bar.animationTick++
	}
	return fill
}

func (bar Bar) unfilled() string {
	return strings.Repeat(bar.animation.Unfilled, bar.RemainingTicks())
}

// TODO: Must be able to increment float64 not int, so we can do .5 for example,
// or base it off a variable rate like with download
func (bar *Bar) Increment(progress float64) bool {
	bar.progress += progress
	if bar.RemainingTicks() == 0 {
		return false
	}
	//bar.progress += (float64(progress) * bar.incrementSize())
	bar.increment <- true
	return true
}

func (bar *Bar) Frame() string {
	fmt.Print(HideCursor())
	fmt.Print(EraseLine(2))
	fmt.Print(CursorStart(1))
	if bar.animation.HidePercent {
		bar.animation.Format = strings.Replace(bar.animation.Format, "%0.2f%%", "", -1)
		return fmt.Sprintf(bar.animation.Format, bar.filled(), bar.unfilled(), bar.Status)
	} else {
		return fmt.Sprintf(bar.animation.Format, bar.filled(), bar.unfilled(), bar.percent(), bar.animation.Status)
	}
}

func (bar *Bar) Animate() {
	for {
		select {
		case <-bar.end:
			fmt.Printf("\n")
			return
		case <-bar.increment:
			fmt.Print(bar.Frame())
		case <-bar.ticker.C:
			fmt.Print(bar.Frame())
		}
	}
}

func (bar *Bar) End() {
	bar.progress = float64(bar.length)
	fmt.Print(EraseLine(2))
	fmt.Print(ShowCursor())
	fmt.Print(CursorStart(1))
	bar.end <- true
}
