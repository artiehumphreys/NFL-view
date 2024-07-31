package models

import "fmt"

type GameDisplay struct {
	Game      string
	PlayID    string
	Team      string
	FirstName string
	LastName  string
}

func (id GameDisplay) String() string {
	return fmt.Sprintf("%s %s %s %s",
		id.PlayID, id.FirstName, id.LastName, id.Team)
}
