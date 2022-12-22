package models

type Teams struct {
	List []Team `json:"teams"`
}

type Team struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Link            string     `json:"link"`
	Venue           Venue      `json:"venue"`
	Abbreviation    string     `json:"abbreviation"`
	TeamName        string     `json:"teamName"`
	LocationName    string     `json:"locationName"`
	FirstYearOfPlay string     `json:"firstYearOfPlay"`
	Division        Division   `json:"division"`
	Conference      Conference `json:"conference"`
	Franchise       Franchise  `json:"franchise"`
	ShortName       string     `json:"shortName"`
	OfficialSiteURL string     `json:"officialSiteUrl"`
	FranshiseID     int        `json:"franchiseId"`
	Active          bool       `json:"active"`
}

type Venue struct {
	Name     string   `json:"name"`
	Link     string   `json:"link"`
	City     string   `json:"city"`
	TimeZone TimeZone `json:"timeZone"`
}

type TimeZone struct {
	TimeZoneId string `json:"id"`
	Offset     int    `json:"offset"`
	TZ         string `json:"tz"`
}

type Division struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	NameShort    string `json:"nameShort"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation"`
}

type Conference struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type Franchise struct {
	FranchiseID int    `json:"franchiseId"`
	TeamName    string `json:"teamName"`
	Link        string `json:"link"`
}
