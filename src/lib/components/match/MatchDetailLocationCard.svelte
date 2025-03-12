<script lang="ts">
  import type { Field } from "bsm.js";
  import { Building2, LandPlot, Link, MapPin } from "lucide-svelte";
  import { MapUtility } from "$lib/service/MapUtility.ts";

  type DisplayedLocation = Omit<
    Field,
    | "created_at"
    | "updated_at"
    | "description_html"
    | "other_information_html"
    | "groundrules_html"
    | "id"
  >;

  interface Props {
    field: DisplayedLocation;
    classes?: string;
    showDividers?: boolean;
  }

  let {
    field,
    classes = "card variant-ghost-surface p-3",
    showDividers = true,
  }: Props = $props();

  const googleMapsLink = $derived(
    MapUtility.buildGoogleMapsURL(
      field.address_addon,
      field.latitude,
      field.longitude,
    ),
  );
  const appleMapsLink = $derived(
    MapUtility.buildAppleMapsURL(
      field.address_addon,
      field.latitude,
      field.longitude,
    ),
  );
</script>

<div class={classes}>
  <div class="flex items-center gap-3">
    <LandPlot />
    <div>
      <p>{field.address_addon} ({field?.name})</p>
      <p class="text-sm font-light">Ballpark</p>
    </div>
  </div>

  <hr class={showDividers ? "" : "!hidden"} />

  <div class="flex items-center gap-3">
    <MapPin />
    <div>
      <p>{field.street}</p>
      <p class="text-sm font-light">Address</p>
    </div>
  </div>

  <hr class={showDividers ? "" : "!hidden"} />

  <div class="flex items-center gap-3">
    <Building2 />
    <div>
      <p>{field.postal_code} {field.city}</p>
      <p class="text-sm font-light">City</p>
    </div>
  </div>

  <hr class={showDividers ? "" : "!hidden"} />

  <div class="flex items-center gap-3 self-end">
    <Link />
    <a class="anchor" target="_blank" href={googleMapsLink}
      >Open in Google Maps</a
    >
  </div>

  <hr class={showDividers ? "" : "!hidden"} />

  <div class="flex items-center gap-3 self-end">
    <Link />
    <a class="anchor" target="_blank" href={appleMapsLink}>Open in Apple Maps</a
    >
  </div>
</div>

<style>
  hr {
    margin: 0.5rem 0;
  }
</style>
