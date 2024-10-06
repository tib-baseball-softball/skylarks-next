<script lang="ts">
    import type { AdditionalMatchStats } from "bsm.js";
    import MatchBoxscoreAdditionalStatsSectionEntry from "$lib/components/boxscore/MatchBoxscoreAdditionalStatsSectionEntry.svelte";

    interface Props {
        stats: AdditionalMatchStats;
    }

    let { stats }: Props = $props();
</script>

<section class="my-3">
    <h3 class="font-bold">BATTING</h3>

    {#if stats.batting.doubles.length > 0}
        <span>
            <span class="font-bold">2B:</span>
            <MatchBoxscoreAdditionalStatsSectionEntry
                entries={stats.batting.doubles}
            />
        </span>
    {/if}

    {#if stats.batting.triples.length > 0}
        <div>
            <span class="font-bold">3B:</span>
            <MatchBoxscoreAdditionalStatsSectionEntry
                entries={stats.batting.triples}
            />
        </div>
    {/if}

    {#if stats.batting.homeruns.length > 0}
        <div>
            <span class="font-bold">HR:</span>
            <MatchBoxscoreAdditionalStatsSectionEntry
                entries={stats.batting.homeruns}
            />
        </div>
    {/if}

    {#if stats.batting.sacrifice_hits.length > 0}
        <div>
            <span class="font-bold">SH:</span>
            <MatchBoxscoreAdditionalStatsSectionEntry
                entries={stats.batting.sacrifice_hits}
            />
        </div>
    {/if}

    {#if stats.batting.sacrifice_flys.length > 0}
        <div>
            <span class="font-bold">SF:</span>
            <MatchBoxscoreAdditionalStatsSectionEntry
                entries={stats.batting.sacrifice_flys}
            />
        </div>
    {/if}
</section>

<section class="my-3">
    <h3 class="font-bold">BASERUNNING</h3>

    {#if stats.baserunning.stolen_bases.length > 0}
        <div>
            <span class="font-bold">SB:</span>
            <MatchBoxscoreAdditionalStatsSectionEntry
                entries={stats.baserunning.stolen_bases}
            />
        </div>
    {/if}

    {#if stats.baserunning.caught_stealings.length > 0}
        <div>
            <span class="font-bold">CS:</span>
            <MatchBoxscoreAdditionalStatsSectionEntry
                entries={stats.baserunning.caught_stealings}
            />
        </div>
    {/if}
</section>

<section class="my-3">
    <h3 class="font-bold">FIELDING</h3>

    {#if stats.fielding.errors.length > 0}
        <div>
            <span class="font-bold">E:</span>
            <MatchBoxscoreAdditionalStatsSectionEntry
                entries={stats.fielding.errors}
            />
        </div>
    {/if}

    {#if stats.fielding.passed_balls.length > 0}
        <div>
            <span class="font-bold">PB:</span>
            <MatchBoxscoreAdditionalStatsSectionEntry
                entries={stats.fielding.passed_balls}
            />
        </div>
    {/if}

    {#if stats.fielding.double_plays.length > 0}
        <div>
            <span class="font-bold">DP:</span>
            <MatchBoxscoreAdditionalStatsSectionEntry
                entries={stats.fielding.double_plays}
            />
        </div>
    {/if}

    {#if stats.fielding.triple_plays.length > 0}
        <div>
            <span class="font-bold">TP:</span>
            <MatchBoxscoreAdditionalStatsSectionEntry
                entries={stats.fielding.triple_plays}
            />
        </div>
    {/if}
</section>
