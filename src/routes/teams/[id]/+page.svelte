<script lang="ts">
    import type { ClubTeam } from "bsm.js";
    import { ProgressRadial } from "@skeletonlabs/skeleton";
    import BaseballStatsDatatable from "$lib/components/datatable/BaseballStatsDatatable.svelte";
    import type { StatsDataset } from "$lib/types/StatsDataset";
    import TeamDetailInfoCard from "$lib/components/team/TeamDetailInfoCard.svelte";
    import StandingsTable from "$lib/components/table/StandingsTable.svelte";

    interface Props {
        data: any;
    }

    let { data }: Props = $props();
    let clubTeam = $derived(data.clubTeam as ClubTeam);

    async function getData(): Promise<StatsDataset> {
        return {
            batting: await data.battingStats!,
            pitching: await data.pitchingStats!,
            fielding: await data.fieldingStats!,
        };
    }
</script>

<h1 class="h1 my-4">{clubTeam.team.name} (Saison {clubTeam.team.season})</h1>

<section class="my-5">
    <h2 class="h2">Informationen</h2>
    <div
        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3 lg:gap-6 xl:gap-8"
    >
        <TeamDetailInfoCard {clubTeam} />

        {#await data?.table}
            <ProgressRadial />
        {:then table}
            {#if table}
                <div class="col-span-2">
                    <StandingsTable {table} />
                </div>
            {/if}
        {:catch error}
            <p>error loading: {error.message}</p>
        {/await}
    </div>
</section>

<section class="my-2">
    <h2 class="h2">Statistiken</h2>
    {#await getData()}
        <ProgressRadial />
    {:then stats}
        {#if stats.batting && stats.pitching && stats.fielding}
            <BaseballStatsDatatable data={stats} tableType="personal" />
        {/if}
    {:catch error}
        <p>error loading: {error.message}</p>
    {/await}
</section>

<style lang="postcss">
    h2 {
        @apply mb-3;
    }
</style>
