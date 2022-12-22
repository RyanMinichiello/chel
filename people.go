package chel

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ryanminichiello/chel/config"
	"github.com/ryanminichiello/chel/models"
)

type PeopleService struct {
	client    *http.Client
	apiConfig config.Config
}

// GetPerson retrieves associated Person from NHLAPI
// Makes no assumption about id's existence in NHLAPI
// Returns models.Person and any error encountered
func (ps *PeopleService) GetPerson(id int) (models.Person, error) {
	var person models.Person
	url := fmt.Sprintf("%s%s%d", ps.apiConfig.Host, ps.apiConfig.People, id)
	resp, err := ps.client.Get(url)
	if err != nil {
		return person, err
	}
	defer resp.Body.Close()

	var people models.People
	err = json.NewDecoder(resp.Body).Decode(&people)
	if err != nil {
		return person, err
	}

	if len(people.People) < 1 {
		return person, fmt.Errorf("no data found")
	}

	person = people.People[0]

	return person, nil
}

// GetPersonStatYear retieves the professional statistics for playerID from NHLAPI
// yearSeason should be 8 digit string, "20192020" would be the NHL's 2019 calendar season
// Makes no assumption about playerID's existence or valid yearSeason for NHLAPI
// Returns a models.PlayerStatYear and any error encountered.
func (ps *PeopleService) GetPlayerStatYear(playerID int, yearSeason string) (models.PlayerStatYear, error) {
	player, err := ps.GetPerson(playerID)
	if err != nil {
		return models.PlayerStatYear{}, err
	}

	url := fmt.Sprintf("%v%v/%v/stats?stats=statsSingleSeason&season=%v", ps.apiConfig.Host, ps.apiConfig.People, playerID, yearSeason)
	resp, err := ps.client.Get(url)
	if err != nil {
		return models.PlayerStatYear{}, err
	}
	defer resp.Body.Close()

	var pyPayload models.PlayerYearlyPayload
	err = json.NewDecoder(resp.Body).Decode(&pyPayload)
	if err != nil {
		return models.PlayerStatYear{}, err
	}
	if len(pyPayload.Stats) < 1 {
		return models.PlayerStatYear{}, fmt.Errorf("empty stats payload for player: %v", playerID)
	}

	playerStatYear := models.PlayerStatYear{
		Player:     player,
		YearSeason: yearSeason,
		Splits:     pyPayload.Stats[0].Splits[0],
	}

	return playerStatYear, nil
}
