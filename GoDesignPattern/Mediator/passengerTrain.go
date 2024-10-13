package mediator

import "fmt"

type PassengerTrain struct {
	mediator Mediator
}

func (t *PassengerTrain) arrive() {
	if !t.mediator.canArrive(t) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (t *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	t.mediator.notifyAboutDeparture()
}

func (t *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	t.arrive()
}
