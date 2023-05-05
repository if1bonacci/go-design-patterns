package factorymethod

type kalmar struct {
	Team
}

func NewKalmar() ITeam {
	return &kalmar{
		Team: Team{
			name: "FC Kalmar",
			game: "football",
		},
	}
}
