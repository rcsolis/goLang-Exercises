package main

import (
	"io"
	"log"
	"nhlproject/nhlapi"
	"os"
	"runtime"
	"sync"
	"time"
)

// loadRoster is the Sender function to push into the channel a roster slice
// for a team identified by its ID
func loadRoster(team nhlapi.Team, ch chan <- nhlapi.FullTeamData, wg *sync.WaitGroup){
	// Marks as done when it finish
	defer wg.Done()
	// Request for the team roster
	rosterArr, err := team.GetRoster()
	// if an error occurs fetching the roster
	if err != nil{
		log.Fatalf("Error getting roster of team #%d. %v", team.ID, err)
	}
	// Send the roster to the channel
	ch <- nhlapi.FullTeamData{Team: team, Roster: rosterArr}
}

// displayRoster Its the Receiver function to listen the channel for a roster slice
// it waits until the channel will close
func displayRoster(ch <- chan nhlapi.FullTeamData){
	// Listen the channel for each team data
	for teamData := range ch{
		log.Printf("****** Recieving information about #%d %s and its roster with %d players", teamData.Team.ID, teamData.Team.ShortName, len(teamData.Roster))
		for _,player := range teamData.Roster{
			log.Println(player.String())
		}
	}
}

// Main
func main(){
	// Help to benchmarking the request and processing time
	startTime := time.Now()
	// Open a file to write and save the response
	teamsFile, err := os.OpenFile("teams.txt", os.O_RDWR|os.O_CREATE, 0666)
	// If can not open
	if err != nil{
		log.Fatalf("Error opening teams.txt: %v",err)
	}
	// Ensure to close the file at the end
	defer teamsFile.Close()
	// Create a multi-writer to send data and write 
	// to the standard output (console) and to the text file
	wrt := io.MultiWriter(os.Stdout, teamsFile)
	// Change the output for the logger
	log.SetOutput(wrt)
	// Calls to our function to fetch the API
	teams, err := nhlapi.GetAllTeams()
	// If error occurs
	if err != nil {
		log.Fatal("Error while getting teams.", err)
	}
	// For each team
	for _, team := range teams{
		log.Println("------ TEAM ------")
		log.Printf("Id: %v \n", team.ID)
		log.Printf("Name: %s \n", team.Name)
		log.Printf("Short name: %s \n", team.ShortName)
		log.Println("------------------")
	}
	log.Printf("\n ++++++ %d Teams found ++++++ \n", len(teams))
	// ROSTERS FOR EACH TEAM
	// Set the cores availables to use for the processing
	runtime.GOMAXPROCS(runtime.NumCPU())
	//Declare a wait group and a channel
	var wg sync.WaitGroup
	dataChannel := make(chan nhlapi.FullTeamData)
	// Add Gorutines to the Wait Group to process all teams
	wg.Add(len(teams))
	// Execute go rutines for each team
	for _, team := range teams {
		go loadRoster(team, dataChannel, &wg)
	}
	// Execute a Go rutine that blocks itself waiting to finish all gorutines
	go func(){
		wg.Wait()
		// Close channel to block sends
		close(dataChannel)
	}()
	//Recieve data and blocks execution until channel is empty
	displayRoster(dataChannel)
	// Show the execution time
	log.Printf("\n Init at %v and Finish at %v Execution time were %v", startTime.String(),time.Now().String(),time.Now().Sub(startTime).String())
}