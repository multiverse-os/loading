package loading

import (
	"fmt"
	"time"
)

// TODO: Consier complete message as assignable field in Spinner so it can be
// automatically used when End() is called which is important for further
// implementing the interface
type Spinner struct {
	animation    []string
	palette      []string
	message      string
	index        int
	paletteIndex int
	ticker       *time.Ticker
	speed        int
	end          chan bool
}

func NewSpinner(animation []string) *Spinner {
   return &Spinner{
		animation: animation,
		palette:   []string{""},
		ticker:    time.NewTicker(time.Millisecond * time.Duration(Normal)),
		end:       make(chan bool),
	}
}

func (self *Spinner) Animation(animation []string) *Spinner {
	return NewSpinner(animation)
}

func (self *Spinner) Start() {
	go self.Animate()
}

func (self *Spinner) Cancel() {
	self.Complete("")
}

func (self *Spinner) End() {
  self.Complete("")
}

func (self *Spinner) Complete(message string) {
	fmt.Print(EraseLine(2))
	fmt.Print(ShowCursor())
	fmt.Print(CursorStart(1))
	fmt.Printf("%v\n", message)
	self.end <- true
}

func (self *Spinner) Animate() {
	for {
		select {
		case <-self.end:
			return
		case <-self.ticker.C:
			fmt.Print(self.Frame() + self.message)
		}
	}
}

func (self *Spinner) Message(message string) *Spinner {
	self.message = "  " + message
	return self
}

func (self *Spinner) Speed(speed int) *Spinner {
	self.ticker = time.NewTicker(time.Millisecond * time.Duration(speed))
	return self
}

func (self *Spinner) Palette(palette []string) *Spinner {
	self.palette = palette
	return self
}

func (self *Spinner) Frame() string {
	fmt.Print(HideCursor())
	fmt.Print(EraseLine(2))
	fmt.Print(CursorStart(1))
	self.index = increment(self.index, len(self.animation))
	self.paletteIndex = increment(self.paletteIndex, len(self.palette))
	return self.palette[self.paletteIndex] + self.animation[self.index] + "\x1b[0m"
}

func (self *Spinner) Increment(frameCount int) bool {
	self.index = increment(self.index, len(self.animation))
  return self.index != len(self.animation) 
}

func increment(index, max int) int {
	index++
	switch index {
	case max:
		return 0
	default:
		return index
	}
}
