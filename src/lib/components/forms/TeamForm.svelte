<script lang="ts">
  import {invalidate} from "$app/navigation";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/model/ExpandedResponse";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import {getDrawerStore, getToastStore, RadioGroup, RadioItem, type ToastSettings,} from "@skeletonlabs/skeleton";
  import {X} from "lucide-svelte";

  const toastStore = getToastStore();
  const drawerStore = getDrawerStore();

  const authModel = authSettings.record as CustomAuthModel;

  const toastSettingsSuccess: ToastSettings = {
    message: "Team data saved successfully.",
    background: "variant-filled-success",
  };

  const toastSettingsError: ToastSettings = {
    message: "An error occurred while saving team data.",
    background: "variant-filled-error",
  };

  const form: ExpandedTeam = $state(
      $drawerStore.meta.team ?? {
        id: "",
        name: "",
        age_group: "",
        signup_key: "",
        club: $drawerStore.meta.club.id, // no binding, cannot be changed via this form
        description: "",
        admins: [],
      },
  );

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ExpandedTeam | null = null;

    try {
      if (form.id) {
        result = await client
            .collection("teams")
            .update<ExpandedTeam>(form.id, form);
      } else {
        // a user creating a team becomes its first admin
        form.admins.push(authModel?.id);

        result = await client
            .collection("teams")
            .create<ExpandedTeam>(form);
      }
    } catch {
      toastStore.trigger(toastSettingsError);
      drawerStore.close();
    }

    if (result) {
      toastStore.trigger(toastSettingsSuccess);
    }
    await invalidate("teams:list");
    await invalidate("club:single");
    await invalidate("nav:load");
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
        <h2 class="h3">Edit Team "{form?.name}"</h2>
      {:else}
        <h2 class="h3">Create new Team</h2>
      {/if}
    </header>
  </div>

  <form onsubmit={submitForm} class="mt-4 space-y-3">
    <div class="grid grid-cols-1 md:grid-cols-2 gap-2 lg:gap-3 xl:gap-4">
      <input
              name="id"
              autocomplete="off"
              class="input"
              type="hidden"
              readonly
              bind:value={form.id}
      />

      <label class="label col-span-2 md:col-span-1">
        Name
        <input
                name="title"
                class="input"
                required
                type="text"
                bind:value={form.name}
        />
      </label>

      <label class="label col-span-2 md:col-span-1">
        Club
        <input
                name="id"
                autocomplete="off"
                class="input"
                type="text"
                readonly
                value={$drawerStore.meta.club?.name}
        />
      </label>

      <label class="label col-span-2">
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
                    New users are automatically added as members to the team corresponding to the signup key used.
                </span>
      </label>

      <label class="label flex flex-col gap-1 col-span-2">
        Type
        <RadioGroup>
          <RadioItem
                  bind:group={form.age_group}
                  name="age_group"
                  value={"adults"}
          >
            Adults
          </RadioItem>
          <RadioItem
                  bind:group={form.age_group}
                  name="type"
                  value={"minors"}
          >
            Minors
          </RadioItem>
        </RadioGroup>
      </label>

      <label class="label col-span-2">
        Description
        <textarea
                name="desc"
                class="textarea"
                rows="8"
                bind:value={form.description}
        ></textarea>
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
