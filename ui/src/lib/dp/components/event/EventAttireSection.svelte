<script lang="ts">
  import Cap from "$lib/dp/components/icons/Cap.svelte";
  import Pants from "$lib/dp/components/icons/Pants.svelte";
  import Shirt from "$lib/dp/components/icons/Shirt.svelte";
  import type {UniformsetsResponse} from "$lib/dp/types/pb-types.ts";

  interface Props {
    attire: UniformsetsResponse;
  }

  let {attire}: Props = $props();
</script>

<h2 class="h2">Uniform</h2>
<div class="attire-container">
  <Cap identifier={attire.cap} size={64}/>
  <Shirt fillColor={attire.jersey} size={48}/>
  <Pants fillColor={attire.pants} size={48}/>
</div>

<style>
  .attire-container {
    padding: calc(var(--spacing) * 4);
    display: flex;
    gap: calc(var(--spacing) * 2);
    justify-content: space-evenly;

    @media (min-width: 48rem) {
      gap: calc(var(--spacing) * 3);
    }
  }
</style>
