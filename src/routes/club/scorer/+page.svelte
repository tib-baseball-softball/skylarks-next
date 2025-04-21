<script lang="ts">
  import type {PageProps} from './$types';
  import { Progress } from "@skeletonlabs/skeleton-svelte";
  import LicenseList from "$lib/components/club/licenses/LicenseList.svelte";

  let {data}: PageProps = $props();
</script>

<h1 class="h1">Scorer Licenses</h1>

{#await data.scorerLicenses}
  <p>Loading Licenses...</p>
  <Progress/>

{:then scorerLicenses}
  <LicenseList licenses={scorerLicenses} showSoftballSection={false}/>
{:catch error}
  <p>error loading: {error.message}</p>
{/await}