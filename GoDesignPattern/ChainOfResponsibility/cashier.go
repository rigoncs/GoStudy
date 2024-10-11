package chainofresponsibility

import "fmt"

type Cashier struct {
	next DepartMent
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
}

func (c *Cashier) setNext(next DepartMent) {
	c.next = next
}
