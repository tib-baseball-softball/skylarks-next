<script lang="ts">
  import {invalidate} from "$app/navigation";
  import type {ExpandedClub} from "$lib/model/ExpandedResponse";
  import type {LocationsResponse} from "$lib/model/pb-types";
  import {save} from "$lib/pocketbase/RecordOperations";
  import {getToastStore, type ToastSettings} from "@skeletonlabs/skeleton";
  import {Edit, Plus} from "lucide-svelte";
  //@ts-ignore
  import * as Sheet from "$lib/components/utility/sheet/index.js";
  import LeafletMapCoordinatePicker from "$lib/components/map/LeafletMapCoordinatePicker.svelte";

  // Gail S. Halvorsen Park coordinates
  const DEFAULT_LATITUDE = 52.482762;
  const DEFAULT_LONGITUDE = 13.407932;

  interface Props {
    club: ExpandedClub;
    location?: LocationsResponse | null;
    buttonClasses?: string;
  }

  let { location = null, club, buttonClasses = "" }: Props = $props();

  const toastStore = getToastStore();

  const toastSettingsSuccess: ToastSettings = {
    message: "Location data saved successfully.",
    background: "variant-filled-success",
  };

  const toastSettingsError: ToastSettings = {
    message: "An error occurred while saving Location.",
    background: "variant-filled-error",
  };

  const form = $state(
    location ?? {
      id: "",
      name: "",
      street: "",
      postal_code: "",
      city: "",
      address_addon: "",
      country: "",
      human_country: "",
      internal_name: "",
      bsm_id: 0,
      latitude: DEFAULT_LATITUDE,
      longitude: DEFAULT_LONGITUDE,
      club: club.id,
    },
  );

  let open = $state(false)

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();
    form.club = club.id;

    let result: LocationsResponse | null = null;

    try {
      result = await save<LocationsResponse>("locations", form);
    } catch {
      toastStore.trigger(toastSettingsError);
    }

    if (result) {
      toastStore.trigger(toastSettingsSuccess);
    }
    await invalidate("club:single");
    open = false;
  }
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger class={buttonClasses}>
    {#if form.id}
      <Edit />
    {:else}
      <Plus />
      <span>Create new</span>
    {/if}
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header class="my-3">
      <h2 class="h2">
        {#if form.id}
          Edit Location "{form?.address_addon ?? form?.internal_name}"
        {:else}
          Add New Location
        {/if}
      </h2>
    </header>

    <form onsubmit={submitForm}>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label for="address_addon" class="">Address Addon*</label>
          <input
            required
            type="text"
            id="address_addon"
            name="address_addon"
            bind:value={form.address_addon}
            class="input"
          />
          <span class="font-light text-sm"
            >Often used in BSM data for the field name.</span
          >
        </div>

        <div>
          <label for="name" class="label">Name (Sport)*</label>
          <input
            required
            type="text"
            id="name"
            name="name"
            bind:value={form.name}
            class="input"
          />
        </div>

        <div>
          <label for="street" class="label">Street</label>
          <input
            type="text"
            id="street"
            name="street"
            bind:value={form.street}
            class="input"
          />
        </div>

        <div>
          <label for="postal_code" class="label">Postal Code</label>
          <input
            type="text"
            id="postal_code"
            name="postal_code"
            bind:value={form.postal_code}
            class="input"
          />
        </div>

        <div>
          <label for="city" class="label">City</label>
          <input
            type="text"
            id="city"
            name="city"
            bind:value={form.city}
            class="input"
          />
        </div>

        <div>
          <label for="bsm_id" class="label">BSM ID</label>
          <input
            readonly
            autocomplete="off"
            type="number"
            id="bsm_id"
            name="bsm_id"
            bind:value={form.bsm_id}
            class="input"
          />
        </div>

        <div>
          <label for="country" class="label">Country Code</label>
          <input
            type="text"
            id="country"
            name="country"
            maxlength="2"
            bind:value={form.country}
            class="input"
          />
        </div>

        <div>
          <label for="human_country" class="label">Country Name</label>
          <input
            type="text"
            id="human_country"
            name="human_country"
            bind:value={form.human_country}
            class="input"
          />
        </div>

        <div class="md:col-span-2">
          <label for="internal_name" class="label">Internal Name</label>
          <input
            type="text"
            id="internal_name"
            name="internal_name"
            bind:value={form.internal_name}
            class="input"
          />
          <span class="font-light text-sm"
            >Custom name field that will not be overwritten by automatic
            imports.</span
          >
        </div>

        <div>
          <label for="latitude" class="label">Latitude*</label>
          <input
            required
            type="number"
            step="any"
            id="latitude"
            name="latitude"
            bind:value={form.latitude}
            class="input"
          />
        </div>

        <div>
          <label for="longitude" class="label">Longitude*</label>
          <input
            required
            type="number"
            step="any"
            id="longitude"
            name="longitude"
            bind:value={form.longitude}
            class="input"
          />
        </div>

        <label class="md:col-span-2">
          Coordinate Map
          <span class="text-sm font-light block">Select a location on the map to set coordinates automatically</span>
          <LeafletMapCoordinatePicker bind:latitude={form.latitude} bind:longitude={form.longitude}/>
        </label>

      </div>

      <div class="mt-4 flex justify-between items-center">
        <button type="submit" class="btn variant-ghost-primary">Submit</button>
        <button type="reset" class="btn variant-ghost ms-2">Reset Form</button>
      </div>
    </form>
  </Sheet.Content>
</Sheet.Root>
