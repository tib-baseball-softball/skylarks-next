<script lang="ts">
    import {range} from "$lib/functions/range";
    import {selectedSeason} from "../../../stores";

    const seasonOptions = range(2021, new Date().getFullYear())
</script>

<select class="select" bind:value={$selectedSeason}>
    {#each seasonOptions as option}
        <option value="{option}">{option}</option>
    {/each}
</select>