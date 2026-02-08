package dp

import "github.com/pocketbase/pocketbase/core"

const (
	TeamsCollection = "teams"
)

var _ core.RecordProxy = (*Team)(nil)

// Team RecordProxy for collection `teams`.
// Provides type-safe struct access to team records.
type Team struct {
	core.BaseRecordProxy
}

func (t *Team) CollectionName() string {
	return TeamsCollection
}

func (t *Team) ID() string {
	return t.Id
}

func (t *Team) Name() string {
	return t.GetString("name")
}

func (t *Team) SetName(name string) {
	t.Set("name", name)
}

func (t *Team) BSMLeagueGroup() int {
	return t.GetInt("bsm_league_group")
}

func (t *Team) SetBSMLeagueGroup(id int) {
	t.Set("bsm_league_group", id)
}

func (t *Team) AgeGroup() string {
	return t.GetString("age_group")
}

func (t *Team) SetAgeGroup(ageGroup string) {
	t.Set("age_group", ageGroup)
}

func (t *Team) Club() string {
	return t.GetString("club")
}

func (t *Team) SetClub(clubID string) {
	t.Set("club", clubID)
}

func (t *Team) Description() string {
	return t.GetString("description")
}

func (t *Team) SetDescription(desc string) {
	t.Set("description", desc)
}

func (t *Team) Admins() []string {
	return t.GetStringSlice("admins")
}

func (t *Team) SetAdmins(adminIDs []string) {
	t.Set("admins", adminIDs)
}

func (t *Team) SignupKey() string {
	return t.GetString("signup_key")
}

func (t *Team) SetSignupKey(key string) {
	t.Set("signup_key", key)
}
