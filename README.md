<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">
## Multiverse OS: CLI Loading bars & spinners library
**URL** [multiverse-os.org](https://multiverse-os.org)
A simple and easy-to-use loading `bar` and `spinner` library, a subcomponent 
of the Multiverse OS `cli` library. Extendable, feature complete and can be
initialized, configred, and started in a single line and stopped in a second.


### Getting started
Keep in mind, that you do not need to initialize new spinner objects, you can
reuse the object, change the loading message and call start again.

### Loading Bars
The following example (from `examples/simple`) is included with the library to demonstrate one way to 
use the loading bars:

```
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/multiverse-os/loading"
	bar "github.com/multiverse-os/loading/bars/thinblocks"
)

func main() {
	fmt.Println("Loading Bar Example")
	fmt.Println("===================")
	fmt.Println("Simple example with one bar")

	loadingBar := loading.Bar().Width(80).Style(bar.Style).Start()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(rand.Intn(135)+22) * time.Millisecond)
		if loadingBar.Increment(1) {
			break
		}
	}
	loadingBar.Status("Completed!").Complete()
}

```


### Loading Spinners
A simple and easy-to-use spinner `library`, a subcomponent of the Multiverse OS 
cli library. Designed to be easily extensible, feature complete and easily 
initialized, configred, and started from a single line of code. 


#### Getting started
Keep in mind, that you do not need to initialize new spinner objects, you can 
reuse the object, change the loading message and call start again.

This example is included in the `/examples/` folder of the repository. 


```Go
package main

import (
	"fmt"
	"math/rand"
	"time"

	loading "github.com/multiverse-os/loading"
	lines "github.com/multiverse-os/loading/spinners/lines"
	moon "github.com/multiverse-os/loading/spinners/moon"
)

func main() {
	s := loading.Spinner(lines.Animation).Message("Loading...").Speed(loading.Normal)
	s.Start()
	start := time.Now()
	timer := time.NewTimer(15 * time.Second)
	tick := time.Tick(15 * time.Millisecond)
	for {
		select {
		case <-tick:
			s.Message(fmt.Sprintf("Loading for: %v", time.Since(start)))
		}
		if time.Since(start) > 15*time.Second {
			break
		}
	}
	<-timer.C
	timer.Stop()
	s.Complete(fmt.Sprintf("Completed after %v", time.Since(start)))
}
```

Options are defined using function chaining, and the available options
are:
  1) Speed(int), with aliases spinner.(Slowest,Slow,Normal,Fast,Fastest)
  2) Message(string)
  3) Palette([]string), expecting ansi colors to cycle through

Currently there are only six (6) available spinner options
any pull requests for additional spinners will be reviewed
and accepted if the code is consistent with the existing
library codebase. Pull requests welcome.
Loader Animation Options:
  1) circle.Animation   ["◐","◓", "◑", "◒"]
  2) clock.Animation    ["🕐","🕑","🕒","🕓",...]
  3) dots.Animation     ["⠋","⠙","⠹","⠸",...]
  4) lines.Animation    ["-","\","|","/"]
  5) moon.Animation     ["🌑","🌒","🌓","🌔",...]
  6) triangle.Animation ["◢","◣","◤","◥"]

To provide a more interactive UI, messages, speed and even palette can
be updated while the spinner animation is active to provide updates
which lets the user know the program is not frozen.
