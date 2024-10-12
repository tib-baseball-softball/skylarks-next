<script lang="ts">
  import {invalidateAll} from "$app/navigation";
  import {client} from "$lib/pocketbase";
  import {getDrawerStore, getToastStore, type ToastSettings,} from "@skeletonlabs/skeleton";
  import {CloseOutline} from "flowbite-svelte-icons";
  import type {ClubsResponse, UsersUpdate} from "$lib/model/pb-types";
  import {authModel} from "$lib/pocketbase/Auth";
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";

  const toastStore = getToastStore();
  const drawerStore = getDrawerStore();

  const model = $authModel as CustomAuthModel

  const toastSettingsSuccess: ToastSettings = {
    message: "Club data saved successfully.",
    background: "variant-filled-success",
  };

  const toastSettingsError: ToastSettings = {
    message: "An error occurred while saving club data.",
    background: "variant-filled-error",
  };

  const form: ClubsResponse = $state(
    $drawerStore.meta.club ?? {
      id: "",
      name: "",
      bsm_id: 0,
      bsm_api_key: "",
      acronym: "",
      admins: [], // TODO: handle
    },
  );

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ClubsResponse | null = null;

    try {
      if (form.id) {
        result = await client
          .collection("clubs")
          .update<ClubsResponse>(form.id, form);
      } else {
        // a user creating a club becomes its first admin
        form.admins.push(model.id)

        result = await client
          .collection("clubs")
          .create<ClubsResponse>(form);

        // a user needs to become a member of the new club
        await client.collection("users").update<UsersUpdate>(model.id, {
          "club+": result.id
        })
      }
    } catch {
      toastStore.trigger(toastSettingsError);
      drawerStore.close();
    }

    if (result) {
      toastStore.trigger(toastSettingsSuccess);
    }
    invalidateAll()
    drawerStore.close();
  }
</script>

<article class="p-6">
    <div class="flex items-center gap-5">
        <button
            aria-label="cancel and close"
            class="btn variant-ghost-surface"
            onclick={drawerStore.close}
        >
            <CloseOutline/>
        </button>
        <header class="text-xl font-semibold">
            {#if form.id}
                <h2 class="h3">Edit Club "{form?.name}"</h2>
            {:else}
                <h2 class="h3">Create new Club</h2>
            {/if}
        </header>
    </div>

    <form onsubmit={submitForm} class="mt-4 space-y-3">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-2 md:gap-3 xl:gap-4">
            <input
                name="id"
                autocomplete="off"
                class="input"
                type="hidden"
                readonly
                bind:value={form.id}
            />

            <label class="label">
                Name
                <input
                    name="name"
                    class="input"
                    required
                    type="text"
                    bind:value={form.name}
                />
            </label>

            <label class="label">
                Acronym
                <input
                    name="acronym"
                    class="input"
                    type="text"
                    bind:value={form.acronym}
                />
            </label>

            <label class="label">
                BSM Club ID
                <input
                    name="bsm_id"
                    class="input"
                    type="number"
                    required
                    bind:value={form.bsm_id}
                >
                <span class="text-sm">Can be found in the BSM address bar while editing
                    (e.g. <span class="italic">https://bsm.baseball-softball.de/clubs/xxx/edit</span>).
                    Needs to be set for the app to function properly.
                </span>
            </label>

            <label class="label">
                BSM API Key
                <input
                    name="bsm_api_key"
                    class="input"
                    type="text"
                    bind:value={form.bsm_api_key}
                />
                <span class="text-sm">
                    Must be created in BSM in a user's account that has the role "Team Administration".
                    If set, all game events for the club can be automatically imported.
                </span>
            </label>
        </div>

        <hr class="!my-5"/>

        <div class="flex justify-center gap-3">
            <button type="submit" class="mt-2 btn variant-ghost-primary">
                Submit
            </button>
        </div>
    </form>
</article>
