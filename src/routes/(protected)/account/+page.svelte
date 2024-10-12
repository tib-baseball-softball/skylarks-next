<script lang="ts">
    import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";
    import UserDataCard from "$lib/components/diamondplanner/user/UserDataCard.svelte";
    import type {CustomAuthModel} from "$lib/model/ExpandedResponse.js";
    import {authModel} from "$lib/pocketbase/Auth";
    import type {ClubsResponse} from "$lib/model/pb-types";
    import PlayerProfileClubsSection from "$lib/components/diamondplanner/user/PlayerProfileClubsSection.svelte";

    const model = $authModel as CustomAuthModel;

    let {data} = $props();
    const clubs: ClubsResponse[] = model?.expand?.club
</script>

<h1 class="h1 lg:mt-4">My Dashboard</h1>

<div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    <UserDataCard {model}/>

    <PlayerProfileClubsSection {clubs}/>
</div>

<h2 class="h2 mt-3">My Teams</h2>

{#await data.teams then teams}
    <div class="grid grid-cols-1 md:grid-cols-2 gap-3 mb-3">
        {#each teams as team}
            <a href="/account/team/{team.id}">
                <TeamTeaserCard {team} link={true}/>
            </a>
        {/each}

        {#if teams?.length === 0}
            <p>You are not a member of any teams yet.</p>
        {/if}
    </div>
{/await}
