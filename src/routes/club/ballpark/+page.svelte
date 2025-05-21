<script lang="ts">
  import type {PageProps} from './$types';
  import { Progress } from "@skeletonlabs/skeleton-svelte";
  import BallparkGrid from "$lib/components/club/ballpark/BallparkGrid.svelte";

  let {data}: PageProps = $props();
</script>

<h1 class="h1">Ballpark</h1>


{#await data.fields}
  <p>Loading Club Officials...</p>
  <Progress/>

{:then fields}
  <BallparkGrid fields={fields}/>
{:catch error}
  <p>error loading: {error.message}</p>
{/await}