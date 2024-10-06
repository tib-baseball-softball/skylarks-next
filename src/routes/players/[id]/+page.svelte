<script lang="ts">
    import type { PageServerData } from "../../../../.svelte-kit/types/src/routes/players/[id]/$types";
    import { ProgressBar } from "@skeletonlabs/skeleton";
    import BaseballStatsDatatable from "$lib/components/datatable/BaseballStatsDatatable.svelte";
    import type { StatsDataset } from "$lib/types/StatsDataset";
    import PlayerHeaderSection from "$lib/components/player/PlayerHeaderSection.svelte";
    import type { Player } from "$lib/model/Player";
    import PlayerDataCard from "$lib/components/player/PlayerDataCard.svelte";

    interface Props {
        data: PageServerData;
    }

    let { data }: Props = $props();

    async function getData(): Promise<StatsDataset> {
        return {
            batting: await data.battingStats!,
            pitching: await data.pitchingStats!,
            fielding: await data.fieldingStats!,
        };
    }

    async function getPlayer(): Promise<Player | undefined> {
        const players = await data.player;

        if (players.length === 1) {
            return players.at(0);
        }
    }
</script>

{#await data.battingStats then batting}
    <h1 class="h1 my-4">
        Spieler*innenprofil für {batting.person?.first_name}
        {batting.person?.last_name}
    </h1>
{/await}

{#await getPlayer()}
    <ProgressBar />
{:then player}
    {#if player}
        <h2 class="h2 my-4">Profildaten</h2>
        <section class="grid grid-cols-1 md:grid-cols-2 gap-3 md:gap-4">
            <PlayerHeaderSection {player} />
            <PlayerDataCard {player} />
        </section>
    {/if}
{:catch error}
    <p>error loading: {error.message}</p>
{/await}

<h2 class="h2 my-6 lg:mt-10 mb-8">Statistiken</h2>

{#await getData()}
    <ProgressBar />
{:then stats}
    <BaseballStatsDatatable data={stats} tableType="seasonal" />
{:catch error}
    <p>error loading: {error.message}</p>
{/await}
