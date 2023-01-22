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

func (self *Bar) DefaultBarAnimation() *Bar {
	return self.Animation(BarAnimation{
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
func (self *Bar) TerminalWidth() *Bar {
	terminalWidth := TerminalWidth()
	return self.Width(terminalWidth)
}

// TODO: Refer to NewSpinner() whjere animation is taken in at initiation as a
// []string. So custom animations can be created without creating a subpackage
// within our library, on the fly custom animations.
// We want to change NewBar() so it takes []string or BarAnimation then if that is
// invalid we go to default
// NewBar(), NewBar(animation)

func NewBar(animation BarAnimation) *Bar {
	bar := &Bar{
		animationTick: 0,
		ticker:        time.NewTicker(time.Millisecond * time.Duration(Fastest)),
		end:           make(chan bool),
		increment:     make(chan bool),
		animation:     animation,
	}
	bar.TerminalWidth()
	return bar
}

func (self *Bar) HidePercent(hide bool) *Bar {
	self.animation.HidePercent = hide
	return self
}

func (self *Bar) Animation(animation BarAnimation) *Bar {
	if len(animation.Status) > 0 {
		self.animation.Status = animation.Status
	}
	if len(animation.Format) > 0 {
		self.animation.Format = animation.Format
	}
	if len(animation.Fill) > 0 {
		self.animation.Fill = animation.Fill
	}
	if len(animation.Unfilled) > 0 {
		self.animation.Unfilled = animation.Unfilled
	}
	self.animation = animation
	return self
}

func (self *Bar) Width(width int) *Bar {
	self.width = float64(width)
	return self
}

func (self *Bar) Fill(characters ...string) *Bar {
	self.animation.Fill = characters
	return self
}

func (self *Bar) Unfilled(character string) *Bar {
	self.animation.Unfilled = character
	return self
}

func (self *Bar) Status(message string) *Bar {
	self.animation.Status = message
	return self
}

func (self *Bar) Start() {
	fmt.Print(self.Frame())
	go self.Animate()
}

func (self Bar) incrementSize() float64 { return float64(self.width) / 100.00 }
func (self Bar) remainingTicks() int    { return int(self.width) - int(self.progress) }
func (self Bar) percent() float64       { return (self.progress / self.width) * 100 }

func (self Bar) filled() string {
	if len(self.animation.Fill) == 0 {
		self.DefaultBarAnimation()
	}
	fill := strings.Repeat(self.animation.Fill[len(self.animation.Fill)-1], int(self.progress))
	if self.progress == self.width {
		return fill
	} else if len(self.animation.Fill) > 1 {
		if (len(self.animation.Fill) - 1) == self.animationTick {
			self.animationTick = 0
		}
		fill += self.animation.Fill[self.animationTick]
		self.animationTick++
	}
	return fill
}

func (self Bar) unfilled() string {
	return strings.Repeat(self.animation.Unfilled, self.remainingTicks())

}
func (self Bar) BarWidth(terminalWidth int) int {
	return terminalWidth - (len(self.animation.Format) + len(self.animation.Status) + 10)
}

func (self *Bar) Increment(progress int) (completed bool) {
	if self.progress >= self.width {
		self.progress = self.width
		completed = true
	} else {
		self.progress += (float64(progress) * self.incrementSize())
		completed = false
	}
	self.increment <- true
	return completed
}

func (self *Bar) Frame() string {
	fmt.Print(HideCursor())
	fmt.Print(EraseLine(2))
	fmt.Print(CursorStart(1))
	if self.animation.HidePercent {
		self.animation.Format = strings.Replace(self.animation.Format, "%0.2f%%", "", -1)
		return fmt.Sprintf(self.animation.Format, self.filled(), self.unfilled(), self.Status)
	} else {
		return fmt.Sprintf(self.animation.Format, self.filled(), self.unfilled(), self.percent(), self.animation.Status)
	}
}

func (self *Bar) Animate() {
	for {
		select {
		case <-self.end:
			fmt.Printf("\n")
			return
		case <-self.increment:
			fmt.Print(self.Frame())
		case <-self.ticker.C:
			fmt.Print(self.Frame())
		}
	}
}

func (self *Bar) End() {
	self.progress = self.width
	fmt.Print(EraseLine(2))
	fmt.Print(ShowCursor())
	fmt.Print(CursorStart(1))
	self.end <- true
}
