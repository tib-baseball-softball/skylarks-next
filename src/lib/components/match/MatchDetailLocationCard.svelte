<script lang="ts">
    import type {Match} from "bsm.js";
    import {
        BuildingOutline, HomeOutline,
        LinkOutline, MapPinOutline
    } from "flowbite-svelte-icons";

    interface Props {
        match: Match;
    }

    let { match }: Props = $props();

    function buildMapsURL(): string {
        const baseURL = "https://www.google.com/maps/search/"

        const params = new URLSearchParams({
            api: "1",
            map_action: "map",
            query_place_id: match?.field.name,
            query: `${match?.field.latitude}, ${match?.field.longitude}`
        })

        return `${baseURL}?${params.toString()}`
    }
</script>

<div class="card variant-ghost-surface p-3">
    <div class="flex items-center gap-3">
        <HomeOutline/>
        <div>
            <p>{match?.field.name}</p>
            <p class="text-sm font-light">Ballpark</p>
        </div>
    </div>

    <hr class="my-2">

    <div class="flex items-center gap-3">
        <MapPinOutline/>
        <div>
            <p>{match?.field.street} </p>
            <p class="text-sm font-light">Adresse</p>
        </div>
    </div>

    <hr class="my-2">

    <div class="flex items-center gap-3">
        <BuildingOutline/>
        <div>
            <p>{match?.field.postal_code} {match?.field.city}</p>
            <p class="text-sm font-light">Stadt</p>
        </div>
    </div>

    <hr class="my-2">

    <div class="flex items-center gap-3 self-end">
        <LinkOutline/>
        <a class="anchor" target="_blank" href="{buildMapsURL()}">In Google Maps öffnen</a>
    </div>
</div>