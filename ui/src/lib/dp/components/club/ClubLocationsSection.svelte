<script lang="ts">
  import { invalidate } from "$app/navigation";
  import LocationForm from "$lib/dp/components/forms/LocationForm.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import { authSettings, client } from "$lib/dp/client.svelte.js";
  import type {
    CustomAuthModel,
    ExpandedClub,
  } from "$lib/dp/types/ExpandedResponse.ts";
  import type { LocationsResponse } from "$lib/dp/types/pb-types.ts";
  import { MapUtility } from "$lib/dp/service/MapUtility.ts";
  import { Map } from "lucide-svelte";
  import { Collection } from "$lib/dp/enum/Collection.ts";

  interface Props {
    club: ExpandedClub;
    locations: LocationsResponse[];
  }

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  const { club, locations }: Props = $props();

  function locationDeleteAction(id: string) {
    client.collection(Collection.Locations).delete(id);
    invalidate("club:locations");
  }
</script>

<ul class="preset-tonal-surface rounded-base shadow-xl">
  {#each locations as location, index (location.id)}
    <li class="location-grid">
      <h3>
        {location.internal_name
          ? location.internal_name
          : location.address_addon} ({location.name})
      </h3>

      <p>
        {location?.street}, {location.postal_code}
        {location.city}
      </p>

      <div class="actions">
        <a
          href={MapUtility.buildGoogleMapsURL(
            location.address_addon,
            location.latitude,
            location.longitude,
          )}
          class="btn btn-icon preset-tonal-primary border border-primary-500"
          title="open location in Google Maps"
          target="_blank"
        >
          <Map />
        </a>

        {#if club?.admins.includes(authRecord.id)}
          <LocationForm
            {club}
            {location}
            triggerVariant="tonal-tertiary"
            triggerIcon={true}
          />

          <DeleteButton
            id={location.id}
            modelName="Location"
            action={locationDeleteAction}
            classes="preset-tonal-error border border-error-500 btn-icon"
          />
        {/if}
      </div>
    </li>

    {#if index < locations.length - 1}
      <hr />
    {/if}
  {:else}
    <p>This club doesn't have saved locations yet.</p>
  {/each}
</ul>

{#if club?.admins.includes(authRecord.id)}
  <div class="add-button-wrapper">
    <LocationForm
      {club}
      location={null}
      triggerVariant="tonal-primary"
      triggerSpaced={true}
    />
  </div>
{/if}

<style>
  .location-grid {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    align-items: center;
    gap: 1rem;
  }

  .actions {
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 2);
    align-items: center;
    justify-self: end;

    @media (min-width: 48rem) {
      flex-direction: row;
    }
  }

  ul {
    font-size: var(--text-sm);
    padding: calc(var(--spacing) * 4);
  }

  h3 {
    font-weight: bold;
    overflow: hidden;
  }

  hr {
    margin-block: calc(var(--spacing) * 2);
  }
  
  .add-button-wrapper {
    margin-block: calc(var(--spacing) * 4);
  }
</style>
