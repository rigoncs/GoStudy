package decorator

type TomatoTopping struct {
	pizza IPizza
}

func (this *TomatoTopping) getPrice() int {
	return this.pizza.getPrice() + 7
}
