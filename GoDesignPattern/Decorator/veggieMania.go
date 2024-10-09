package decorator

type VeggieMania struct{}

func (this *VeggieMania) getPrice() int {
	return 15
}
