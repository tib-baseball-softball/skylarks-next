<script lang="ts">
  import { Plus, SquarePen } from "lucide-svelte";
  import { invalidateAll } from "$app/navigation";
  //@ts-ignore
  import * as Sheet from "$lib/dp/components/modal/sheet";
  import { authSettings, client } from "$lib/dp/client.svelte.js";
  import { toastController } from "$lib/dp/service/ToastController.svelte.ts";
  import type {
    CustomAuthModel,
    ExpandedAnnouncement,
  } from "$lib/dp/types/ExpandedResponse.ts";
  import type {
    AnnouncementsResponse,
    ClubsResponse,
    TeamsResponse,
  } from "$lib/dp/types/pb-types.ts";

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  interface Props {
    announcement: ExpandedAnnouncement | null;
    club: ClubsResponse | null;
    team: TeamsResponse | null;
    triggerVariant?:
      | "filled-primary"
      | "tonal-primary"
      | "tonal-secondary"
      | "tonal-tertiary"
      | "tonal-surface";
    triggerSize?: "default" | "sm";
    triggerIcon?: boolean;
    triggerSpaced?: boolean;
    showLabel?: boolean;
  }

  const {
    announcement = null,
    club = null,
    team = null,
    triggerVariant = "filled-primary",
    triggerSize = "default",
    triggerIcon = false,
    triggerSpaced = false,
    showLabel = true,
  }: Props = $props();

  function formFromProps(data: ExpandedAnnouncement | null) {
    return (
      data ?? {
        title: "",
        bodytext: "",
        link: "",
        link_text: "",
        author: authRecord?.id,
        club: club?.id,
        team: team?.id,
      }
    );
  }

  let form: Partial<ExpandedAnnouncement> = $derived.by(() => {
    const formData = $state(formFromProps(announcement));
    return formData;
  });

  let open = $state(false);

  const isEditing = $derived(announcement !== null);

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    if (!form.priority) {
      form.priority = "info";
    }

    let result: AnnouncementsResponse | null = null;

    try {
      if (form.id) {
        result = await client
          .collection("announcements")
          .update<AnnouncementsResponse>(form.id, form);
      } else {
        result = await client
          .collection("announcements")
          .create<AnnouncementsResponse>(form);
      }
    } catch {
      toastController.triggerGenericFormErrorMessage("Announcement");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Announcement");
      open = false;
    }
    await invalidateAll();
  }
</script>

<Sheet.Root bind:open>
  <Sheet.Trigger
    class={[
      "btn",
      "announcement-form-trigger",
      "trigger-button",
      `trigger-variant-${triggerVariant}`,
      triggerSize === "sm" && "btn-sm",
      triggerIcon && "btn-icon",
      triggerSpaced && "trigger-spaced",
      triggerVariant === "filled-primary" && "preset-filled-primary-500",
      triggerVariant === "tonal-primary" &&
        "preset-tonal-primary border-primary",
      triggerVariant === "tonal-secondary" &&
        "preset-tonal-secondary border-secondary",
      triggerVariant === "tonal-tertiary" &&
        "preset-tonal-tertiary border-tertiary",
      triggerVariant === "tonal-surface" &&
        "preset-tonal-surface border-surface",
    ]}
    data-testid="announcement-form-trigger-{isEditing ? 'edit' : 'create'}"
  >
    {#if form.id}
      <SquarePen />
      {#if showLabel}
        <span>Edit Announcement</span>
      {/if}
    {:else}
      <Plus />
      {#if showLabel}
        <span>Create new</span>
      {/if}
    {/if}
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header>
      {#if form.id}
        <h2 class="h3">Edit Announcement "{form?.title}"</h2>
      {:else}
        <h2 class="h3">Create new Announcement</h2>
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

        <input
          autocomplete="off"
          class="input"
          name="author"
          readonly
          type="hidden"
          value={form.author}
        />

        {#if club}
          <input
            autocomplete="off"
            value={club.id}
            class="input"
            name="club"
            readonly
            type="hidden"
          />
        {/if}

        {#if team}
          <input
            autocomplete="off"
            value={team.id}
            class="input"
            name="team"
            readonly
            type="hidden"
          />
        {/if}

        <label class="label field-wide">
          <span>Title</span>
          <input
            bind:value={form.title}
            class="input"
            name="title"
            required
            type="text"
          />
        </label>

        <label class="label field-wide">
          <span>Announcement Text</span>
          <textarea
            bind:value={form.bodytext}
            class="textarea"
            name="desc"
            required
            rows="10"
          ></textarea>
        </label>

        <fieldset
          class="field-wide border border-surface-200-800 p-3 rounded-base"
        >
          <legend class="legend mb-3">Priority</legend>
          {#each ["info", "warning", "danger"] as prio}
            <label
              class="label priority-radio-label"
            >
              <input
                class="radio"
                type="radio"
                name="priority"
                value={prio}
                required
                checked={prio === "info"}
                bind:group={form.priority}
              />
              {prio}
            </label>
          {/each}
        </fieldset>

        <fieldset
          class="field-wide border border-surface-200-800 rounded-base"
        >
          <legend class="legend h6">Link Settings</legend>

          <div class=" link-grid">
            <label class="label">
              <span>Link</span>
              <input
                bind:value={form.link}
                class="input"
                name="link"
                pattern="https?://.+"
                placeholder="https://example.com"
                title="Please enter a valid URL"
                type="url"
              />
              <span class="text-sm"
                >Single link in case the announcement is used as a call to
                action.</span
              >
            </label>

            <label class="label">
              <span>Link Text</span>
              <input
                bind:value={form.link_text}
                class="input"
                name="link_text"
                placeholder="Click here"
                type="text"
              />
              <span class="text-sm"
                >If not set, the link itself will be used as a the text.</span
              >
            </label>
          </div>
        </fieldset>
      </div>

      <hr />

      <div class="submit-box">
        <button class="btn preset-filled-primary-500" type="submit">
          Submit
        </button>
      </div>
    </form>
  </Sheet.Content>
</Sheet.Root>

<style>
  :global(.announcement-form-trigger .lucide-icon) {
    flex-shrink: 1;
  }
  
  .legend {
    margin-block: calc(var(--spacing) * 2);
  }
  
  .priority-radio-label {
    text-transform: capitalize;
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 2);
    margin-block: var(--spacing);
  }

  hr {
    margin-block: calc(var(--spacing) * 5);
  }

  .submit-box {
    display: flex;
    justify-content: center;
  }

  .edit-form-grid {
    gap: calc(var(--spacing) * 2);
  }
  
  .text-sm {
    font-size: var(--text-sm);
  }
  
  .link-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 2);
    margin-block: var(--spacing);
    
    @media (min-width: 32rem) {
      grid-template-columns: repeat(2, 1fr);
    }
  }
</style>
