package loading

import (
	"fmt"
	"strings"
	"time"
)

const (
	Unfilled = 0
	Filled   = 1
)

type Bar struct {
	format        string
	length        uint
	progress      float64
	spinner       *Spinner
	animationTick int
	ticker        *time.Ticker
	end           chan bool
	increment     chan bool
	runeWidth     uint
	animation     []string
	speed         uint
	status        string
	percent       bool
}

// TODO
// Would like to support b/sec Kb/sec Mb/sec
// Would like to support Time left MM:SS + HH:MM:SS
// TODO: This needs to take into consideration if there is a status or percent
// and then base the length on that.
// Why not go all the way down to 8+ (9 = 8 spaces + 1 loading bar)
func NewBar(animationFrames []string) *Bar {
	if len(animationFrames) == 0 &&
		(len(animationFrames[Filled]) == 0 ||
			len(animationFrames[Unfilled]) == 0) {
		animationFrames = DefaultBar()
	}

	// NOTE
	// Ticker is for when we have a specific wait time known prior to creation of
	// the loader. Otherwise we can simply increment manually
	// TODO: How about just assign color to each filled and unfilled
	bar := &Bar{
		status:        "",
		animationTick: 0,
		end:           make(chan bool),
		increment:     make(chan bool),
		animation:     animationFrames,
		spinner:       NewSpinner(DefaultSpinner()).LoadingBar(true),
		runeWidth:     RuneWidth(animationFrames[Unfilled]),
		format:        defaultFormat(),
		percent:       true,
		ticker:        time.NewTicker(time.Millisecond * time.Duration(Average)),
	}
	bar.TerminalWidth()
	return bar
}

func (bar *Bar) Animation(frames []string) {
	bar.animation = frames
}

func (bar *Bar) NewTicker(speed int) *Bar {
	bar.ticker = time.NewTicker(time.Millisecond * time.Duration(speed))
	return bar
}

func (bar *Bar) TerminalWidth() *Bar {
	if TerminalWidth() < 20 {
		return bar.Length(12)
	} else {
		return bar.Length(uint(TerminalWidth() - 20))
	}
}

func (bar *Bar) ShowPercent(visible bool) *Bar {
	bar.percent = visible
	return bar
}

func (bar *Bar) Length(length uint) *Bar {
	bar.length = length
	return bar
}

func (bar *Bar) Status(message string) *Bar {
	bar.status = message
	return bar
}

func (bar *Bar) Start() {
	go bar.Animate()
}

// TODO
// Remaining ticks becomes important if we use this for our while loop
func (bar Bar) RemainingTicks() uint {
	if float64(bar.length) <= bar.progress {
		bar.progress = float64(bar.length)
		return 0
	} else {
		return bar.length - uint(bar.progress)
	}
}

// TODO
// This is where we are failing to do the animation correctly, where
// with dots we have more than just 1 dot and full dots.
func (bar Bar) filled() string {
	return strings.Repeat(
		bar.animation[Filled],
		int(uint(bar.progress)/bar.runeWidth),
	)
}

// TODO
// Why does unfilled count end on 45
// Why am I getting a flashing symbol at the end as if the unfilled count
// is increasing and decreasing
func (bar Bar) unfilled() string {
	return strings.Repeat(
		bar.animation[Unfilled],
		int(bar.RemainingTicks()/bar.runeWidth),
	)
}

func (bar *Bar) Increment(percent float64) bool {
	if bar.RemainingTicks() == 0 {
		return false
	}
	incrementAmount := roundFloat((float64(bar.length) / 100 * percent), 2)
	bar.progress += incrementAmount
	bar.increment <- true
	return true
}

// TODO
// Frame should probably take into account if the bar is overflowing into
// a second line to prevent issues we were having before
func (bar *Bar) Frame() string {
	fmt.Print(HideCursor())
	fmt.Print(EraseLine(2))
	fmt.Print(CursorStart(1))

	return fmt.Sprintf(
		bar.format,
		bar.filled()+bar.spinner.Frame(),
		Text(bar.unfilled()).sgr(style(Black, Foreground)).String(),
		bar.percentStatus()+bar.status,
	)
}

func defaultFormat() string { return " %s%s%s" }

func (bar Bar) percentage() float64 {
	return ((bar.progress / float64(bar.length)) * 100)
}

func (bar Bar) percentStatus() string {
	return fmt.Sprintf(" %0.2f%% ", bar.percentage())
}

func (bar *Bar) Animate() {
	for {
		select {
		case <-bar.end:
			// TODO
			// Fix the bar ending wrong value by putting it at 100%, with status
			// message here before new line
			fmt.Printf("\n")
			return
		case <-bar.increment:
			// TODO
			// THIS IS WHERE WE CAN ENSURE IT NEVER GOES OVER WIDTH LIMIT
			fmt.Print(bar.Frame())
		case <-bar.ticker.C:
			fmt.Print(bar.Frame())
		}
	}
}

func (bar *Bar) End() {
	// TODO
	// If enter is detected go back up or preferably disable CR
	bar.progress = float64(100)
	fmt.Print(EraseLine(2))
	fmt.Print(ShowCursor())
	fmt.Print(CursorStart(1))
	bar.end <- true
}
