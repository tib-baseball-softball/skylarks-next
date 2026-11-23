with mq as (select
--p.id,
--p.event,
e.title,
e.type,
e.starttime,
p.state,
p.created                                              as participation_ts,
--p.user,
concat(u.first_name, ' ', u.last_name)                 as Name,

julianday(e.starttime) - julianday(p.created)          as time_diff_day,
((julianday(e.starttime) - julianday(p.created)) * 24) as time_diff_hour,
case
    when ((julianday(e.starttime) - julianday(p.created)) * 24) < 0 then 'Time travelling alert!'
    when ((julianday(e.starttime) - julianday(p.created)) * 24) >= 0 and
         ((julianday(e.starttime) - julianday(p.created)) * 24) <= 12 then 'Naughty alert!'
    when ((julianday(e.starttime) - julianday(p.created)) * 24) > 12 and
         ((julianday(e.starttime) - julianday(p.created)) * 24) <= 24 and p.state != 'in'
        then 'I am not angry, just dissapointed'
    when ((julianday(e.starttime) - julianday(p.created)) * 24) > 12 and
         ((julianday(e.starttime) - julianday(p.created)) * 24) <= 24 and p.state = 'in'
        then 'I am glad you could make it'
    when ((julianday(e.starttime) - julianday(p.created)) * 24) > 24 and
         julianday(e.starttime) - julianday(p.created) < 5 then 'Thank you for being you!'
    when julianday(e.starttime) - julianday(p.created) >= 5 then 'Look at you go - you are a planning wizard!'
    end                                                as participation_grade

            from participations p
                     left join users u on p.user = u.id
                     left join events e on p.event = e.id
            )

select
Name,
type,
state,
participation_grade,
count(participation_ts) as count

from mq

where state = 'out'
and participation_grade = 'Naughty alert!'
and type = 'practice'
and starttime >= '2026-01-01'
group by Name, type, state, participation_grade
order by count(participation_ts) desc