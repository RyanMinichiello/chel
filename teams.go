package chel

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ryanminichiello/chel/config"
	"github.com/ryanminichiello/chel/models"
)

type TeamService struct {
	client    *http.Client
	apiConfig config.Config
}

// List retrieves all active NHL teams from NHLAPI
// Returns slice of []model.Team and any error encountered
func (ts *TeamService) List() ([]models.Team, error) {
	var teams []models.Team
	url := fmt.Sprintf("%s%s", ts.apiConfig.Host, ts.apiConfig.Teams)
	resp, err := ts.client.Get(url)
	if err != nil {
		return teams, err
	}
	defer resp.Body.Close()

	var x models.Teams
	err = json.NewDecoder(resp.Body).Decode(&x)
	if err != nil {
		return teams, err
	}

	teams = x.List
	return teams, nil
}

// GetTeam retrieves associated NHL team from NHLAPI
// Makes no assumption about teamID's existence in NHLAPI
// Returns models.Team and any error encountered
func (ts *TeamService) GetTeam(id int) (models.Team, error) {
	var team models.Team
	url := fmt.Sprintf("%s%s%d", ts.apiConfig.Host, ts.apiConfig.Teams, id)
	resp, err := ts.client.Get(url)
	if err != nil {
		return team, err
	}
	defer resp.Body.Close()

	var teams models.Teams
	err = json.NewDecoder(resp.Body).Decode(&teams)

	if err != nil {
		return team, err
	}
	team = teams.List[0]

	return team, nil
}

// GetRoster retrieves active roster for supplied NHL team by their ID
// Makes no assumption about teamID's existence in NHLAPI
// Returns models.Roster and any error encountered
func (ts *TeamService) GetRoster(id int) (models.Roster, error) {
	var roster models.Roster
	url := fmt.Sprintf("%s%s%d%s", ts.apiConfig.Host, ts.apiConfig.Teams, id, ts.apiConfig.Roster)
	resp, err := ts.client.Get(url)
	if err != nil {
		return roster, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&roster)
	if err != nil {
		return roster, err
	}
	return roster, nil
}

// Roster retrieves active roster for supplied NHL Team by their ID
// Similar to GetRoster but allows a Team struct to be passed in place
// Makes no assumption about teamID's existence in NHLAPI
// Returns models.Roster and any error encountered
func (ts *TeamService) Roster(t models.Team) (models.Roster, error) {
	var roster models.Roster
	url := fmt.Sprintf("%s%s%d%s", ts.apiConfig.Host, ts.apiConfig.Teams, t.ID, ts.apiConfig.Roster)
	resp, err := ts.client.Get(url)
	if err != nil {
		return roster, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&roster)
	if err != nil {
		return roster, err
	}
	return roster, nil
}
