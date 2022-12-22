package models

type People struct {
	People []Person `json:"people"`
}
type Person struct {
	ID               int      `json:"id"`
	FullName         string   `json:"fullName"`
	Link             string   `json:"link"`
	FirstName        string   `json:"firstName"`
	LastName         string   `json:"lastName"`
	PrimaryNumber    string   `json:"primaryNumber"`
	BirthDate        string   `json:"birthDate"`
	CurrentAge       int      `json:"currentAge"`
	BirthCity        string   `json:"birthCity"`
	BirthCountry     string   `json:"birthCountry"`
	Nationality      string   `json:"nationality"`
	Height           string   `json:"height"`
	Weight           int      `json:"weight"`
	Active           bool     `json:"active"`
	AlternateCaptain bool     `json:"alternateCaptain"`
	Captain          bool     `json:"captain"`
	Rookie           bool     `json:"rookie"`
	ShootsCatches    string   `json:"shootsCatches"`
	RosterStatus     string   `json:"rosterStatus"`
	CurrentTeam      Team     `json:"currentTeam"`
	PrimaryPosition  Position `json:"primaryPosition"`
}
