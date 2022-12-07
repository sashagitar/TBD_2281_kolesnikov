package list

type Date struct {
	day   int
	month int
	year  int
}

type Timer struct {
	start    Date
	finish   Date
	interval int
}

type Produkt struct {
	name   string
	weight float64
	bought bool
	Timer
}
