package state

import "log"

func main() {
	v := newVendingMachine(1, 10)
	err := v.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = v.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = v.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = v.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = v.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = v.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = v.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
