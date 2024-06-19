<script lang="ts">
    import type {StatisticsData} from "bsm.js/dist/model/AbstractStatisticsEntry";

    export let row: StatisticsData<"BattingStatistics">
</script>

<td>{row.values.games}</td>
<td>{row.values.games_started}</td>
<td>{row.values.plate_appearances}</td>
<td>{row.values.at_bats}</td>
<td>{row.values.runs}</td>
<td>{row.values.runs_batted_in}</td>
<td>{row.values.hits}</td>
<td>{row.values.doubles}</td>
<td>{row.values.triples}</td>
<td>{row.values.homeruns}</td>
<td>{row.values.strikeouts}</td>
<td>{row.values.base_on_balls}</td>
<td>{row.values.intentional_base_on_balls}</td>
<td>{row.values.hit_by_pitches}</td>
<td>{row.values.stolen_bases}</td>
<td>{row.values.caught_stealings}</td>
<td>{row.values.sacrifice_hits}</td>
<td>{row.values.sacrifice_flys}</td>
<td>{row.values.interferences}</td>
<td>{row.values.batting_average}</td>
<td>{row.values.on_base_percentage}</td>
<td>{row.values.slugging_percentage}</td>
<td>{row.values.on_base_plus_slugging}</td>