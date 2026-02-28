package bsm

// LeagueGroup represents the BSM structure "Liga/Gruppe".
// It is always unique to single season.
// Caution: several LeagueGroup might have the same GameClass.
type LeagueGroup struct {
	ID                       int       `json:"id"`
	Season                   int       `json:"season"`
	Name                     string    `json:"name"`
	Acronym                  string    `json:"acronym"`
	Sport                    string    `json:"sport"`
	HumanSport               string    `json:"human_sport"`
	HumanSportShort          string    `json:"human_sport_short"`
	AgeGroup                 string    `json:"age_group"`
	HumanAgeGroup            string    `json:"human_age_group"`
	HumanAgeGroupShort       string    `json:"human_age_group_short"`
	Classification           string    `json:"classification"`
	HumanClassification      string    `json:"human_classification"`
	HumanClassificationShort string    `json:"human_classification_short"`
	PlannedInnings           int       `json:"planned_innings"`
	GameClass                GameClass `json:"game_class"`
}
