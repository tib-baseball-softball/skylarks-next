<script lang="ts">
  import {range} from "$lib/functions/range";
  import {preferences} from "$lib/stores";
  import type {AppPreferences} from "$lib/types/AppPreferences";

  const seasonOptions = range(2021, new Date().getFullYear())

  let selectedSeason: number = $state($preferences.selectedSeason)

  function updatePreferences(event: Event) {
    const target = event.target as HTMLSelectElement;
    preferences.update((current: AppPreferences) => ({
      ...current,
      selectedSeason: parseInt(target?.value, 10)
    }));
  }
</script>

<select class="select" bind:value={selectedSeason} onchange={updatePreferences}>
    {#each seasonOptions as option}
        <option value="{option}">{option}</option>
    {/each}
</select>