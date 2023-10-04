package loading

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/text/width"
)

// TODO: Add a cutoff to prevent more than the length of the screen being
// printed or at least track the number of lines so we can clear it correctly
// and never have a condition the screen renders a new frame on a new line.
type Bar struct {
	format         string
	length         uint
	progress       float64
	animationTick  int
	ticker         *time.Ticker
	end            chan bool
	increment      chan bool
	runeWidth      uint
	frames         map[string][]string
	speed          uint
	status         string
	percentVisible bool
}

// TODO
// Would like to support b/sec Kb/sec Mb/sec
// Would like to support Time left MM:SS + HH:MM:SS
type Format struct {
	Before, After           string
	Prefix, Suffix          string
	RuneWidth               uint8
	StatusLength, BarLength uint8
	Visible                 bool
}

func FormatWithPercent() string    { return " %s%s %0.2f%% %s" }
func FormatWithoutPercent() string { return " %s%s %s" }

func UndefinedBar() map[string][]string {
	return map[string][]string{
		"fill":     []string{},
		"unfilled": []string{},
	}
}

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

func NewBar(animationFrames map[string][]string) *Bar {
	fmt.Printf("animationFrames(%v)\n", animationFrames)

	fmt.Printf("\nanimationFrames(%v)\n", animationFrames)
	fmt.Printf("len(animationFrames['fill'])(%v)\n", len(animationFrames["fill"]))
	fmt.Printf("len(animationFrames['unfilled'])(%v)\n\n", len(animationFrames["unfilled"]))

	fmt.Printf("\nanimationFrames != nil (%v)\n", animationFrames != nil)
	fmt.Printf("len(animationFrames['fill']) == 0 (%v)\n", len(animationFrames["fill"]) != 0)
	fmt.Printf("len(animationFrames['unfilled']) == 0 (%v)\n\n", len(animationFrames["unfilled"]) != 0)

	if animationFrames != nil &&
		(len(animationFrames["fill"]) == 0 ||
			len(animationFrames["unfilled"]) == 0) {
		animationFrames = map[string][]string{
			"fill":     []string{"■"},
			"unfilled": []string{"□"},
		}
	}

	runeProperties, _ := width.Lookup([]byte(animationFrames["unfilled"][0]))
	fmt.Printf("runeProperties.Kind()(%v)\n", runeProperties.Kind())
	fmt.Printf("\n\n")
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
	bar := &Bar{
		status:         "",
		animationTick:  0,
		ticker:         time.NewTicker(time.Millisecond * time.Duration(Fastest)),
		end:            make(chan bool),
		increment:      make(chan bool),
		frames:         animationFrames,
		runeWidth:      uint(runeWidth),
		format:         FormatWithPercent(),
		percentVisible: true,
	}
	bar.TerminalWidth()
	return bar
}

func (bar *Bar) ShowPercent(show bool) *Bar {
	bar.percentVisible = show
	if show {
		bar.format = FormatWithPercent()
	} else {
		bar.format = FormatWithoutPercent()
	}
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

func (bar *Bar) Start() { go bar.Animate() }

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
		bar.frames["fill"][len(bar.frames["fill"])-1],
		int(uint(bar.progress)/bar.runeWidth),
	) + bar.frames["fill"][bar.animationTick]
}

// TODO
//
//   - Why does unfilled count end on 45
//
//   - Why am I getting a flashing symbol at the end as if the unfilled count
//     is increasing and decreasing
func (bar Bar) unfilled() string {
	return strings.Repeat(
		bar.frames["unfilled"][0],
		int(bar.RemainingTicks()/bar.runeWidth),
	)
}

func (bar Bar) percent() float64 {
	return (bar.progress / float64(bar.length)) * 100
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

// TODO: Frame should probably take into account if the bar is overflowing into
// a second line to prevent issues we were having before
func (bar *Bar) Frame() string {
	fmt.Print(HideCursor())
	fmt.Print(EraseLine(2))
	fmt.Print(CursorStart(1))

	// TODO: NOVERB is showing up after percent and would be in this frame
	// function

	// TODO
	// Add optional spinner animation at end of the loading bar; could simplify
	// the loading bar by having just filled and unfilled
	// then have a spinner inbetween
	if bar.animationTick < len(bar.frames["fill"])-1 {
		bar.animationTick += 1
	} else {
		bar.animationTick = 0
	}

	// TODO
	// Here is our problem with the flashing end item; it goes back and forth
	// between difference of 3 total as the:
	//   - filled increases
	//   - unfilled decreases
	//
	//
	//fmt.Printf(
	//	"filled(%v) + unfilled(%v) = (%v) \n",
	//	len(bar.filled()),
	//	len(bar.unfilled()),
	//	len(bar.filled())+len(bar.unfilled()),
	//)

	//fmt.Printf("bar.length(%v)\n", bar.length)

	if bar.percentVisible {
		return fmt.Sprintf(
			bar.format,
			bar.filled(),
			bar.unfilled(),
			bar.percent(),
			bar.status,
		)
	} else {
		return fmt.Sprintf(
			bar.format,
			bar.filled(),
			bar.unfilled(),
			bar.status,
		)
	}
}

func (bar *Bar) Animate() {
	for {
		select {
		case <-bar.end:
			fmt.Printf("\n")
			return
		case <-bar.increment:
			// TODO: THIS IS WHERE WE CAN ENSURE IT NEVER GOES OVER WIDTH LIMIT
			fmt.Printf("%v", bar.Frame())
		case <-bar.ticker.C:
			fmt.Printf("%v", bar.Frame())
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
