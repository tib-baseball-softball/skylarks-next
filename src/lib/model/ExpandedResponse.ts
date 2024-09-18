import type { ClubsResponse, EventsResponse, ParticipationsResponse, TeamsResponse, UniformsetsResponse, UsersResponse } from "./pb-types"

export type Extension<T, E> = T & E

export type ExpandedEvent = Extension<EventsResponse, {
    expand: {
        participations_via_event?: Extension<ParticipationsResponse, {
            expand:{
                user: UsersResponse
            } 
        }>[],
        attire?: UniformsetsResponse
    }
}>

export type ExpandedTeam = Extension<TeamsResponse, {
    expand: {
        club: ClubsResponse
    }
}>