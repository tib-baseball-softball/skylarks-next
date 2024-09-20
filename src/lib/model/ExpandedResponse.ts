import type { AuthModel } from "pocketbase"
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

export type CustomAuthModel = Extension<AuthModel, {
    first_name?: string,
    last_name?: string,
    email: string,
    last_login: string,
}>