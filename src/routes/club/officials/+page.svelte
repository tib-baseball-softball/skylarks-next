<script lang="ts">
  import type {PageProps} from './$types';
  import ClubOfficialsList from "$lib/components/club/officials/ClubOfficialsList.svelte";
  import {ProgressBar} from "@skeletonlabs/skeleton";

  let {data}: PageProps = $props();
</script>

<h1 class="h1">Club Officials</h1>


{#await data.clubOfficials}
  <p>Loading Club Officials...</p>
  <ProgressBar/>

{:then clubOfficials}
  <ClubOfficialsList clubOfficials={clubOfficials}/>
{:catch error}
  <p>error loading: {error.message}</p>
{/await}