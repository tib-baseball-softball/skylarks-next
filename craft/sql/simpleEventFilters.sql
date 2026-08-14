-- single team

SELECT events.title, events.starttime, events.bsm_id
FROM events
WHERE team = 'lz468bc0ce2zvfk'

UNION

SELECT events.title, events.starttime, events.bsm_id
FROM events, json_each(events.additional_teams)
WHERE json_each.value = 'lz468bc0ce2zvfk'
;

-- same thing, but different

SELECT events.title, events.starttime, events.bsm_id
FROM events
WHERE team = 'lz468bc0ce2zvfk'
   OR (SELECT 1 FROM json_each(events.additional_teams) WHERE value = 'lz468bc0ce2zvfk')
;