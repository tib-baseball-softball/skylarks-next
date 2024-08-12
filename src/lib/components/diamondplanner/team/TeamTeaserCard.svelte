<script lang="ts">
    import {
        ClipboardCleanOutline,
        ClipboardOutline,
        FileLinesOutline, InfoCircleOutline
    } from "flowbite-svelte-icons";
    import type {TeamsResponse} from "$lib/model/pb-types";

    interface props {
        team: TeamsResponse
    }

    let {team}: props = $props()
</script>

<a href="/account/team/{team.id}" class="card card-hover block variant-ghost-surface p-3">

    <div class="flex items-center gap-3 self-end">
        <FileLinesOutline/>
        <div>
            <p>{team.name}</p>
            <p class="text-sm font-light">Name</p>
        </div>
    </div>

    <hr class="my-2">

    <div class="flex items-center gap-3">
        <InfoCircleOutline/>
        <div>
            <p>{team.bsm_league_group}</p>
            <p class="text-sm font-light">BSM-Liga (für aktuelle Saison)</p>
        </div>
    </div>
    
     <hr class="my-2">
    
    <div class="flex items-center gap-3">
        <ClipboardOutline/>
        <div>
            <p>{team?.expand?.club.name}</p>
            <p class="text-sm font-light">Club</p>
        </div>
    </div>
    
     <hr class="my-2">
    
    <div class="flex items-center gap-3">
        <ClipboardCleanOutline/>
        <div>
            <p>{team.age_group}</p>
            <p class="text-sm font-light">Altersgruppe</p>
        </div>
    </div>

</a>