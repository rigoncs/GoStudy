package bridge

func main() {
	epson := &Epson{}
	hp := &Hp{}
	mac := &Mac{}
	win := &Windows{}

	mac.SetPrinter(epson)
	mac.Print()

	win.SetPrinter(hp)
	win.Print()
}
