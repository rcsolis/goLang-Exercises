package nhlapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type division struct{
	ID int `json:"id"`
	Name string	`json:"name"`
	Link string	`json:"link"`
}
type timeZone struct{
	ID string	`json:"id"`
	Offset int	`json:"offset"`
	Tz string	`json:"tz"`
}
type venue struct {
	Name string `json:"name"`
	Link string `json:"link"`
	City string `json:"city"`
	TimeZone timeZone `json:"timeZone"`
}
type conference struct {
	ID int	`json:"id"`
	Name string	`json:"name"`
	Link string	`json:"link"`
}
type franchise struct {
	FranchiseID int    `json:"franchiseId"`
	TeamName string `json:"teamName"`
	Link string `json:"link"`
}
// Team data model for store information about a NHL team.
type Team struct{
	ID int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
	ShortName string `json:"shortName"`
	OfficialSiteURL string `json:"officialSiteUrl"`
	FranchiseID int    `json:"franchiseId"`
	Active bool   `json:"active"`
	Venue venue `json:"venue,omitempty"`
	Abbreviation string `json:"abbreviation"`
	TeamName string `json:"teamName"`
	LocationName string `json:"locationName"`
	FirstYearOfPlay string `json:"firstYearOfPlay"`
	Division division `json:"division"`
	Conference conference `json:"conference"`
	Franchise franchise `json:"franchise"`
}
// GetRoster Method of Team to fetch the players list
func (t *Team)GetRoster() ([]Player, error){
	// Send a request for the roster
	res, err := http.Get(fmt.Sprintf("%s/teams/%d/roster", baseURL, t.ID))
	// If error return 
	if err != nil{
		return nil, err
	}
	// Ensure to close the resource
	defer res.Body.Close()
	// Parse the response with the container
	var response nhlRosterResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	// Return the roster or error
	if err != nil {
		return nil, err
	}
	return response.Roster, nil
}
