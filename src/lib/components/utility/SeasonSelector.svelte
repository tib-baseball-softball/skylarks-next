<script lang="ts">
    import {range} from "$lib/functions/range";
    import {preferences} from "$lib/stores";
    import type {AppPreferences} from "$lib/types/AppPreferences";

    const seasonOptions = range(2021, new Date().getFullYear())

    let selectedSeason: number

    $: {
        preferences.subscribe((value: AppPreferences) => {
            if (value.selectedSeason !== selectedSeason) {
                selectedSeason = value.selectedSeason;
            }
        });
    }

    function updatePreferences(event: Event) {
        const target = event.target as HTMLSelectElement;
        preferences.update((current: AppPreferences) => ({
            ...current,
            selectedSeason: parseInt(target?.value, 10)
        }));
    }
</script>

<select class="select" bind:value={selectedSeason} on:change={updatePreferences}>
    {#each seasonOptions as option}
        <option value="{option}">{option}</option>
    {/each}
</select>