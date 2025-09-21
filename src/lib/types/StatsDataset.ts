import {
  type BattingStatisticsEntry,
  type FieldingStatisticsEntry,
  type PitchingStatisticsEntry,
  StatsType
} from "bsm.js";
import type {BattingStatisticValues} from "bsm.js/dist/model/BattingStatisticsEntry";
import type {PitchingStatisticValues} from "bsm.js/dist/model/PitchingStatisticsEntry";
import type {FieldingStatisticValues} from "bsm.js/dist/model/FieldingStatisticsEntry";

export type StatsDataset = {
  batting: BattingStatisticsEntry,
  pitching: PitchingStatisticsEntry,
  fielding: FieldingStatisticsEntry
}

export type BattingRow = { values: BattingStatisticValues };
export type PitchingRow = { values: PitchingStatisticValues };
export type FieldingRow = { values: FieldingStatisticValues };

export type RowByType = {
  [StatsType.batting]: BattingRow;
  [StatsType.pitching]: PitchingRow;
  [StatsType.fielding]: FieldingRow;
};