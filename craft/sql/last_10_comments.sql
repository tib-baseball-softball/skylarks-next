-- last 10 comments per team/club

select
                concat(u.first_name, ' ', u.last_name) as Name,
                --u.teams,
                --t.id as team_id,

                c.text,


                date(c.created) as comment_date, --leave as ts if you prefer
                --c.id as comment_id,
                --c.event as event_id,
                e.type as event_type,
                e.title as event_title,
                date(e.starttime) as event_date,
                --e.team as event_team_id,
                t.name as team_name,
                clubs.name as club_name
            from comments c
            left join events e on e.id = c.event
            left join teams t on e.team = t.id
            left join users u on c.user = u.id
            --left join teams t on (SELECT 1 FROM json_each(u.teams) WHERE value = t.id)
            left join clubs on t.club = clubs.id

where
    clubs.id = 'ww1yx36pohg0ap4' --by club id
    --e.team = 'qq9o8gru7npr82u' --by event team
    --u.id = 'rvm43s2gpapxyo1' --by user id
--and e.type = 'practice' --add this for limiting event types
order by c.created desc
limit 10