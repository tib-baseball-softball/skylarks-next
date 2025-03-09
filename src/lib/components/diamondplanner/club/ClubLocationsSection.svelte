<script lang="ts">
  import { invalidate, invalidateAll } from "$app/navigation";
  import DeleteButton from "$lib/components/utility/DeleteButton.svelte";
  import type { LocationsResponse } from "$lib/model/pb-types";
  import { client } from "$lib/pocketbase/index.svelte";

  interface Props {
    locations: LocationsResponse[];
  }

  let { locations }: Props = $props();

  function locationDeleteAction(id: string) {
    client.collection("locations").delete(id);
    invalidate("club:single");
  }
</script>

<ul class="variant-soft rounded-token p-4 shadow-xl text-sm">
  {#each locations as location, index}
    <li class="location-grid">
      <h3 class="font-bold">{location.address_addon ? location.address_addon : location.internal_name} ({location.name})</h3>
      <div class="">
        {location?.street}, {location.postal_code}
        {location.city}
      </div>

      <div class="place-self-end">
        <DeleteButton
          id={location.id}
          modelName="Location"
          action={locationDeleteAction}
          classes="variant-ghost-error mx-1 btn-sm"
        />
      </div>
    </li>

    {#if index < locations.length - 1}
      <hr class="my-3" />
    {/if}
  {/each}
</ul>

<style lang="postcss">
  .location-grid {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    align-items: center;
    gap: 1rem;
  }
</style>
