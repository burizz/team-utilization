package teams

// ItgixTeam - list of devops engineers within a team
type ItgixTeam struct {
	Team      string           `json:"Team"`
	Engineers []DevOpsEngineer `json:"Engineers"`
}

// DevOpsEngineer - engineer definition
type DevOpsEngineer struct {
	Firstname     string `json:"Firstname"`
	Lastname      string `json:"Lastname"`
	Level         string `json:"Level"`
	TrackedHours  int    `json:"TrackedHours"`
	TrackingMonth string `json:"TrackingMonth"`
	TrackingYear  int    `json:"TrackingYear"`
}
