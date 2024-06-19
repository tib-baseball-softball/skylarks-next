<script lang="ts">
    import ThSort from "$lib/components/datatable/ThSort.svelte";
    import type {DataHandler} from "@vincjo/datatables";

    export let handler: DataHandler<any>
</script>

<ThSort {handler} orderBy="games">G</ThSort>
<ThSort {handler} orderBy="games_started">GS</ThSort>
<ThSort {handler} orderBy="complete_games">CG</ThSort>
<ThSort {handler} orderBy="batters_faced">BF</ThSort>
<ThSort {handler} orderBy="at_bats">AB</ThSort>
<ThSort {handler} orderBy="runs">R</ThSort>
<ThSort {handler} orderBy="earned_runs">ER</ThSort>
<ThSort {handler} orderBy="hits">H</ThSort>
<ThSort {handler} orderBy="doubles">2B</ThSort>
<ThSort {handler} orderBy="triples">3B</ThSort>
<ThSort {handler} orderBy="homeruns">HR</ThSort>
<ThSort {handler} orderBy="strikeouts">K</ThSort>
<ThSort {handler} orderBy="balks">BK</ThSort>
<ThSort {handler} orderBy="wild_pitches">WP</ThSort>
<ThSort {handler} orderBy="base_on_balls_allowed">BB</ThSort>
<ThSort {handler} orderBy="intentional_base_on_balls">IBB</ThSort>
<ThSort {handler} orderBy="hit_by_pitches">HBP</ThSort>
<ThSort {handler} orderBy="interferences">I</ThSort>
<ThSort {handler} orderBy="earned_runs_average">ERA</ThSort>
<ThSort {handler} orderBy="innings_pitched">IP</ThSort>
<ThSort {handler} orderBy="walks_and_hits_per_innings_pitched">WHIP</ThSort>
<ThSort {handler} orderBy="wins">W</ThSort>
<ThSort {handler} orderBy="losses">L</ThSort>
<ThSort {handler} orderBy="saves">SV</ThSort>
