package dp

import "git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/base"

type PersonalAttendanceStatsItem struct {
	Season              int                          `json:"season"`
	TotalPossibleEvents int                          `json:"totalPossibleEvents"`
	AttendanceTotals    []AttendanceTotal            `json:"attendanceTotals"`
	ParticipationTotals []ParticipationTotal         `json:"participationTotals"`
	Type                base.EventType               `json:"type"`
	Values              []ParticipationStatsByPerson `json:"values"`
	TeamName            string                       `json:"teamName"`
}

type AttendanceTotal struct {
	Type     base.EventType `json:"type"`
	Attended int            `json:"attended"`
	Total    int            `json:"total"`
}

type ParticipationTotal struct {
	Type   ParticipationType `json:"type"`
	Amount int               `json:"amount"`
}
