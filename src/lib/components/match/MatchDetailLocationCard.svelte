<script lang="ts">
  import type {Match} from "bsm.js";
  import {Building2, LandPlot, Link, MapPin} from "lucide-svelte";
  import {MapUtility} from "$lib/service/MapUtility.ts";

  interface Props {
    match: Match;
  }

  let {match}: Props = $props();

  const mapsLink = $derived(MapUtility.buildGoogleMapsURL(match?.field.name, match?.field.latitude, match?.field.longitude));
</script>

<div class="card variant-ghost-surface p-3">
  <div class="flex items-center gap-3">
    <LandPlot/>
    <div>
      <p>{match?.field.name}</p>
      <p class="text-sm font-light">Ballpark</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <MapPin/>
    <div>
      <p>{match?.field.street} </p>
      <p class="text-sm font-light">Adresse</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <Building2/>
    <div>
      <p>{match?.field.postal_code} {match?.field.city}</p>
      <p class="text-sm font-light">Stadt</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3 self-end">
    <Link/>
    <a class="anchor" target="_blank" href="{mapsLink}">In Google Maps öffnen</a>
  </div>
</div>