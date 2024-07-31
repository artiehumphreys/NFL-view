package models

import "fmt"

type GameDisplay struct {
	PlayID    string
	FirstName string
	LastName  string
}

func (id GameDisplay) String() string {
	return fmt.Sprintf("%s %s %s",
		id.PlayID, id.FirstName, id.LastName)
}
