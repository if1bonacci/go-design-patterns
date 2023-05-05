package chainofresponsibility

type Order struct {
	Id           string
	DeliveryDone bool
	PaymentDone  bool
}
