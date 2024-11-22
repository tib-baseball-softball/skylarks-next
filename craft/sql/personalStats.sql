SELECT users.id,
       users.last_name,
       users.first_name,
       events.type,
       COUNT(CASE WHEN participations.state = 'in' THEN participations.id END)    AS inCount,
       COUNT(CASE WHEN participations.state = 'out' THEN participations.id END)   AS outCount,
       COUNT(CASE WHEN participations.state = 'maybe' THEN participations.id END) AS maybeCount,
       COUNT(participations.id)                                                   AS TotalCount
FROM users
         LEFT JOIN participations ON participations.user = users.id
         INNER JOIN events ON participations.event = events.id
WHERE 1 = 1
  AND users.id = 'USER'
GROUP BY users.id, users.last_name, users.first_name, events.type;
