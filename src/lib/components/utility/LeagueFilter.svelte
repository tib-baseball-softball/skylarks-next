<script lang="ts">
    import {preferences} from "$lib/stores";
    import type {AppPreferences} from "$lib/types/AppPreferences";
    import type {LeagueGroup} from "bsm.js";

    export let leagueGroups: LeagueGroup[] = []

    let selectedID = 0

    $: {
        preferences.subscribe((value: AppPreferences) => {
            if (value.leagueGroupID !== selectedID) {
                selectedID = value.leagueGroupID;
            }
        });
    }

    function updatePreferences(event: Event) {
        const target = event.target as HTMLSelectElement;
        preferences.update((current: AppPreferences) => ({
            ...current,
            leagueGroupID: parseInt(target?.value, 10)
        }));
    }
</script>

<select class="select min-w-52" bind:value={selectedID} on:change={updatePreferences}>
    <option value="{0}" selected>Alle Ligen</option>

    {#each leagueGroups as leagueGroup}
        <option value="{leagueGroup.id}">{leagueGroup.name}</option>
    {/each}
</select>