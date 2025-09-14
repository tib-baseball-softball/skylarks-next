<script lang="ts">
  import type {Player} from "$lib/model/Player.ts";
  import PlayerRow from "$lib/components/player/PlayerRow.svelte";
  import {Segment} from "@skeletonlabs/skeleton-svelte";

  interface Props {
    players: Player[];
  }

  let {players}: Props = $props();
  let visual: "number" | "image" = $state("number");

  // @ts-expect-error
  function changeValue(details) {
    visual = details.value ?? "number";
  }
</script>

<label class="label flex flex-col gap-1 my-4 lg:max-w-[50%]">
  Visualization
  <Segment name="visual" onValueChange={changeValue} value={visual}>
    <Segment.Item classes="flex-grow" value={"number"}>
      Number
    </Segment.Item>
    <Segment.Item classes="flex-grow" value={"image"}>
      Image
    </Segment.Item>
  </Segment>
</label>

<div class="card p-3 preset-tonal dark:border dark:border-surface-500 shadow-xl">
  {#each players as player (player.uid)}
    <PlayerRow {player} {visual}/>

    <hr class="my-3">
  {/each}
</div>
