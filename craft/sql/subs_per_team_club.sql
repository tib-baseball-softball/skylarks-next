-- Team 2

SELECT sub.id, sub.endpoint, users.email
FROM pushsubscriptions AS sub
         INNER JOIN users ON sub.user = users.id
WHERE EXISTS (SELECT 1 FROM json_each(users.teams) WHERE value = 'qq9o8gru7npr82u')
;

-- Bogus Team

SELECT sub.id, sub.endpoint, users.email
FROM pushsubscriptions AS sub
         INNER JOIN users ON sub.user = users.id
WHERE EXISTS (SELECT 1 FROM json_each(users.teams) WHERE value = 'xkgyhs024uup2l0')
;

-- München

SELECT sub.id, sub.endpoint, users.email
FROM pushsubscriptions AS sub
         INNER JOIN users ON sub.user = users.id
--INNER JOIN teams ON teams.id = 'w945fqi7pjk0nuc'
WHERE EXISTS (SELECT 1 FROM json_each(users.teams) WHERE value = 'w945fqi7pjk0nuc')
;
