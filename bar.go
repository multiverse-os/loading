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

// TODO: This needs to take into consideration if there is a status or percent
// and then base the length on that.
func (bar *Bar) TerminalWidth() *Bar {
	if TerminalWidth() < 20 {
		return bar.Length(12)
	} else {
		return bar.Length(TerminalWidth() - 20)
	}
}

// TODO: Add ability to run the loader for x amount of time to fill up, so we
// have a simple interface with it and we don't have to deal with the loop
// directly (but we should still be able to when we want to)

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

func (bar *Bar) HidePercent() *Bar {
	bar.animation.HidePercent = true
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
	// TODO: This is redundant if we are tracking progress
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

// TODO: Remaining ticks becomes important if we use this for our while loop
func (bar Bar) RemainingTicks() int {
	if float64(bar.length) <= bar.progress {
		bar.progress = float64(bar.length)
		return 0
	} else {
		return bar.length - int(bar.progress)
	}
}

// TODO: Incredibly overly complicated just to have a filled version of the bar;
// it should be the filled in version x the length
func (bar Bar) filled() string {
	// TODO: Seems overly complex
	// TODO: This is where we are failing to do the animation correctly, where
	// with dots we have more than just 1 dot and full dots.
	fill := strings.Repeat(bar.animation.Fill[len(bar.animation.Fill)-1], int(bar.progress))
	// TODO: This condition is everything is filled completely?
	//if bar.progress == float64(bar.length) {
	//	return fill
	//	// TODO: This is likely where our animation is failing to work now
	//} else if 1 < len(bar.animation.Fill) {

	// TODO: There is an issue with this code; this is where the animation cycling
	// should be occuring but instead its only showing the first position of the
	// animation or the last position of the animation; and without it we only
	// show unfilled and last position of the animation

	//if 1 < len(bar.animation.Fill) {
	//	if (len(bar.animation.Fill) - 1) == bar.animationTick {
	//		bar.animationTick = 0
	//	}
	//	fill += bar.animation.Fill[bar.animationTick]
	//	bar.animationTick++
	//}
	return fill
}

func (bar Bar) unfilled() string {
	return strings.Repeat(bar.animation.Unfilled, bar.RemainingTicks())
}

func (bar Bar) percent() float64 {
	return (bar.progress / float64(bar.length)) * 100
}

func (bar *Bar) Increment(progress float64) bool {
	if bar.RemainingTicks() == 0 {
		return false
	}
	bar.progress += (float64(bar.length) / 100) * progress
	bar.increment <- true
	return true
}

// TODO: Frame should probably take into account if the bar is overflowing into
// a second line to prevent issues we were having before
func (bar *Bar) Frame() string {
	fmt.Print(HideCursor())
	fmt.Print(EraseLine(2))
	fmt.Print(CursorStart(1))

	// TODO: There is a simpler way of doing this where we have only 1
	// fmt.Sprintf( and the percent is just empty)
	if bar.animation.HidePercent {
		return fmt.Sprintf(
			"%v%v%v",
			bar.filled(),
			bar.unfilled(),
			bar.animation.Status,
		)
	}

	return fmt.Sprintf(
		bar.animation.Format,
		bar.filled(),
		bar.unfilled(),
		bar.percent(),
		bar.animation.Status,
	)
}

// TODO: Consider what conditions the ticker is necessary, when we wouldn't just
// increment (possibly a loadingBar.Time(15s))
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
	// TODO: This line is being repeated
	bar.progress = float64(bar.length)
	fmt.Print(EraseLine(2))
	fmt.Print(ShowCursor())
	fmt.Print(CursorStart(1))
	bar.end <- true
}
