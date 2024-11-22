package stats

type ParticipationByPerson struct {
	ID         string `db:"id" json:"id"`
	LastName   string `db:"last_name" json:"lastName"`
	FirstName  string `db:"first_name" json:"firstName"`
	Type       string `db:"type" json:"type"`
	InCount    int    `db:"inCount" json:"inCount"`
	OutCount   int    `db:"outCount" json:"outCount"`
	MaybeCount int    `db:"maybeCount" json:"maybeCount"`
	TotalCount int    `db:"TotalCount" json:"totalCount"`
}
