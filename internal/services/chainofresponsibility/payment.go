package chainofresponsibility

import "fmt"

type Payment struct {
	next Step
}

func (p *Payment) Execute(o *Order) {
	if o.PaymentDone {
		fmt.Println("Order payment already done")
		p.next.Execute(o)
		return
	}

	fmt.Println("Process of payment...")
	o.PaymentDone = true
	p.next.Execute(o)
}

func (p *Payment) SetNext(next Step) {
	p.next = next
}
