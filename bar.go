package loading

import (
	"fmt"
	"strings"
	"time"

	width "golang.org/x/text/width"

	dots "github.com/multiverse-os/loading/spinners/dots"
)

// TODO: Add a cutoff to prevent more than the length of the screen being
// printed or at least track the number of lines so we can clear it correctly
// and never have a condition the screen renders a new frame on a new line.

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
	frames        []string
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
// TODO
// MUST Inlcude if percentage is shown; currently assumes it IS shown, so we
// need to make modifications for cleaner output; in addition, we may want the
// option to tell the loading bar how long our status messages will be, or
// instead maybe load status messages into a slice or map, so we can use that
// to determine maximum loading bar width
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

// TODO: Maybe we should move to very basic on/off style and use spinner on the
// edge
//type Animation struct {
//	Filled   rune
//	Unfilled rune
//	Spinner  *Spinner
//}

func NewBar(loadingFrames []string) *Bar {
	if len(loadingFrames) == 0 &&
		(len(loadingFrames[Filled]) == 0 || len(loadingFrames[Unfilled]) == 0) {
		loadingFrames = []string{"□", "■"}
	}

	runeProperties, _ := width.Lookup([]byte(loadingFrames[Unfilled]))
	//fmt.Printf("runeProperties.Kind()(%v)\n", runeProperties.Kind())
	var runeWidth int
	switch runeProperties.Kind() {
	case width.EastAsianWide, width.EastAsianFullwidth:
		runeWidth = 2
	case width.EastAsianAmbiguous, width.EastAsianNarrow, width.EastAsianHalfwidth:
		runeWidth = 1
	default: // Neutral
		runeWidth = 1
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
		frames:        loadingFrames,
		spinner:       NewSpinner(dots.Animation).LoadingBar(true),
		runeWidth:     uint(runeWidth),
		format:        defaultFormat(),
		percent:       true,
		ticker:        time.NewTicker(time.Millisecond * time.Duration(Average)),
	}
	bar.TerminalWidth()
	return bar
}

func (bar *Bar) NewTicker(speed int) *Bar {
	bar.ticker = time.NewTicker(time.Millisecond * time.Duration(speed))
	return bar
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

	// TODO
	// Turn the spinner on
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
	return strings.Repeat(
		bar.frames[Filled],
		int(uint(bar.progress)/bar.runeWidth),
	)
	// TODO This is where we want to put our spinner or maybe separated further
	//bar.frames["fill"][bar.animationTick]
}

// TODO
//
//   - Why does unfilled count end on 45
//
//   - Why am I getting a flashing symbol at the end as if the unfilled count
//     is increasing and decreasing
func (bar Bar) unfilled() string {
	return strings.Repeat(
		bar.frames[Unfilled],
		int(bar.RemainingTicks()/bar.runeWidth),
	)
}

func (bar *Bar) Increment(percent float64) bool {
	if bar.RemainingTicks() == 0 {
		return false
	}
	incrementAmount := roundFloat((float64(bar.length) / 100 * percent), 2)
	bar.progress += incrementAmount
	//bar.spinner.Increment(0)
	bar.increment <- true
	return true
}

// TODO: Frame should probably take into account if the bar is overflowing into
// a second line to prevent issues we were having before
func (bar *Bar) Frame() string {
	fmt.Sprintf("test %v", Text("test").sgr(style(Blue, Foreground)).String())
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
			// TODO: THIS IS WHERE WE CAN ENSURE IT NEVER GOES OVER WIDTH LIMIT
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
