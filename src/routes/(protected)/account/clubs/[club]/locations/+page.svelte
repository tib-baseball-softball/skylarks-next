<script lang="ts">
  import type {PageProps} from "./$types";
  import type {ExpandedClub} from "$lib/model/ExpandedResponse.ts";
  import ClubLocationsSection from "$lib/components/diamondplanner/club/ClubLocationsSection.svelte";

  let { data }: PageProps = $props();

  let club: ExpandedClub = $derived(data.club);
</script>

<svelte:head>
  <title>Locations for {club.name}</title>
  <meta
          name="description"
          content="All configured locations for {club.name}, including baseball fields and other venues."
  />
</svelte:head>

<h1 class="h1">{club.name} - Locations</h1>

<section class="mt-8!">
  <ClubLocationsSection club={club} locations={data?.locations ?? []} />
</section>