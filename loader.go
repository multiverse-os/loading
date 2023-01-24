package loading

// Interface for both Spinner & Bar
type Loader interface {
	Start()
	// TODO: Build on this and make the two types similar enough that we can use
	// this loader interface instead of typing out to their type
	Increment(float64) bool
	End()
}

func ToBar(loader Loader) *Bar         { return loader.(*Bar) }
func ToSpinner(loader Loader) *Spinner { return loader.(*Spinner) }
