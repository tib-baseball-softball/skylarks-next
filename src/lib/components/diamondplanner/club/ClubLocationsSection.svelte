<script lang="ts">
  import {invalidate} from "$app/navigation";
  import LocationForm from "$lib/components/forms/LocationForm.svelte";
  import DeleteButton from "$lib/components/utility/DeleteButton.svelte";
  import type {CustomAuthModel, ExpandedClub,} from "$lib/model/ExpandedResponse";
  import type {LocationsResponse} from "$lib/model/pb-types";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";

  interface Props {
    club: ExpandedClub;
    locations: LocationsResponse[];
  }

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  let { club, locations }: Props = $props();

  function locationDeleteAction(id: string) {
    client.collection("locations").delete(id);
    invalidate("club:single");
  }
</script>

<ul class="variant-soft rounded-token p-4 shadow-xl text-sm">
  {#each locations as location, index (location.id)}
    <li class="location-grid">
      <h3 class="font-bold">
        {location.address_addon
          ? location.address_addon
          : location.internal_name} ({location.name})
      </h3>
      <p>
        {location?.street}, {location.postal_code}
        {location.city}
      </p>

      <div class="place-self-end">
        {#if club?.admins.includes(authRecord.id)}
          <LocationForm
            {club}
            {location}
            buttonClasses="btn btn-icon variant-ghost-tertiary my-3"
          />

          <DeleteButton
            id={location.id}
            modelName="Location"
            action={locationDeleteAction}
            classes="variant-ghost-error btn-icon ms-1"
          />
        {/if}
      </div>
    </li>

    {#if index < locations.length - 1}
      <hr class="my-2 md:my-0"/>
    {/if}

    {:else }
      <p>This club doesn't have saved locations yet.</p>
  {/each}
</ul>

{#if club?.admins.includes(authRecord.id)}
  <LocationForm
    {club}
    location={null}
    buttonClasses="btn variant-ghost-primary my-3"
  />
{/if}

<style lang="postcss">
  .location-grid {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    align-items: center;
    gap: 1rem;
  }
</style>
