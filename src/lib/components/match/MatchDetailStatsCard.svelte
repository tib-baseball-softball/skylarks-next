<script lang="ts">
    import type {Match} from "bsm.js";
    import {
        HourglassOutline,
        InfoCircleOutline,
        LinkOutline,
        TableRowOutline
    } from "flowbite-svelte-icons";

    export let match: Match
</script>

<div class="card variant-ghost-surface p-3">
    <div class="flex items-center gap-3">
        <TableRowOutline/>
        <div>
            <p>{match.league.name} ({match.league.acronym})</p>
            <p class="text-sm font-light">Liga</p>
        </div>
    </div>

    <hr class="my-2">

    <div class="flex items-center gap-3">
        <InfoCircleOutline/>
        <div>
            <p>{match.match_id}</p>
            <p class="text-sm font-light">Spiel-ID</p>
        </div>
    </div>

    <hr class="my-2">

    <div class="flex items-center gap-3">
        <HourglassOutline/>
        <div>
            <p>{match.planned_innings}</p>
            <p class="text-sm font-light">geplante Innings</p>
        </div>
    </div>

    <hr class="my-2">

    <div class="flex items-center gap-3">
        <LinkOutline/>
        {#if match.scoresheet_url}
            <a class="anchor" target="_blank" href="{match.scoresheet_url}">Link zum Scoresheet</a>
        {:else}
            <span>Kein Scoresheet vorhanden.</span>
        {/if}
    </div>
</div>