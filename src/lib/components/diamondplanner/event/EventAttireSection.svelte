<script lang="ts">
  import Pants from "$lib/components/icons/Pants.svelte";
  import Shirt from "$lib/components/icons/Shirt.svelte";
  import type {UniformsetsResponse} from "$lib/model/pb-types";
  import Cap from "$lib/components/icons/Cap.svelte";

  interface Props {
    attire: UniformsetsResponse;
  }

  let {attire}: Props = $props();
</script>

<h3 class="h3">Attire</h3>
<div class="p-4 flex gap-2 md:gap-3 justify-evenly">
    <Cap identifier={attire.cap} classes="w-16"/>
    <Shirt fillColor={attire.jersey} classes="w-12 h-12"/>
    <Pants fillColor={attire.pants} classes="w-12 h-12"/>
</div>