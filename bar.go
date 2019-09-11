package loading

import (
	"fmt"
	"strings"
	"time"

	color "github.com/multiverse-os/cli/text/ansi/color"
)

type LoadingBar struct {
	width         float64
	progress      float64
	animationTick int
	ticker        *time.Ticker
	end           chan bool
	increment     chan bool
	style         BarStyle
}

type BarStyle struct {
	AnimationSpeed int
	Status         string
	HidePercent    bool
	Fill           []string
	Unfilled       string
	Format         string
}

func DefaultStyle() BarStyle {
	return BarStyle{
		AnimationSpeed: Slowest,
		Status:         "Loading...",
		HidePercent:    false,
		Fill:           []string{color.White("#")},
		Unfilled:       "_",
		Format:         "[%s%s] %0.2f%% %s",
	}
}

func Bar() *LoadingBar {
	terminalWidth, err := TerminalWidth()
	if err != nil {
		fmt.Errorf("[text/loading] failed to calculate terminal width:\n", err)
		return &LoadingBar{}
	}
	bar := &LoadingBar{
		animationTick: 0,
		ticker:        time.NewTicker(time.Millisecond * time.Duration(Fastest)),
		end:           make(chan bool),
		increment:     make(chan bool),
		style:         DefaultStyle(),
	}
	bar.width = float64(bar.barWidth(terminalWidth))
	return bar
}

func (self *LoadingBar) HidePercent(hide bool) *LoadingBar {
	self.style.HidePercent = hide
	return self
}

func (self *LoadingBar) Style(style BarStyle) *LoadingBar {
	if len(style.Status) > 0 {
		self.style.Status = style.Status
	}
	if len(style.Format) > 0 {
		self.style.Format = style.Format
	}
	if len(style.Fill) > 0 {
		self.style.Fill = style.Fill
	}
	if len(style.Unfilled) > 0 {
		self.style.Unfilled = style.Unfilled
	}
	self.style = style
	return self
}

func (self *LoadingBar) Width(width int) *LoadingBar {
	self.width = float64(width)
	return self
}

func (self *LoadingBar) Fill(characters ...string) *LoadingBar {
	self.style.Fill = characters
	return self
}

func (self *LoadingBar) Unfilled(character string) *LoadingBar {
	self.style.Unfilled = character
	return self
}

func (self *LoadingBar) Status(message string) *LoadingBar {
	self.style.Status = message
	return self
}

func (self *LoadingBar) Start() *LoadingBar {
	fmt.Print(self.Frame())
	go self.Animate()
	return self
}

func (self *LoadingBar) incrementSize() float64 { return float64(self.width) / 100.00 }
func (self *LoadingBar) remainingTicks() int    { return int(self.width) - int(self.progress) }
func (self *LoadingBar) percent() float64       { return (self.progress / self.width) * 100 }
func (self *LoadingBar) filled() string {
	if len(self.style.Fill) == 0 {
		self.style = DefaultStyle()
	}
	fill := strings.Repeat(self.style.Fill[len(self.style.Fill)-1], int(self.progress))
	if self.progress == self.width {
		return fill
	} else if len(self.style.Fill) > 1 {
		if (len(self.style.Fill) - 1) == self.animationTick {
			self.animationTick = 0
		}
		fill += self.style.Fill[self.animationTick]
		self.animationTick++
	}
	return fill
}

func (self *LoadingBar) unfilled() string {
	return strings.Repeat(self.style.Unfilled, self.remainingTicks())

}
func (self *LoadingBar) barWidth(terminalWidth int) int {
	return terminalWidth - (len(self.style.Format) + len(self.style.Status) + 10)
}

func (self *LoadingBar) Increment(progress int) (completed bool) {
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

func (self *LoadingBar) Frame() string {
	fmt.Print(HideCursor())
	fmt.Print(EraseLine(2))
	fmt.Print(CursorStart(1))
	if self.style.HidePercent {
		self.style.Format = strings.Replace(self.style.Format, "%0.2f%%", "", -1)
		return fmt.Sprintf(self.style.Format, self.filled(), self.unfilled(), self.Status)
	} else {
		return fmt.Sprintf(self.style.Format, self.filled(), self.unfilled(), self.percent(), self.style.Status)
	}
}

func (self *LoadingBar) Animate() {
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

func (self *LoadingBar) Complete() {
	self.progress = self.width
	fmt.Print(EraseLine(2))
	fmt.Print(ShowCursor())
	fmt.Print(CursorStart(1))
	self.end <- true
}
