<script lang="ts">
  import type {Field} from "bsm.js";
  import GenericDetailRow from "$lib/components/utility/GenericDetailRow.svelte";
  import {Diamond, Info, MapPin, Ticket, Users} from "lucide-svelte";

  interface Props {
    field: Field,
  }

  let {field}: Props = $props();
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

  <GenericDetailRow
          categoryName="Address"
          rowValue="{field.street}, {field.postal_code} {field.city}"
  >
    {#snippet icon()}
      <div class="tertiary">
        <MapPin/>
      </div>
    {/snippet}
  </GenericDetailRow>
</article>

<style>
    .primary {
        color: rgba(var(--color-primary-500));
    }

    .tertiary {
        color: light-dark(rgba(var(--color-secondary-500)), rgba(var(--color-tertiary-500)));
    }
</style>