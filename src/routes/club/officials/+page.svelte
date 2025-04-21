<script lang="ts">
  import type {PageProps} from './$types';
  import ClubOfficialsList from "$lib/components/club/officials/ClubOfficialsList.svelte";
  import { Progress } from "@skeletonlabs/skeleton-svelte";

  let {data}: PageProps = $props();
</script>

<h1 class="h1">Club Officials</h1>


{#await data.clubOfficials}
  <p>Loading Club Officials...</p>
  <Progress/>

{:then clubOfficials}
  <ClubOfficialsList clubOfficials={clubOfficials}/>
{:catch error}
  <p>error loading: {error.message}</p>
{/await}