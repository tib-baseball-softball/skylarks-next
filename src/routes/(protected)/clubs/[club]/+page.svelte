<script lang="ts">
  import ClubDetailCard from "$lib/components/diamondplanner/club/ClubDetailCard.svelte";

  let {data} = $props()

  let club = $derived(data.club)
</script>

<h1 class="h1">{club.name}</h1>

<div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    <ClubDetailCard {club}/>
</div>