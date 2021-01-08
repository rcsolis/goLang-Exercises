package nhlapi

import "fmt"

type person struct {
	ID int `json:"id"`
	FullName string `json:"fullName"`
	Link     string `json:"link"`
}
type position struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	TypeOf         string `json:"type"`
	Abbreviation string `json:"abbreviation"`
}
// Player is the main model for represent a player of a roster
type Player struct {
	Person       person   `json:"person"`
	JerseyNumber string   `json:"jerseyNumber"`
	Position     position `json:"position"`
}
//String method to parse to print out as string
func (r *Player) String() string {
	str := fmt.Sprintf("Name:%s Position:%s (%s) Jersey:%s", r.Person.FullName, r.Position.Name,r.Position.Code,r.JerseyNumber)
	return "--->"+str
}

