package factorymethod

type barselona struct {
	Team
}

func NewBarselona() ITeam {
	return &barselona{
		Team: Team{
			name: "FC BARCELONA",
			game: "football",
		},
	}
}
