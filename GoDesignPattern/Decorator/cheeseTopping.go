package decorator

type CheeseTopping struct {
	pizza IPizza
}

func (this *CheeseTopping) getPrice() int {
	return this.pizza.getPrice() + 10
}
