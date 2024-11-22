SELECT COUNT(CASE WHEN events.type = 'game' THEN events.id END)     AS games,
       COUNT(CASE WHEN events.type = 'practice' THEN events.id END) AS practices,
       COUNT(CASE WHEN events.type = 'misc' THEN events.id END)     AS misc,
       COUNT(events.id)                                             AS all_events
FROM events
WHERE 1 = 1
  AND events.team = 'TEAM'
;