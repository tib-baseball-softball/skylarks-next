<script lang="ts">
    import {range} from "$lib/functions/range";
    import {preferences} from "$lib/stores";

    const seasonOptions = range(2021, new Date().getFullYear());
</script>

<select class="select" bind:value={$preferences.selectedSeason}>
  {#each seasonOptions as option}
    <option value="{option}">{option}</option>
  {/each}
</select>