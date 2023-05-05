package dependencyinjection

type cheeze struct {
	name string
}

func NewCheeze() ICheeze {
	return &cheeze{
		name: "cheeze name",
	}
}

func (c *cheeze) GetName() string {
	return c.name
}
