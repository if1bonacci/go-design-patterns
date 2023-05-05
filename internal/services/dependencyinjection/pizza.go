package dependencyinjection

type pizza struct {
	cheeze ICheeze
}

func NewPizza(cheeze ICheeze) IPizza {
	return &pizza{
		cheeze: cheeze,
	}
}

func (p *pizza) GetCheeze() string {
	return p.cheeze.GetName()
}
