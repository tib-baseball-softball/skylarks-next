<script lang="ts">
  import {invalidateAll} from "$app/navigation";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import {getDrawerStore, getToastStore, type ToastSettings,} from "@skeletonlabs/skeleton";
  import type {ClubsResponse, UsersResponse, UsersUpdate} from "$lib/model/pb-types";
  import type {CustomAuthModel, ExpandedClub} from "$lib/model/ExpandedResponse";
  import MultiSelectCombobox from "$lib/components/utility/MultiSelectCombobox.svelte";
  import {X} from "lucide-svelte";

  const toastStore = getToastStore();
  const drawerStore = getDrawerStore();

  const model = authSettings.record as CustomAuthModel;

  const toastSettingsSuccess: ToastSettings = {
    message: "Club data saved successfully.",
    background: "variant-filled-success",
  };

  const toastSettingsError: ToastSettings = {
    message: "An error occurred while saving club data.",
    background: "variant-filled-error",
  };

  const form: ExpandedClub = $state(
      $drawerStore.meta.club ?? {
        id: "",
        name: "",
        bsm_id: 0,
        bsm_api_key: "",
        signup_key: "",
        acronym: "",
        admins: [],
      },
  );

  let selectedAdmins: UsersResponse[] = $state(form.expand.admins);

  const allUsersForClub = client.collection("users").getFullList<UsersResponse>({
    filter: `club ?~ '${$drawerStore.meta.club.id}'`
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ClubsResponse | null = null;

    try {
      if (form.id) {
        form.admins = selectedAdmins.map((admin) => {
          return admin.id;
        });

        result = await client
            .collection("clubs")
            .update<ClubsResponse>(form.id, form);
      } else {
        // a user creating a club becomes its first admin
        form.admins.push(model.id);

        result = await client
            .collection("clubs")
            .create<ClubsResponse>(form);

        // a user needs to become a member of the new club
        await client.collection("users").update<UsersUpdate>(model.id, {
          "club+": result.id
        });
      }
    } catch {
      toastStore.trigger(toastSettingsError);
      drawerStore.close();
    }

    if (result) {
      toastStore.trigger(toastSettingsSuccess);
    }
    await invalidateAll();
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
      <X/>
    </button>
    <header class="text-xl font-semibold">
      {#if form.id}
        <h2 class="h3">Edit Club "{form?.name}"</h2>
      {:else}
        <h2 class="h3">Create new Club</h2>
      {/if}
    </header>
  </div>

  <form class="mt-4 space-y-3" onsubmit={submitForm}>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-2 md:gap-3 xl:gap-4">
      <input
              autocomplete="off"
              bind:value={form.id}
              class="input"
              name="id"
              readonly
              type="hidden"
      />

      <label class="label">
        <span>Name</span>
        <input
                bind:value={form.name}
                class="input"
                name="name"
                required
                type="text"
        />
      </label>

      <label class="label">
                <span>
                Acronym
                </span>
        <input
                bind:value={form.acronym}
                class="input"
                name="acronym"
                type="text"
        />
      </label>

      <label class="label">
                <span>
                BSM Club ID
                </span>
        <input
                bind:value={form.bsm_id}
                class="input"
                name="bsm_id"
                required
                type="number"
        >
        <span class="text-sm">Can be found in the BSM address bar while editing
                    (e.g. <span class="italic">https://bsm.baseball-softball.de/clubs/xxx/edit</span>).
                    Needs to be set for the app to function properly.
                </span>
      </label>

      <label class="label">
                <span>
                BSM API Key
                </span>
        <input
                bind:value={form.bsm_api_key}
                class="input"
                name="bsm_api_key"
                type="text"
                placeholder="current value not shown for security"
        />
        <span class="text-sm">
                    Must be created in BSM in a user's account that has the role "Team Administration".
                    If set, all game events for the club can be automatically imported.
                </span>
      </label>

      <label class="label">
                <span>
                Signup Key
                </span>
        <input
                bind:value={form.signup_key}
                class="input"
                name="signup_key"
                placeholder="minimum 8 characters"
                minlength="8"
                required
                type="text"
        />
        <span class="text-sm">
                    A valid signup key needs to be entered upon user account creation.
                    New users are automatically added as members to the club corresponding to the signup key used.
                </span>
      </label>

      {#if form.id /* we are editing and might want to change admins */}
        <label class="label space-y-3 md:col-span-2">
          <span>Club Admins</span><br>

          {#await allUsersForClub then users}
            <MultiSelectCombobox itemName="Admin" bind:selectedItems={selectedAdmins} allItems={users}/>
          {/await}

          <span class="text-sm">
                        It is not possible to remove the last admin from a club.
                        <br>
                        Caution: You can remove yourself as admin, losing all access rights!
                    </span>
        </label>
      {/if}
    </div>

    <hr class="!my-5"/>

    <div class="flex justify-center gap-3">
      <button class="mt-2 btn variant-ghost-primary" type="submit">
        Submit
      </button>
    </div>
  </form>
</article>
