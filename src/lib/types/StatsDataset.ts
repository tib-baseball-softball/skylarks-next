import type {BattingStatisticsEntry, FieldingStatisticsEntry, PitchingStatisticsEntry} from "bsm.js";

export type StatsDataset = {
    batting: BattingStatisticsEntry,
    pitching: PitchingStatisticsEntry,
    fielding: FieldingStatisticsEntry
}