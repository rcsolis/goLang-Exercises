/*
Package nhlapi implements an API to fetch NHL Teams
*/
package nhlapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ------------- Private Access

const baseURL="https://statsapi.web.nhl.com/api/v1"

// Container for the HTTP response JSON file for Teams
type nhlTeamsResponse struct{
	Teams []Team `json:"teams"`
}
// Container for the HTTP response JSON file for Roster
type nhlRosterResponse struct{
	Roster []Player `json:"roster"`
}

// ------------- Public Access

// FullTeamData base data model with information about team and its roster
type FullTeamData struct{
	Team Team
	Roster []Player
}

// GetAllTeams to fetch api and get the NHL teams.
// Returns all the teams as a slice or an error.
func GetAllTeams() ([]Team, error){
	// send a request to the API 
	res, err:= http.Get(
		fmt.Sprintf("%s/teams",baseURL))
	// if an error occurs when fetching data
	if err != nil {
		return nil, err
	}
	// ensure that the resourse is close
	defer res.Body.Close()
	// declare the response container
	var response nhlTeamsResponse
	// create a decoder and parse the response
	err = json.NewDecoder(res.Body).Decode(&response)
	// Return teams or error
	if err != nil {
		return nil, err
	}
	return response.Teams, nil
}