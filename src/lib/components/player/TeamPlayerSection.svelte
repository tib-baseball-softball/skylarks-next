<script lang="ts">
import type { Player } from "$lib/model/Player.ts"
import PlayerRow from "$lib/components/player/PlayerRow.svelte"
// @ts-ignore
import { Tabs } from "bits-ui"

interface Props {
  players: Player[]
}

let { players }: Props = $props()
let visual: "number" | "image" = $state("number")
</script>

<label class="label flex flex-col gap-1 my-4 lg:max-w-[50%]">
  Visualization
  <Tabs.Root bind:value={visual}>
    <Tabs.List class="tabs-list input">
      <Tabs.Trigger class="tabs-trigger btn flex-grow" value="number">Number</Tabs.Trigger>
      <Tabs.Trigger class="tabs-trigger btn flex-grow" value="image">Image</Tabs.Trigger>
    </Tabs.List>
  </Tabs.Root>
</label>

<div class="card p-3 preset-tonal dark:border dark:border-surface-500 shadow-xl">
  {#each players as player (player.uid)}
    <PlayerRow {player} {visual}/>

    <hr class="my-3">
  {/each}
</div>
