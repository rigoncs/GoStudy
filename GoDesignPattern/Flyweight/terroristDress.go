package flyweight

type TerroristDress struct {
	color string
}

func (d *TerroristDress) getColor() string {
	return d.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "Red"}
}
