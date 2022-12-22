package models

type PlayerYearlyPayload struct {
	Stats []Stats `json:"stats"`
}

type Stats struct {
	// TODO: Implement this stats type struct
	Type   interface{} `json:"type"`
	Splits []Splits    `json:"splits"`
}

type Splits struct {
	Season string     `json:"season"`
	Stat   PlayerStat `json:"stat"`
}

type PlayerStat struct {
	TOI                        string  `json:"timeOnIce"`
	Assists                    int     `json:"assists"`
	Goals                      int     `json:"goals"`
	PIM                        int     `json:"pim"`
	SOG                        int     `json:"shots"`
	Games                      int     `json:"games"`
	Hits                       int     `json:"hits"`
	PPG                        int     `json:"powerPlayGoals"`
	PPP                        int     `json:"powerPlayPoints"`
	PPTOI                      string  `json:"powerPlayTimeOnIce"`
	EvenTOI                    string  `json:"evenTimeOnIce"`
	PenaltyMinutes             string  `json:"penaltyMinutes"`
	FaceOffPercentage          float64 `json:"faceOffPct"`
	ShotPercentage             float64 `json:"shotPct"`
	GWG                        int     `json:"gameWinningGoals"`
	OTG                        int     `json:"overTimeGoals"`
	SHG                        int     `json:"shortHandedGoals"`
	SHP                        int     `json:"shortHandedPoints"`
	SHTOI                      string  `json:"shortHandedTimeOnIce"`
	Blocks                     int     `json:"blocked"`
	PlusMinus                  int     `json:"plusMinus"`
	Points                     int     `json:"points"`
	Shifts                     int     `json:"shifts"`
	AvgEvenTOI                 string  `json:"evenTimeOnIcePerGame"`
	AvgShortHandedTOI          string  `json:"shortHandedTimeOnIcePerGame"`
	AvgPowerPlayTOI            string  `json:"powerPlayTimeOnIcePerGame"`
	OT                         int     `json:"ot"`
	Shutouts                   int     `json:"shutouts"`
	Ties                       int     `json:"ties"`
	Wins                       int     `json:"wins"`
	Losses                     int     `json:"losses"`
	Saves                      int     `json:"saves"`
	PowerPlaySaves             int     `json:"powerPlaySaves"`
	ShortHandedSaves           int     `json:"shortHandedSaves"`
	EvenSaves                  int     `json:"evenSaves"`
	ShortHandedShots           int     `json:"shotHandedShots"`
	EvenShots                  int     `json:"evenShots"`
	PowerPlayShots             int     `json:"powerPlayShots"`
	SavePercentage             float64 `json:"savePercentage"`
	GoalsAgainstAverage        float64 `json:"goalsAgainstAverage"`
	GamesStarted               int     `json:"gamesStarted"`
	ShotsAgainst               int     `json:"shotsAgainst"`
	GoalsAgainst               int     `json:"goalsAgainst"`
	TimeOnIcePerGame           string  `json:"timeOnIcePerGame"`
	PowerPlaySavePercentage    float64 `json:"powerPlaySavePercentage"`
	ShortHandedSavePercentage  float64 `json:"shortHandedSavePercentage"`
	EvenStrengthSavePercentage float64 `json:"evenStrengthSavePercentage"`
}

type PlayerStatYear struct {
	Player     Person
	YearSeason string
	Splits     Splits
}
