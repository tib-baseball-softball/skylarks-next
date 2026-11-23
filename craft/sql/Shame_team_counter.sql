with event as (select distinct e.id,
                               e.additional_teams,
                               e.type               as event_type,
                               e.title              as event_title,
                               date(e.starttime)    as event_date,
                               e.team               as event_team_id,
                               t.name               as team_name,
                               clubs.name           as club_name,
                               p.state,
                               COUNT(distinct p.id) AS state_count

               from events e
                        left join teams t on e.team = t.id
                        left join clubs on t.club = clubs.id
                        left join users u on (SELECT 1 FROM json_each(u.teams) WHERE value = t.id)
                        left join participations p on p.event = e.id

               group by e.id),

     team as
         (select distinct t.id,
                          count(u.id) over (partition by t.id) as user_count_in_team
          from teams t
                   left join users u on (SELECT 1 FROM json_each(u.teams) WHERE value = t.id))
        ,

     additional_team as
         (select distinct t.id,
                          count(u.id) over (partition by t.id) as user_count_in_additional_team
          from teams t
                   left join users u on (SELECT 1 FROM json_each(u.teams) WHERE value = t.id))

select event_type,
       event_title,
       event_date,
    e.id,
--e.additional_teams,
       team_name,
       club_name,
       t.user_count_in_team + ifnull(a.user_count_in_additional_team, 0)                      as user_in_team_s,
       case WHEN state = 'in' then state_count else 0 end                                     as inCount,
       case when state = 'out' then state_count else 0 end                                    as outCount,
       case when state = 'maybe' then state_count else 0 end                                  as maybeCount,
       t.user_count_in_team + ifnull(a.user_count_in_additional_team, 0) - (sum(state_count)) as ghostCount,
       (case WHEN state = 'in' then state_count else 0 end) /
       (t.user_count_in_team + ifnull(a.user_count_in_additional_team, 0)) * 100              as participation_rate


from event e
         left join team t on e.event_team_id = t.id
         left join additional_team a on (SELECT 1 FROM json_each(e.additional_teams) where value = a.id)
--left join stats p on e.id = p.event

--where e.id = 'xkwvz8ztynkv0a0'
where e.event_type = 'practice'
group by e.id
