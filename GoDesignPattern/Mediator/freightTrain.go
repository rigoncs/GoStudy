package mediator

import "fmt"

type FreightTrain struct {
	mediator Mediator
}

func (t *FreightTrain) arrive() {
	if !t.mediator.canArrive(t) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (t *FreightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	t.mediator.notifyAboutDeparture()
}

func (t *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted, arriving")
	t.arrive()
}
