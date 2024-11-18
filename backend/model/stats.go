package model

import "github.com/tib-baseball-softball/skylarks-next/enum"

type PersonalAttendanceStatsItem struct {
	Season              int                  `json:"season"`
	AttendanceTotals    []AttendanceTotal    `json:"attendanceTotals"`
	ParticipationTotals []ParticipationTotal `json:"participationTotals"`
	TotalPossibleEvents int                  `json:"totalPossibleEvents"`
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
