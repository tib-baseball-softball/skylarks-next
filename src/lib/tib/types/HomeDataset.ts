import type {LeagueGroup, Match, Table} from "bsm.js";
//@ts-expect-error
import type {TableRow} from "bsm.js/dist/model/Table";
import type {PlayoffSeries} from "$lib/tib/types/PlayoffSeries.ts";

export type HomeDataset = {
  team_id: number
  season: number
  league_group_id: number
  league_group: LeagueGroup
  next_game: Match | null
  last_game: Match | null
  playoff_series: PlayoffSeries | null
  table: Table
  table_row: TableRow | null
  streak_data: StreakDataEntry[] | null
}

export type StreakDataEntry = {
  game: string
  won_game: boolean
  wins_count: number
}
