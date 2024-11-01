<script lang="ts">
    import UserDataCard from "$lib/components/diamondplanner/user/UserDataCard.svelte";
    import type {CustomAuthModel} from "$lib/model/ExpandedResponse.js";
    import {authModel} from "$lib/pocketbase/Auth";
    import PlayerProfileClubsSection from "$lib/components/diamondplanner/user/PlayerProfileClubsSection.svelte";
    import TeamListTeaser from "$lib/components/diamondplanner/team/TeamListTeaser.svelte";

    const model = $authModel as CustomAuthModel;

    let {data} = $props();
</script>

<h1 class="h1 lg:mt-4">My Dashboard</h1>

<div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    <UserDataCard {model}/>

    {#await data.clubs then clubs}
        <PlayerProfileClubsSection {clubs}/>
    {/await}
</div>

<h2 class="h2 mt-3">My Teams</h2>

{#await data.teams then teams}
    <div class="grid grid-cols-1 md:grid-cols-2 gap-3 mb-3">
        {#each teams as team}
            <TeamListTeaser {team} link={true}/>
        {/each}

        {#if teams?.length === 0}
            <p>You are not a member of any teams yet.</p>
        {/if}
    </div>
{/await}
