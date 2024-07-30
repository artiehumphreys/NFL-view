package models

import "fmt"

type InjuryDisplay struct {
	Type         string
	GamePosition string
	Team         string
	JerseyNumber string
	FirstName    string
	LastName     string
}

func (id InjuryDisplay) String() string {
	return fmt.Sprintf("%s %s: %s %s #%s %s",
		id.FirstName, id.LastName, id.GamePosition, id.Team, id.JerseyNumber, id.Type)
}
