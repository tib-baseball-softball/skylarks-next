<script lang="ts">
  import type {PageProps} from './$types';
  import { Progress } from "@skeletonlabs/skeleton-svelte";
  import LicenseList from "$lib/components/club/licenses/LicenseList.svelte";

  let {data}: PageProps = $props();
</script>

<h1 class="h1">Umpire Licenses</h1>

{#await data.umpireLicenses}
  <p>Loading Licenses...</p>
  <Progress/>

{:then umpireLicenses}
  <LicenseList licenses={umpireLicenses} showSoftballSection={true}/>
{:catch error}
  <p>error loading: {error.message}</p>
{/await}