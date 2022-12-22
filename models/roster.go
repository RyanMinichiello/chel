package models

type Roster struct {
	PlayersList []Player `json:"roster"`
}

type Player struct {
	Person       Person   `json:"person"`
	JerseyNumber string   `json:"jerseyNumber"`
	Position     Position `json:"position"`
}

type Position struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Abbreviation string `json:"abbreviation"`
}
