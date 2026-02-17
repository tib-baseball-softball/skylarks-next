<script lang="ts">
  // @ts-ignore
  import {Tabs} from "bits-ui";
  import PlayerRow from "$lib/dp/components/player/PlayerRow.svelte";
  import type {Player} from "$lib/dp/types/Player.ts";

  interface Props {
    players: Player[];
  }

  let {players}: Props = $props();
  let visual: "number" | "image" = $state("number");
</script>

<label class="label visual-selector-container">
  Visualization
  <Tabs.Root bind:value={visual}>
    <Tabs.List class="tabs-list selector-tabs">
      <Tabs.Trigger class="tabs-trigger btn selector-trigger" value="number">Number</Tabs.Trigger>
      <Tabs.Trigger class="tabs-trigger btn selector-trigger" value="image">Image</Tabs.Trigger>
    </Tabs.List>
  </Tabs.Root>
</label>

<div class="card players-list-card preset-tonal shadow-xl">
  {#each players as player (player.uid)}
    <PlayerRow {player} {visual}/>

    <hr class="row-divider">
  {/each}
</div>

<style>
    .visual-selector-container {
        display: flex;
        flex-direction: column;
        gap: calc(var(--spacing) * 1);
        margin-block: calc(var(--spacing) * 4);
        
        @media (min-width: 64rem) {
            max-width: 50%;
        }
    }

    .selector-tabs {
        display: flex;
        background-color: var(--color-surface-50-950);
        border: 1px solid var(--color-surface-200);
        border-radius: var(--radius-base);
        padding: calc(var(--spacing) * 1);
        
        :global([data-theme='dark']) & {
            border-color: var(--color-surface-700);
        }
    }

    .selector-trigger {
        flex-grow: 1;
    }

    :global(.selector-trigger[data-state='active']) {
        background-color: var(--color-surface-100-900);
        
        :global([data-theme='dark']) & {
            background-color: var(--color-surface-800);
        }
    }

    .players-list-card {
        padding: calc(var(--spacing) * 3);
        
        :global([data-theme='dark']) & {
            border: 1px solid var(--color-surface-500);
        }
    }

    .row-divider {
        margin-block: calc(var(--spacing) * 3);
    }
</style>
