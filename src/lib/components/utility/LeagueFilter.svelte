<script lang="ts">
    import { run } from 'svelte/legacy';

    import {preferences} from "$lib/stores";
    import type {AppPreferences} from "$lib/types/AppPreferences";
    import type {LeagueGroup} from "bsm.js";

    interface Props {
        leagueGroups?: LeagueGroup[];
    }

    let { leagueGroups = [] }: Props = $props();

    let selectedID = $state(0)

    run(() => {
        preferences.subscribe((value: AppPreferences) => {
            if (value.leagueGroupID !== selectedID) {
                selectedID = value.leagueGroupID;
            }
        });
    });

    function updatePreferences(event: Event) {
        //const target = event.target as HTMLSelectElement;
        preferences.update((current: AppPreferences) => ({
            ...current,
            leagueGroupID: selectedID
        }));
    }
</script>

<select class="select min-w-52" bind:value={selectedID} onchange={updatePreferences}>
    <option value="{0}" selected>Alle Ligen</option>

    {#each leagueGroups as leagueGroup}
        <option value="{leagueGroup.id}">{leagueGroup.name}</option>
    {/each}
</select>