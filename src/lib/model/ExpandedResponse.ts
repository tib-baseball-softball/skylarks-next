import type { ClubsResponse, EventsResponse, ParticipationsResponse, TeamsResponse, UsersResponse } from "./pb-types"

export type Extension<T, E> = T & E

export type ExpandedEvent = Extension<EventsResponse, {
    expand: {
        participations_via_event?: Extension<ParticipationsResponse, {
            expand:{
                user: UsersResponse
            } 
        }>[]
    }
}>

export type ExpandedTeam = Extension<TeamsResponse, {
    expand: {
        club: ClubsResponse
    }
}>