package factorymethod

type Team struct {
	name string
	game string
}

func (t *Team) SetName(name string) {
	t.name = name
}

func (t *Team) GetName() string {
	return t.name
}

func (t *Team) SetGame(game string) {
	t.game = game
}

func (t *Team) GetGame() string {
	return t.game
}
