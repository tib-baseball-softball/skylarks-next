<script lang="ts">
  import ClubDetailCard from "$lib/components/diamondplanner/club/ClubDetailCard.svelte";
  import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";

  let {data} = $props()

  let club = $derived(data.club)
  let teams = $derived(data.teams)
</script>

<h1 class="h1">{club.name}</h1>

<div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    <ClubDetailCard {club}/>
</div>

<h2 class="h2">Club Teams</h2>

<div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    {#each teams as team}
        <a href="/account/team/{team.id}">
            <TeamTeaserCard {team} link={true}/>
        </a>
    {/each}
</div>