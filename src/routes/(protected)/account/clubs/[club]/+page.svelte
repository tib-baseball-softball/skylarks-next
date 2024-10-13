<script lang="ts">
  import ClubDetailCard from "$lib/components/diamondplanner/club/ClubDetailCard.svelte";
  import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";
  import UniformSetInfoCard from "$lib/components/diamondplanner/uniformset/UniformSetInfoCard.svelte";
  import {PlusOutline} from "flowbite-svelte-icons";

  let {data} = $props()

  let club = $derived(data.club)
  let teams = $derived(data.teams)
  let uniformSets = $derived(data.uniformSets)
</script>

<h1 class="h1">{club.name}</h1>

<section class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    <ClubDetailCard {club}/>
</section>

<h2 class="h2">Club Teams</h2>

<section class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    {#each teams as team}
        <a href="/account/team/{team.id}">
            <TeamTeaserCard {team} link={true}/>
        </a>
    {/each}
</section>

<h2 class="h2">Uniform Sets</h2>
<section class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    {#each uniformSets as uniformSet}
        <UniformSetInfoCard {uniformSet}/>
    {/each}
</section>
<button class="btn variant-ghost-primary">
    <PlusOutline/>
    <span>Create new</span>
</button>
