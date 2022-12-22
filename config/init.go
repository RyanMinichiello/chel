package config

func InitConfig() (Config, error) {
	cfg := &Config{
		Host:             "https://statsapi.web.nhl.com/api/v1/",
		Franchises:       "franchises/",
		Teams:            "teams/",
		Divisions:        "divisions/",
		Conferences:      "conferences/",
		People:           "people/",
		Game:             "game/",
		GameType:         "gameTypes/",
		GameStatus:       "gameStatus/",
		PlayTypes:        "playTypes/",
		TournamentTypes:  "tournamentTypes/",
		PlayoffStructure: "playoffStructure/",
		Schedule:         "schedule/",
		Seasons:          "seasons/",
		Standings:        "standings/",
		StandingsType:    "standingsType/",
		StatTypes:        "statTypes/",
		TeamStats:        "teamStats/",
		Draft:            "draft/",
		Prospects:        "prospects/",
		Awards:           "awards/",
		Venues:           "venues/",
		EventTypes:       "eventTypes/",
		Roster:           "roster/",
	}
	return *cfg, nil
}

type Config struct {
	Host             string `json:"host"`
	Franchises       string `json:"franchises"`
	Teams            string `json:"teams"`
	Divisions        string `json:"divisions"`
	Conferences      string `json:"conferences"`
	People           string `json:"people"`
	Game             string `json:"game"`
	GameType         string `json:"gameType"`
	GameStatus       string `json:"gameStatus"`
	PlayTypes        string `json:"playTypes"`
	TournamentTypes  string `json:"tournamentTypes"`
	PlayoffStructure string `json:"playoffStructure"`
	Schedule         string `json:"schedule"`
	Seasons          string `json:"seasons"`
	Standings        string `json:"standings"`
	StandingsType    string `json:"standingsType"`
	StatTypes        string `json:"statTypes"`
	TeamStats        string `json:"teamStats"`
	Draft            string `json:"draft"`
	Prospects        string `json:"prospects"`
	Awards           string `json:"awards"`
	Venues           string `json:"venues"`
	EventTypes       string `json:"eventTypes"`
	Roster           string `json:"roster"`
}
