package migrations

import (
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp"
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId(dp.ParticipationsCollection)
		if err != nil {
			return err
		}

		userField := collection.Fields.GetByName("user").(*core.RelationField)
		userField.Required = true

		eventField := collection.Fields.GetByName("event").(*core.RelationField)
		eventField.Required = true

		stateField := collection.Fields.GetByName("state").(*core.SelectField)
		stateField.Required = true

		err = app.Save(collection)
		if err != nil {
			return err
		}

		return nil
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId(dp.ParticipationsCollection)
		if err != nil {
			return err
		}

		userField := collection.Fields.GetByName("user").(*core.RelationField)
		userField.Required = false

		eventField := collection.Fields.GetByName("event").(*core.RelationField)
		eventField.Required = false

		stateField := collection.Fields.GetByName("state").(*core.SelectField)
		stateField.Required = false

		err = app.Save(collection)
		if err != nil {
			return err
		}

		return nil
	})
}
