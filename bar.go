package loading

import (
	"fmt"
	"math"
	"strings"
	"time"

	width "golang.org/x/text/width"
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

	//_, runeWidth := utf8.DecodeRuneInString(animationFrames["fill"][0])

	//runeProperties := width.LookupRune(bar.animation.Unfilled)

	//rune, size := utf8.DecodeRune([]byte(bar.animation.Unfilled))

	//runeProperties := width.LookupRune(bar.animation.Unfilled)

	_, runeWidth := width.LookupString(animationFrames["fill"][0])
	fmt.Printf("runeWidth(%v)\n", runeWidth)

	bar := &Bar{
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
	fmt.Printf("bar.ShowPercent(%v)\n", show)
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
	fill := strings.Repeat(
		bar.frames["fill"][len(bar.frames["fill"])-1],
		int(uint(bar.progress)/bar.runeWidth),
	)
	return fill
}

func (bar Bar) unfilled() string {

	return strings.Repeat(
		bar.frames["unfilled"][0],
		int(math.Floor(float64(bar.RemainingTicks()/bar.runeWidth))),
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

	// TODO
	// Add optional spinner animation at end of the loading bar
	fill := bar.filled() + bar.frames["fill"][bar.animationTick]
	if bar.animationTick < len(bar.frames["fill"])-1 {
		bar.animationTick += 1
	} else {
		bar.animationTick = 0
	}

	if bar.percentVisible {
		return fmt.Sprintf(
			bar.format,
			fill,
			bar.unfilled(),
			bar.percent(),
			bar.status,
		)
	} else {
		return fmt.Sprintf(
			bar.format,
			fill,
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
			fmt.Print(bar.Frame())
		case <-bar.ticker.C:
			fmt.Printf(bar.Frame())
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
