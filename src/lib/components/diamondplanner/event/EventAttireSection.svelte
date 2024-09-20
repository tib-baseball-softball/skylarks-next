<script lang="ts">
  import Baseball from "$lib/components/icons/Baseball.svelte";
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
  <Baseball strokeColor={attire.cap} />
  <Shirt fillColor={attire.jersey} />
  <Pants fillColor={attire.pants} />
</div>
