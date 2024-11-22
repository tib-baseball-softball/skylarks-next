package stats

import (
	"github.com/tib-baseball-softball/skylarks-next/enum"
)

type PersonalAttendanceStatsItem struct {
	Season              int                          `json:"season"`
	TotalPossibleEvents int                          `json:"totalPossibleEvents"`
	AttendanceTotals    []AttendanceTotal            `json:"attendanceTotals"`
	ParticipationTotals []ParticipationTotal         `json:"participationTotals"`
	Type                enum.EventType               `json:"type"`
	Values              []ParticipationStatsByPerson `json:"values"`
	TeamName            string                       `json:"teamName"`
}

type AttendanceTotal struct {
	Type     enum.EventType `json:"type"`
	Attended int            `json:"attended"`
	Total    int            `json:"total"`
}

type ParticipationTotal struct {
	Type   enum.ParticipationType `json:"type"`
	Amount int                    `json:"amount"`
}
