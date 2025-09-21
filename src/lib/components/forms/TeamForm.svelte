<script lang="ts">
  import {invalidate} from "$app/navigation";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/model/ExpandedResponse";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  // @ts-ignore
  import {Tabs} from "bits-ui";
  import {Edit, Plus} from "lucide-svelte";
  import type {ClubsResponse, UsersResponse} from "$lib/model/pb-types.ts";
  import MultiSelectCombobox from "$lib/components/utility/MultiSelectCombobox.svelte";
  //@ts-ignore
  import * as Sheet from "$lib/components/utility/sheet/index.js";
  import {toastController} from "$lib/service/ToastController.svelte.ts";

  interface Props {
    club: ClubsResponse,
    team: ExpandedTeam | null,
    buttonClasses?: string,
    showLabel?: boolean,
  }

  let {club, team, buttonClasses = "", showLabel = true}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  let open = $state(false);


  const form: Partial<ExpandedTeam> = $state(
      team ?? {
        id: "",
        name: "",
        age_group: "adults",
        signup_key: "",
        club: club.id, // no binding, cannot be changed via this form
        description: "",
        admins: [],
      },
  );

  let selectedAdmins: UsersResponse[] = $state(form?.expand?.admins ?? []);

  const allTeamMembers = client.collection("users").getFullList<UsersResponse>({
    filter: `teams ?~ '${team?.id}'`,
    requestKey: `allTeamMembers-${team?.id}`
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ExpandedTeam | null = null;

    form.admins = selectedAdmins.map((admin) => {
      return admin.id;
    });

    try {
      if (form.id) {
        result = await client
            .collection("teams")
            .update<ExpandedTeam>(form.id, form);
      } else {
        // a user creating a team becomes its first admin
        form.admins.push(authRecord?.id);

        result = await client
            .collection("teams")
            .create<ExpandedTeam>(form);
      }
    } catch {
      toastController.triggerGenericFormErrorMessage("Team");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Team");
      open = false;
      await invalidate("teams:list");
      await invalidate("club:single");
      await invalidate("nav:load");
    }
  }
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger class={buttonClasses}>
    {#if form.id}

      <Edit/>
      {#if showLabel}
        <span>Edit Team</span>
      {/if}

    {:else}

      <Plus/>
      {#if showLabel}
        <span>Create new</span>
      {/if}

    {/if}
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header class="text-xl font-semibold">
      {#if form.id}
        <h2 class="h3">Edit Team "{form?.name}"</h2>
      {:else}
        <h2 class="h3">Create new Team</h2>
      {/if}
    </header>

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
                  value={club?.name}
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
          <Tabs.Root bind:value={form.age_group} class="dark:preset-outlined-surface-600-400">
            <Tabs.List class="tabs-list">
              <Tabs.Trigger value="adults" class="tabs-trigger flex-grow">Adults</Tabs.Trigger>
              <Tabs.Trigger value="minors" class="tabs-trigger flex-grow">Minors</Tabs.Trigger>
            </Tabs.List>
          </Tabs.Root>
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

        <label class="label space-y-3 md:col-span-2">
          <span>Team Admins</span><br>

          {#await allTeamMembers then users}
            <MultiSelectCombobox itemName="Admin" bind:selectedItems={selectedAdmins} allItems={users}/>
          {/await}
        </label>
      </div>

      <hr class="my-5!"/>

      <div class="flex justify-center gap-3">
        <button type="submit" class="mt-2 btn preset-filled-primary-500">
          Submit
        </button>
      </div>
    </form>
  </Sheet.Content>
</Sheet.Root>
