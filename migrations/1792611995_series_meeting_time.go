package migrations

import (
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp"
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId(dp.EventSeriesCollection)
		if err != nil {
			return err
		}

		meetingField := core.NumberField{
			Name:        "meet",
			Id:          "meet",
			System:      false,
			Hidden:      false,
			Presentable: false,
			Help:        "Time (in minutes) to meet before the start of each event",
			Min:         nil,
			Max:         nil,
			OnlyInt:     true,
			Required:    false,
		}

		collection.Fields.Add(&meetingField)

		err = app.Save(collection)
		if err != nil {
			return err
		}

		return nil
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId(dp.EventSeriesCollection)
		if err != nil {
			return err
		}
		collection.Fields.RemoveByName("meet")

		err = app.Save(collection)
		if err != nil {
			return err
		}

		return nil
	})
}
