<script lang="ts">
    import SeasonSelector from "$lib/components/utility/SeasonSelector.svelte";
    import MatchTeaserCard from "$lib/components/match/MatchTeaserCard.svelte";
    import {ProgressBar, Tab, TabGroup} from "@skeletonlabs/skeleton";
    import {Gameday} from "bsm.js";
    import {preferences} from "$lib/stores";
    import LeagueFilter from "$lib/components/utility/LeagueFilter.svelte";
    import {browser} from "$app/environment";

    async function reloadGameData(gameday: Gameday, season: number) {
        if (!browser) {
            return
        }

        const basePath = "/api/matches"
        const params = [
           [ "season", season.toString() ],
           [ "gameday", gameday]
        ]

        const urlParams = new URLSearchParams(params)
        const urlWithParams = `${basePath}?${urlParams.toString()}`

        data.streamed.matches = (await fetch(urlWithParams)).json()
    }

    export let data
    $: leagueGroups = data.leagueGroups

    $: reloadGameData($preferences.gameday, $preferences.selectedSeason)
</script>


<div class="my-2 md:flex justify-between items-center">
    <h1 class="h1">Gamecenter</h1>

    <div class="flex gap-2">
        {#await leagueGroups then groups}
            <LeagueFilter leagueGroups="{groups}"/>
        {/await}
        <SeasonSelector/>
    </div>
</div>

<section class="mb-5 mt-3">
    <label id="gameday_label">Spieltag</label>
    <TabGroup justify="justify-center" labelledby="gameday_label">
        <Tab bind:group={$preferences.gameday} name="tabPrevious" value={Gameday.previous}>Voriger</Tab>
        <Tab bind:group={$preferences.gameday} name="tabCurrent" value={Gameday.current}>Aktuell</Tab>
        <Tab bind:group={$preferences.gameday} name="tabNext" value={Gameday.next}>Nächster</Tab>
        <Tab bind:group={$preferences.gameday} name="tabAny" value={Gameday.any}>Alle</Tab>
        <!-- Tab Panels --->
        <svelte:fragment slot="panel">


        </svelte:fragment>
    </TabGroup>
</section>

{#await data.streamed.matches}
    <p>Loading matches...</p>
    <ProgressBar/>
{:then matches}
    <section>
        <header>
            <h2 class="h2 my-3">Spiele</h2>
        </header>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 lg:gap-6 mb-6">
            {#each matches as match}
                <MatchTeaserCard {match}/>
            {/each}
        </div>
    </section>

{:catch error}
    <p>error loading matches: {error.message}</p>
{/await}
