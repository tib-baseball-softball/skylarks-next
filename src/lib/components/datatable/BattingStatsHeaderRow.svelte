<script lang="ts">
    import type {DataHandler} from "@vincjo/datatables";
    import ThSort from "$lib/components/datatable/ThSort.svelte";

    export let handler: DataHandler<any>
</script>

<ThSort {handler} orderBy="values.games">G</ThSort>
<ThSort {handler} orderBy="games_started">GS</ThSort>
<ThSort {handler} orderBy="plate_appearances">PA</ThSort>
<ThSort {handler} orderBy="at_bats">AB</ThSort>
<ThSort {handler} orderBy="runs">R</ThSort>
<ThSort {handler} orderBy="runs_batted_in">RBI</ThSort>
<ThSort {handler} orderBy="hits">H</ThSort>
<ThSort {handler} orderBy="doubles">2B</ThSort>
<ThSort {handler} orderBy="triples">3B</ThSort>
<ThSort {handler} orderBy="homeruns">HR</ThSort>
<ThSort {handler} orderBy="strikeouts">K</ThSort>
<ThSort {handler} orderBy="base_on_balls">BB</ThSort>
<ThSort {handler} orderBy="intentional_base_on_balls">IBB</ThSort>
<ThSort {handler} orderBy="hit_by_pitches">HP</ThSort>
<ThSort {handler} orderBy="stolen_bases">SB</ThSort>
<ThSort {handler} orderBy="caught_stealings">CS</ThSort>
<ThSort {handler} orderBy="sacrifice_hits">SH</ThSort>
<ThSort {handler} orderBy="sacrifice_flys">SF</ThSort>
<ThSort {handler} orderBy="interferences">I</ThSort>
<ThSort {handler} orderBy="batting_average">AVG</ThSort>
<ThSort {handler} orderBy="on_base_percentage">OBP</ThSort>
<ThSort {handler} orderBy="slugging_percentage">SLG</ThSort>
<ThSort {handler} orderBy="on_base_plus_slugging">OPS</ThSort>