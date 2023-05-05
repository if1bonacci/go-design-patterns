package chainofresponsibility

type Step interface {
	Execute(*Order)
	SetNext(Step)
}
