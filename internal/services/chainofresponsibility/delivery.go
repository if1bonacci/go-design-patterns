package chainofresponsibility

import "fmt"

type Delivery struct {
	next Step
}

func (d *Delivery) Execute(o *Order) {
	if o.DeliveryDone {
		fmt.Println("Order delivery already done")
		return
	}

	fmt.Println("Process of delivery...")
	o.DeliveryDone = true
}

func (d *Delivery) SetNext(next Step) {
	d.next = next
}
