package loading

import (
	"fmt"
	"strings"
	"time"
)

// TODO: Add a cutoff to prevent more than the length of the screen being
// printed or at least track the number of lines so we can clear it correctly
// and never have a condition the screen renders a new frame on a new line.
type Bar struct {
	length        uint
	progress      float64
	animationTick int
	ticker        *time.Ticker
	end           chan bool
	increment     chan bool
	animation     BarAnimation
}

type BarAnimation struct {
	AnimationSpeed uint
	Status         string
	HidePercent    bool
	Fill           []string
	RuneWidth      uint
	Unfilled       string
	Format         string
}

// TODO: This needs to take into consideration if there is a status or percent
// and then base the length on that.
func (bar *Bar) TerminalWidth() *Bar {
	if TerminalWidth() < 20 {
		return bar.Length(12)
	} else {
		return bar.Length(uint(TerminalWidth() - 20))
	}
}

// TODO: Add ability to run the loader for x amount of time to fill up, so we
// have a simple interface with it and we don't have to deal with the loop
// directly (but we should still be able to when we want to)

func NewBar(animation BarAnimation) *Bar {

	//if 0 < len(animation.Status) {
	//	bar.animation.Status = animation.Status
	//}

	//if 0 == len(animation.Format) {

	//   animation.Format = " "
	//}
	//	bar.animation.Format = animation.Format
	//}
	//if 0 < len(animation.Fill) {
	//	bar.animation.Fill = animation.Fill
	//}
	//// TODO: This is redundant if we are tracking progress
	//if 0 < len(animation.Unfilled) {
	//	bar.animation.Unfilled = animation.Unfilled
	//}

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

func (bar *Bar) ShowPercent(show bool) *Bar {
	bar.animation.HidePercent = show
	return bar
}

// TODO: I kinda hate this function and I want to simplify it
//
//	this doesnt have fail conditions either, we could simply check against
//	if the animation exists too
func (bar *Bar) Animation(animation BarAnimation) *Bar {
	bar.animation = animation
	return bar
}

func (bar *Bar) Length(length uint) *Bar {
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
func (bar Bar) RemainingTicks() uint {
	if float64(bar.length) <= bar.progress {
		bar.progress = float64(bar.length)
		return 0
	} else {
		return bar.length - uint(bar.progress)
	}
}

// TODO: This is where we are failing to do the animation correctly, where
// with dots we have more than just 1 dot and full dots.
func (bar Bar) filled() string {
	fill := strings.Repeat(
		bar.animation.Fill[len(bar.animation.Fill)-1],
		int(uint(bar.progress)/bar.animation.RuneWidth),
	)
	return fill
}

func (bar Bar) unfilled() string {

	//rune, size := utf8.DecodeRune([]byte(bar.animation.Unfilled))

	//runeProperties := width.LookupRune(bar.animation.Unfilled)

	return strings.Repeat(
		bar.animation.Unfilled,
		int(bar.RemainingTicks()/bar.animation.RuneWidth),
	)
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

	// TODO: NOVERB is showing up after percent and would be in this frame
	// function

	// NOTE: Spinner animation at end of the loading bar
	fill := bar.filled() + bar.animation.Fill[bar.animationTick]
	if bar.animationTick < len(bar.animation.Fill)-1 {
		bar.animationTick += 1
	} else {
		bar.animationTick = 0
	}

	// TODO: There is a simpler way of doing this where we have only 1
	// fmt.Sprintf( and the percent is just empty)
	//if bar.animation.HidePercent {
	//	return fmt.Sprintf(
	//		"%v%v%v",
	//		fill,
	//		bar.unfilled(),

	//		bar.animation.Status,
	//	)
	//}

	return fmt.Sprintf(
		bar.animation.Format,
		fill,
		bar.unfilled(),
		bar.percent(),
		bar.animation.Status,
	)
}

// TODO: Consider what conditions the ticker is necessary, when we wouldn't just
// increment (possibly a loadingBar.Time(15s))
func (bar *Bar) Animate() {
	for {
		// TODO: A ticker here would likely use less processor but test it to
		// guarantee this is more than a guess
		select {
		case <-bar.end:
			fmt.Printf("\n")
			return
		case <-bar.increment:
			fmt.Print(bar.Frame())
		case <-bar.ticker.C:
			fmt.Printf(bar.Frame())
			//fmt.Print(bar.Frame())
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
