<script lang="ts">
    import type {LayoutServerData} from "../../../.svelte-kit/types/src/routes/gamecenter/$types";
    import SeasonSelector from "$lib/components/utility/SeasonSelector.svelte";
    import MatchTeaserCard from "$lib/components/match/MatchTeaserCard.svelte";
    import {ProgressBar} from "@skeletonlabs/skeleton";

    export let data: LayoutServerData
</script>

<SeasonSelector/>
{#await data.streamed.matches}
    <p>Loading matches...</p>
    <ProgressBar />
{:then matches}
    <section>
        <header>
            <h2 class="h2">Spiele</h2>
        </header>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 lg:gap-6">
            {#each matches as match}
                <MatchTeaserCard {match}/>
            {/each}
        </div>
    </section>

{:catch error}
    <p>error loading matches: {error.message}</p>
{/await}
