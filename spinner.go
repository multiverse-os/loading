package loading

import (
	"fmt"
	"time"
)

type LoadingSpinner struct {
	animation    []string
	palette      []string
	message      string
	index        int
	paletteIndex int
	ticker       *time.Ticker
	speed        int
	end          chan bool
}

func Spinner(animation []string) *LoadingSpinner {
	return &LoadingSpinner{
		animation: animation,
		palette:   []string{""},
		ticker:    time.NewTicker(time.Millisecond * time.Duration(Normal)),
		end:       make(chan bool),
	}
}

func Animation(animation []string) *LoadingSpinner {
	return Spinner(animation)
}

func (self *LoadingSpinner) Start() *LoadingSpinner {
	go self.Animate()
	return self
}

func (self *LoadingSpinner) Cancel() {
	self.Complete("")
}

func (self *LoadingSpinner) Complete(message string) {
	self.end <- true
	fmt.Print(EraseLine(2))
	fmt.Print(ShowCursor())
	fmt.Print(CursorStart(1))
	fmt.Println(message)
}

func (self *LoadingSpinner) Animate() {
	for {
		select {
		case <-self.end:
			return
		case <-self.ticker.C:
			fmt.Print(self.Frame() + self.message)
		}
	}
}

func (self *LoadingSpinner) Message(message string) *LoadingSpinner {
	self.message = message
	return self
}

func (self *LoadingSpinner) Speed(speed int) *LoadingSpinner {
	self.ticker = time.NewTicker(time.Millisecond * time.Duration(speed))
	return self
}

func (self *LoadingSpinner) Palette(palette []string) *LoadingSpinner {
	self.palette = palette
	return self
}

func (self *LoadingSpinner) Frame() string {
	fmt.Print(HideCursor())
	fmt.Print(EraseLine(2))
	fmt.Print(CursorStart(1))
	self.index = Increment(self.index, len(self.animation))
	self.paletteIndex = Increment(self.paletteIndex, len(self.palette))
	return self.palette[self.paletteIndex] + self.animation[self.index] + "\x1b[0m"
}

func Increment(index, max int) int {
	index++
	switch index {
	case max:
		return 0
	default:
		return index
	}
}
