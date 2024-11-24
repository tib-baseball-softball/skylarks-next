import type {RecordModel} from "pocketbase"
import type {
  ClubsResponse,
  EventsResponse,
  ParticipationsResponse,
  TeamsResponse,
  UniformsetsResponse,
  UsersResponse
} from "./pb-types"

export type Extension<T, E> = T & E

export type ExpandedEvent = Extension<EventsResponse, {
  expand: {
    participations_via_event?: Extension<ParticipationsResponse, {
      expand: {
        user: UsersResponse
      }
    }>[],
    attire?: UniformsetsResponse,
    team?: ExpandedTeam
  },
  participations: {
    in: ExpandedParticipation[],
    out: ExpandedParticipation[],
    maybe: ExpandedParticipation[],
  },
  userParticipation: ExpandedParticipation
}>

export type ExpandedTeam = Extension<TeamsResponse, {
  expand: {
    club: ClubsResponse
  }
}>

export type ExpandedClub = Extension<ClubsResponse, {
  expand: {
    admins: UsersResponse[]
  }
}>

export type ExpandedUniformSet = Extension<UniformsetsResponse, {
  expand?: {
    club?: ClubsResponse
  }
}>

export type ExpandedParticipation = Extension<ParticipationsResponse, {
  expand?: {
    user?: CustomAuthModel
  }
}>

export type CustomAuthModel = Extension<RecordModel, {
  email?: string,
} & UsersResponse>

export type EventType = EventsResponse['type']
export type ParticipationType = ParticipationsResponse['state']
