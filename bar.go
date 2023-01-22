package loading

import (
	"fmt"
	"strings"
	"time"

	color "github.com/multiverse-os/ansi"
)

type Bar struct {
	width         float64
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

func (bar *Bar) DefaultBarAnimation() *Bar {
	return bar.Animation(BarAnimation{
		AnimationSpeed: Slowest,
		Status:         "Loading...",
		HidePercent:    false,
		Fill:           []string{color.White("#")},
		Unfilled:       "_",
		Format:         "[%s%s] %0.2f%% %s",
	}).TerminalWidth()
}

// TODO: Something specific to the system should not really be a method of Bar,
// that doesnt really flow with the logical structure we are trying to create.
// Say if you ahve a folder system, and in that system you have a folder anem
// "BackgruondMusic" and another folder named "HoldMusic" our design comes to to
// the essential concept that we would create a folder named Music and within it
// have music files ,and Background folder and Hold folder.
func (bar *Bar) TerminalWidth() *Bar {
	terminalWidth := TerminalWidth()
	return bar.Width(terminalWidth)
}

// TODO: Refer to NewSpinner() whjere animation is taken in at initiation as a
// []string. So custom animations can be created without creating a subpackage
// within our library, on the fly custom animations.
// We want to change NewBar() so it takes []string or BarAnimation then if that is
// invalid we go to default
// NewBar(), NewBar(animation)

func NewBar(animation BarAnimation) *Bar {
	bar := &Bar{
		width:         80,
		animationTick: 0,
		ticker:        time.NewTicker(time.Millisecond * time.Duration(Fastest)),
		end:           make(chan bool),
		increment:     make(chan bool),
		animation:     animation,
	}
	bar.TerminalWidth()
	return bar
}

func (bar *Bar) HidePercent(hide bool) *Bar {
	bar.animation.HidePercent = hide
	return bar
}

func (bar *Bar) Animation(animation BarAnimation) *Bar {
	if len(animation.Status) > 0 {
		bar.animation.Status = animation.Status
	}
	if len(animation.Format) > 0 {
		bar.animation.Format = animation.Format
	}
	if len(animation.Fill) > 0 {
		bar.animation.Fill = animation.Fill
	}
	if len(animation.Unfilled) > 0 {
		bar.animation.Unfilled = animation.Unfilled
	}
	bar.animation = animation
	return bar
}

func (bar *Bar) Width(width int) *Bar {
	bar.width = float64(width)
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
	fmt.Print(bar.Frame())
	go bar.Animate()
}

func (bar Bar) incrementSize() float64 { return float64(bar.width) / 100.00 }
func (bar Bar) remainingTicks() int    { return int(bar.width) - int(bar.progress) }
func (bar Bar) percent() float64       { return (bar.progress / bar.width) * 100 }

func (bar Bar) filled() string {
	if len(bar.animation.Fill) == 0 {
		bar.DefaultBarAnimation()
	}
	fill := strings.Repeat(bar.animation.Fill[len(bar.animation.Fill)-1], int(bar.progress))
	if bar.progress == bar.width {
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
	return strings.Repeat(bar.animation.Unfilled, bar.remainingTicks())

}
func (bar Bar) BarWidth(terminalWidth int) int {
	return terminalWidth - (len(bar.animation.Format) + len(bar.animation.Status) + 10)
}

func (bar *Bar) Increment(progress int) (completed bool) {
	fmt.Printf("frame+1")
	if bar.progress >= bar.width {
		bar.progress = bar.width
		completed = true
	} else {
		bar.progress += (float64(progress) * bar.incrementSize())
		completed = false
	}
	bar.increment <- true
	return completed
}

func (bar *Bar) Frame() string {
	fmt.Printf("frame+1")
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
	bar.progress = bar.width
	fmt.Print(EraseLine(2))
	fmt.Print(ShowCursor())
	fmt.Print(CursorStart(1))
	bar.end <- true
}
