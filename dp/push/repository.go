package push

import (
	"ganigeorgiev.com/pocketbase/dbx"
	"ganigeorgiev.com/pocketbase/pocketbase/core"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/base"
)

func GetSubscriptionsForTeamOrClub(id string, coll string, app core.App) ([]base.PushSubscription, error) {
	var whereClause dbx.Expression

	if coll == base.ClubsCollection {
		whereClause = dbx.Exists(dbx.NewExp("SELECT 1 FROM json_each(users.club) WHERE value = {:clubID}", dbx.Params{"clubID": id}))
	} else {
		whereClause = dbx.Exists(dbx.NewExp("SELECT 1 FROM json_each(users.teams) WHERE value = {:teamID}", dbx.Params{"teamID": id}))
	}

	var subs []base.PushSubscription
	err := app.RecordQuery(base.PushSubscriptionsCollection).
		InnerJoin("users", dbx.NewExp("pushsubscriptions.user = users.id")).
		AndWhere(whereClause).
		All(&subs)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func GetSubscriptionsForUserIDs(userIDs []string, app core.App) ([]base.PushSubscription, error) {
	ids := make([]any, len(userIDs))
	for i, id := range userIDs {
		ids[i] = id
	}

	var subs []base.PushSubscription
	err := app.RecordQuery(base.PushSubscriptionsCollection).
		AndWhere(dbx.In("user", ids...)).
		All(&subs)
	if err != nil {
		return nil, err
	}

	return subs, nil
}
