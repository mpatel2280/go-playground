package challenges

// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// Define a new type
type ByPrice []Flight

func (f ByPrice) Len() int           { return len(f) }
func (f ByPrice) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ByPrice) Less(i, j int) bool { return f[i].Price < f[j].Price }
