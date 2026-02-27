<script lang="ts">
  import {Plus, SquarePen} from "lucide-svelte";
  import {invalidate} from "$app/navigation";
  import TabsRadioGroup from "$lib/dp/components/formElements/TabsRadioGroup.svelte";
  import MultiSelectCombobox from "$lib/dp/components/formElements/MultiSelectCombobox.svelte";
  //@ts-ignore
  import * as Sheet from "$lib/dp/components/modal/sheet";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import type {ClubsResponse, UsersResponse} from "$lib/dp/types/pb-types.ts";
  import {Collection} from "$lib/dp/enum/Collection.ts";

  interface Props {
    club: ClubsResponse;
    team: ExpandedTeam | null;
    triggerVariant?: "filled-primary" | "tonal-primary" | "tonal-secondary" | "tonal-tertiary" | "tonal-surface";
    triggerSize?: "default" | "sm";
    triggerIcon?: boolean;
    triggerSpaced?: boolean;
    showLabel?: boolean;
  }

  const {
    club,
    team,
    triggerVariant = "tonal-primary",
    triggerSize = "default",
    triggerIcon = false,
    triggerSpaced = false,
    showLabel = true,
  }: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  let open = $state(false);

  function formFromProps(data: ExpandedTeam | null) {
    return data ?? {
      id: "",
      name: "",
      age_group: "adults",
      signup_key: "",
      club: club.id, // no binding, cannot be changed via this form
      description: "",
      admins: [],
    };
  }

  let form: Partial<ExpandedTeam> & { age_group: string } = $derived.by(() => {
    const formData = $state(formFromProps(team));
    return formData;
  });

  let selectedAdmins: UsersResponse[] = $derived(form?.expand?.admins ?? []);

  const allTeamMembers = $derived(client.collection(Collection.Users).getFullList<UsersResponse>({
    filter: `teams ?~ '${team?.id}'`,
    sort: "+last_name",
    requestKey: `allTeamMembers-${team?.id}`,
  }));

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ExpandedTeam | null = null;

    form.admins = selectedAdmins.map((admin) => {
      return admin.id;
    });

    try {
      if (form.id) {
        result = await client.collection(Collection.Teams).update<ExpandedTeam>(form.id, form);
      } else {
        // a user creating a team becomes its first admin
        form.admins.push(authRecord?.id);

        result = await client.collection(Collection.Teams).create<ExpandedTeam>(form);
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
  <Sheet.Trigger
    class={[
      "btn",
      "trigger-button",
      "team-form-trigger",
      `trigger-variant-${triggerVariant}`,
      triggerSize === "sm" && "btn-sm",
      triggerIcon && "btn-icon",
      triggerSpaced && "trigger-spaced",
      triggerVariant === "filled-primary" && "preset-filled-primary-500",
      triggerVariant === "tonal-primary" && "preset-tonal-primary border-primary",
      triggerVariant === "tonal-secondary" && "preset-tonal-secondary border-secondary",
      triggerVariant === "tonal-tertiary" && "preset-tonal-tertiary border-tertiary",
      triggerVariant === "tonal-surface" && "preset-tonal-surface",
    ]}
  >
    {#if form.id}

      <SquarePen/>
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

    <header>
      {#if form.id}
        <h2 class="h3">Edit Team "{form?.name}"</h2>
      {:else}
        <h2 class="h3">Create new Team</h2>
      {/if}
    </header>

    <form onsubmit={submitForm}>
      <div class="edit-form-grid">
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
            name="title"
            required
            type="text"
          />
        </label>

        <label class="label">
          <span>Club</span>
          <input
            autocomplete="off"
            class="input"
            name="id"
            readonly
            type="text"
            value={club?.name}
          />
        </label>

        <label class="label">
                <span>
                Signup Key
                </span>
          <input
            bind:value={form.signup_key}
            class="input"
            minlength="8"
            name="signup_key"
            placeholder="minimum 8 characters"
            required
            type="text"
          />
          <span class="hint">
                    A valid signup key needs to be entered upon user account creation.
                    New users are automatically added as members to the team corresponding to the signup key used.
                </span>
        </label>

        <div>
          <TabsRadioGroup
            bind:value={form.age_group}
            label="Age Group"
            listClass="tabs-list input"
            name="age_group"
            options={["adults", "minors"]}
            required={true}
          />
        </div>

        <label class="label col-span-2">
          Description
          <textarea
            bind:value={form.description}
            class="textarea"
            name="desc"
            rows="8"
          ></textarea>
        </label>

        <label class="label space-y-3 field-wide">
          <span>Team Admins</span><br>

          {#await allTeamMembers then users}
            <MultiSelectCombobox itemName="Admin" bind:selectedItems={selectedAdmins} allItems={users}/>
          {/await}
        </label>
      </div>

      <hr/>

      <div class="flex justify-center gap-3">
        <button class="mt-2 btn preset-filled-primary-500" type="submit">
          Submit
        </button>
      </div>
    </form>
  </Sheet.Content>
</Sheet.Root>

<style>
  :global(.team-form-trigger .lucide-icon) {
    flex-shrink: 1;
  }

  .hint {
    font-size: var(--text-sm);
    font-weight: var(--font-weight-light);
  }

  hr {
    margin-block: calc(var(--spacing) * 5);
  }

  .edit-form-grid {
    gap: calc(var(--spacing) * 4);
  }
</style>
