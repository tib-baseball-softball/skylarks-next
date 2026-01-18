package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(16, []byte(`{
			"hidden": false,
			"id": "elzgs2s7",
			"maxSelect": 1,
			"name": "bats",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"none",
				"left",
				"right",
				"switch"
			]
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(17, []byte(`{
			"hidden": false,
			"id": "gahyjvnk",
			"maxSelect": 1,
			"name": "throws",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"none",
				"left",
				"right",
				"switch"
			]
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(19, []byte(`{
			"hidden": false,
			"id": "wqhoic1y",
			"maxSelect": 1,
			"name": "umpire",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"none",
				"A",
				"B",
				"C",
				"D"
			]
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(20, []byte(`{
			"hidden": false,
			"id": "0plahfdr",
			"maxSelect": 1,
			"name": "scorer",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"none",
				"A",
				"B",
				"C",
				"D"
			]
		}`)); err != nil {
			return err
		}

		err = app.Save(collection)
		if err != nil {
			return err
		}

		_, err = app.DB().NewQuery(`
UPDATE users SET bats = 'none' WHERE bats = '0';
UPDATE users SET bats = 'left' WHERE bats = '1';
UPDATE users SET bats = 'right' WHERE bats = '2';
UPDATE users SET bats = 'switch' WHERE bats = '3';

UPDATE users SET throws = 'none' WHERE throws = '0';
UPDATE users SET throws = 'left' WHERE throws = '1';
UPDATE users SET throws = 'right' WHERE throws = '2';
UPDATE users SET throws = 'switch' WHERE throws = '3';

UPDATE users SET umpire = 'none' WHERE umpire = '0';
UPDATE users SET umpire = 'A' WHERE umpire = '1';
UPDATE users SET umpire = 'B' WHERE umpire = '2';
UPDATE users SET umpire = 'C' WHERE umpire = '3';
UPDATE users SET umpire = 'D' WHERE umpire = '4';

UPDATE users SET scorer = 'none' WHERE scorer = '0';
UPDATE users SET scorer = 'A' WHERE scorer = '1';
UPDATE users SET scorer = 'B' WHERE scorer = '2';
UPDATE users SET scorer = 'C' WHERE scorer = '3';
UPDATE users SET scorer = 'D' WHERE scorer = '4';
`).Execute()

		if err != nil {
			return err
		}

		return nil
	}, func(app core.App) error {
		// ROLLBACK
		collection, err := app.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		if err := collection.Fields.AddMarshaledJSONAt(16, []byte(`{
			"hidden": false,
			"id": "elzgs2s7",
			"maxSelect": 1,
			"name": "bats",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"0",
				"1",
				"2",
				"3"
			]
		}`)); err != nil {
			return err
		}

		if err := collection.Fields.AddMarshaledJSONAt(17, []byte(`{
			"hidden": false,
			"id": "gahyjvnk",
			"maxSelect": 1,
			"name": "throws",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"0",
				"1",
				"2",
				"3"
			]
		}`)); err != nil {
			return err
		}

		if err := collection.Fields.AddMarshaledJSONAt(19, []byte(`{
			"hidden": false,
			"id": "wqhoic1y",
			"maxSelect": 1,
			"name": "umpire",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"0",
				"1",
				"2",
				"3",
				"4"
			]
		}`)); err != nil {
			return err
		}

		if err := collection.Fields.AddMarshaledJSONAt(20, []byte(`{
			"hidden": false,
			"id": "0plahfdr",
			"maxSelect": 1,
			"name": "scorer",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"0",
				"1",
				"2",
				"3"
			]
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
