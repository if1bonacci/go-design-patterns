package factorymethod

type ITeam interface {
	SetName(name string)
	GetName() string
	SetGame(game string)
	GetGame() string
}
