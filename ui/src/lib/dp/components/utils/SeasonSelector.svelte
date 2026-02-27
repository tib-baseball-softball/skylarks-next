<script lang="ts">
  import {page} from "$app/state";

  interface Props {
    selectedSeason: number;
    seasonOptions: number[];
    onChangeCallback: () => void;
  }

  let {selectedSeason = $bindable(), seasonOptions, onChangeCallback}: Props = $props();

  const seasonParam = $derived(page.url.searchParams.get("season"));

  function parseParam() {
    if (seasonParam) {
      return parseInt(seasonParam);
    }
  }

  selectedSeason = parseParam() ?? 0;
</script>

<div class="container preset-tonal-surface rounded-base">
  <label class="season-label">
    Season
    <select bind:value={selectedSeason} class="select" onchange={onChangeCallback}>
      {#each seasonOptions as option}
        <option value="{option}">{option}</option>
      {/each}
    </select>
  </label>
</div>

<style>
  .container {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    gap: calc(var(--spacing) * 4);
    padding-inline: calc(var(--spacing) * 4);
    padding-block: calc(var(--spacing) * 3);

    @media (min-width: 64rem) {
      gap: calc(var(--spacing) * 8);
    }
  }

  .season-label {
    display: flex;
    align-items: center;
    flex-grow: 1;
    justify-content: space-between;
    gap: calc(var(--spacing) * 2);

    @media (min-width: 64rem) {
      flex-grow: 0;
    }
  }
</style>
