<script lang="ts">
    import {StatsType} from "bsm.js";
    import type {StatisticsData} from "bsm.js";
    import StatsBlockContent from "$lib/components/utility/StatsBlockContent.svelte";

    interface props {
        type: StatsType,
        row: StatisticsData<"BattingStatistics" | "PitchingStatistics" | "FieldingStatistics">
    }
    let {type, row}: props = $props()
</script>

{#if row}
    {#if type === StatsType.batting}
        <StatsBlockContent block={{ title: 'AVG', value: row.values.batting_average, desc: 'Batting Average' }}/>
        <StatsBlockContent block={{ title: 'OBP', value: row.values.on_base_percentage, desc: 'On Base Percentage' }}/>
        <StatsBlockContent block={{ title: 'SLG', value: row.values.slugging_percentage, desc: 'Slugging Average' }}/>
        <StatsBlockContent block={{ title: 'OPS', value: row.values.on_base_plus_slugging, desc: 'On Base Plus Slugging' }}/>
    {/if}

    {#if type === StatsType.pitching}
        <StatsBlockContent block={{ title: 'W - L', value: `${row.values.wins} - ${row.values.losses}`, desc: 'Wins - Losses' }}/>
        <StatsBlockContent block={{ title: 'IP', value: row.values.innings_pitched, desc: 'Innings Pitched' }}/>
        <StatsBlockContent block={{ title: 'ERA', value: row.values.earned_runs_average, desc: 'Earned Run Average' }}/>
        <StatsBlockContent block={{ title: 'WHIP', value: row.values.walks_and_hits_per_innings_pitched, desc: 'Walks & Hits per Innings Pitched' }}/>
    {/if}

    {#if type === StatsType.fielding}
        <StatsBlockContent block={{ title: 'G', value: row.values.games, desc: 'Games' }}/>
        <StatsBlockContent block={{ title: 'IP', value: row.values.innings_played, desc: 'Innings Played' }}/>
        <StatsBlockContent block={{ title: 'AVG', value: row.values.fielding_average, desc: 'Fielding Average' }}/>
    {/if}
{/if}