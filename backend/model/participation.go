package model

import "github.com/pocketbase/pocketbase/core"

type ParticipationsByType struct {
	In    []*core.Record `json:"in"`
	Out   []*core.Record `json:"out"`
	Maybe []*core.Record `json:"maybe"`
}
