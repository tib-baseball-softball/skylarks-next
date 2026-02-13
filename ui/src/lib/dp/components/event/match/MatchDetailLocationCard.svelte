<script lang="ts">
  import type {Field} from "bsm.js";
  import {Building2, LandPlot, Link, MapPin} from "lucide-svelte";
  import {MapUtility} from "$lib/dp/service/MapUtility.ts";

  type DisplayedLocation = Omit<
    Field,
    | "created_at"
    | "updated_at"
    | "description_html"
    | "other_information_html"
    | "groundrules_html"
    | "id"
  >

  interface Props {
    field: DisplayedLocation;
    classes?: string;
    showDividers?: boolean;
  }

  const {
    field,
    classes = "card preset-tonal-surface border border-surface-500 p-3",
    showDividers = true,
  }: Props = $props();

  const googleMapsLink = $derived(
    MapUtility.buildGoogleMapsURL(field.address_addon, field.latitude, field.longitude)
  );
  const appleMapsLink = $derived(
    MapUtility.buildAppleMapsURL(field.address_addon, field.latitude, field.longitude)
  );
</script>

<div class="root {classes}">
  <div class="location-row">
    <LandPlot/>
    <div>
      <p>{field.address_addon} ({field?.name})</p>
      <p class="label">Ballpark</p>
    </div>
  </div>

  <hr class={showDividers ? "" : "hidden!"}/>

  <div class="location-row">
    <MapPin/>
    <div>
      <p>{field.street}</p>
      <p class="label">Address</p>
    </div>
  </div>

  <hr class={showDividers ? "" : "hidden!"}/>

  <div class="location-row">
    <Building2/>
    <div>
      <p>{field.postal_code} {field.city}</p>
      <p class="label">City</p>
    </div>
  </div>

  <hr class={showDividers ? "" : "hidden!"}/>

  <div class="location-row link">
    <Link/>
    <a class="anchor" href={googleMapsLink} target="_blank"
    >Open in Google Maps</a
    >
  </div>

  <hr class={showDividers ? "" : "hidden!"}/>

  <div class="location-row link">
    <Link/>
    <a class="anchor" href={appleMapsLink} target="_blank">Open in Apple Maps</a
    >
  </div>
</div>

<style>
  .root {
    display: flex;
    flex-direction: column;
  }

  .location-row {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 3);

    &.link {
      align-self: flex-end;
    }
  }

  .label {
    font-size: 0.875rem; /* text-sm */
    font-weight: 300; /* font-light */
  }

  hr {
    margin: 0.5rem 0;
  }
</style>
