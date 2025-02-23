<script lang="ts">
  import type {Field} from "bsm.js";
  import GenericDetailRow from "$lib/components/utility/GenericDetailRow.svelte";
  import {Diamond, Info, MapPin, Ticket, Users} from "lucide-svelte";
  import {MapUtility} from "$lib/service/MapUtility.ts";

  interface Props {
    field: Field,
  }

  let {field}: Props = $props();

  const mapsLink = $derived(MapUtility.buildGoogleMapsURL(field?.name, field?.latitude, field?.longitude));
</script>

<article class="card p-3 variant-soft dark:border dark:border-surface-500 shadow-xl">
  <GenericDetailRow categoryName="Ballpark Name" rowValue={field.name}>
    {#snippet icon()}
      <div class="primary">
        <Diamond/>
      </div>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow categoryName="Address Addon" rowValue={field.address_addon}>
    {#snippet icon()}
      <Info/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow categoryName="Capacity" rowValue={field.spectator_total}>
    {#snippet icon()}
      <Users/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow categoryName="Seats" rowValue={field.spectator_seats}>
    {#snippet icon()}
      <Ticket/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <div class="container">
    <div class="tertiary">
      <MapPin/>
    </div>
    <dl class="container">
      <dt><strong class="">Address:</strong></dt>
      <dd>
        <a class="anchor" href="{mapsLink}" target="_blank">
          {field.street}, {field.postal_code} {field.city}
        </a>
      </dd>
    </dl>
  </div>

</article>

<style lang="postcss">
    .primary {
        color: rgba(var(--color-primary-500));
    }

    .tertiary {
        color: light-dark(rgba(var(--color-secondary-500)), rgba(var(--color-tertiary-500)));
    }

    .container {
        display: flex;
        align-items: center;
        @apply gap-4;
    }
</style>