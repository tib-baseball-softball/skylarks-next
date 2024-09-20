<script lang="ts">
  import Pants from "$lib/components/icons/Pants.svelte";
  import Shirt from "$lib/components/icons/Shirt.svelte";
  import type { UniformsetsResponse } from "$lib/model/pb-types";

  interface props {
    attire: UniformsetsResponse;
  }

  let { attire }: props = $props();
</script>

<h3 class="h3">Attire</h3>
<div class="p-4 flex gap-2 md:gap-3 justify-evenly">
    <!-- We are cheating here, this does not work for other clubs -->

  {#if attire.cap === "B"}
    <img src="/uniform/B_cap.svg" alt="Berlin Skylarks Home Cap">
    {:else if attire.cap === "Bird"}
    <img src="/uniform/Bird_cap.svg" alt="Berlin Skylarks Road Cap">
  {/if}
  <Shirt fillColor={attire.jersey} />
  <Pants fillColor={attire.pants} />
</div>

<style lang="postcss">
    img {
        @apply w-16
    }
</style>