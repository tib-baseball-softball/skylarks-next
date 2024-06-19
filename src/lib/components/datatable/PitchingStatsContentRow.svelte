<script lang="ts">
    import type {StatisticsData} from "bsm.js/dist/model/AbstractStatisticsEntry";

    export let row: StatisticsData<"PitchingStatistics">
</script>

<td>{row.values.games}</td>
<td>{row.values.games_started}</td>
<td>{row.values.complete_games}</td>
<td>{row.values.batters_faced}</td>
<td>{row.values.at_bats}</td>
<td>{row.values.runs}</td>
<td>{row.values.earned_runs}</td>
<td>{row.values.hits}</td>
<td>{row.values.doubles}</td>
<td>{row.values.triples}</td>
<td>{row.values.homeruns}</td>
<td>{row.values.strikeouts}</td>
<td>{row.values.balks}</td>
<td>{row.values.wild_pitches}</td>
<td>{row.values.base_on_balls_allowed}</td>
<td>{row.values.intentional_base_on_balls}</td>
<td>{row.values.hit_by_pitches}</td>
<td>{row.values.interferences}</td>
<td>{row.values.earned_runs_average}</td>
<td>{row.values.innings_pitched}</td>
<td>{row.values.walks_and_hits_per_innings_pitched}</td>
<td>{row.values.wins}</td>
<td>{row.values.losses}</td>
<td>{row.values.saves}</td>