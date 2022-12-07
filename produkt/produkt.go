package produkt

type Date struct {
	Day   int
	Month int
	Year  int
}

type Timer struct {
	Start    Date
	Finish   Date
	Interval int
}

type Produkt struct {
	Name   string
	Weight float64
	Bought bool
	Timer
}
