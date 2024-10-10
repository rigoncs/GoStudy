package flyweight

type CounterTerroristDress struct {
	color string
}

func (d *CounterTerroristDress) getColor() string {
	return d.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "Green"}
}
