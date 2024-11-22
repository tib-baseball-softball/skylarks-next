import type {EventType, ParticipationType} from "$lib/model/ExpandedResponse";

export interface PersonalAttendanceStatsItem {
  season: number
  attendanceTotals: AttendanceTotal[]
  participationTotals: ParticipationTotal[]
  totalPossibleEvents: number
}

export type ParticipationByPerson = {
  id: string;
  lastName: string;
  firstName: string;
  type: Exclude<EventType, "">;
  inCount: number;
  outCount: number;
  maybeCount: number;
  totalCount: number;
}

export interface AttendanceTotal {
  type: EventType
  attended: number
  total: number
}

export interface ParticipationTotal {
  type: Exclude<ParticipationType, "">
  amount: number
}