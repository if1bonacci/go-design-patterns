package factorymethod

import "fmt"

func GetTeam(teanType string) (ITeam, error) {
	if teanType == "kalmar" {
		return NewKalmar(), nil
	}
	if teanType == "barsa" {
		return NewBarselona(), nil
	}
	return nil, fmt.Errorf("Wrong team type passed")
}
