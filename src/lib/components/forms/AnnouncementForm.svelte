<script lang="ts">
  import {invalidateAll} from "$app/navigation";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import type {AnnouncementsResponse, ClubsResponse, TeamsResponse} from "$lib/model/pb-types";
  import type {CustomAuthModel, ExpandedAnnouncement} from "$lib/model/ExpandedResponse";
  import {Edit, Plus} from "lucide-svelte";
  //@ts-ignore
  import * as Sheet from "$lib/components/utility/sheet/index.js";
  import {toastController} from "$lib/service/ToastController.svelte.ts";


  const authRecord = $derived(authSettings.record as CustomAuthModel);

  interface Props {
    announcement: ExpandedAnnouncement | null;
    club: ClubsResponse | null;
    team: TeamsResponse | null;
    buttonClasses?: string;
    showLabel?: boolean,
  }

  let {announcement = null, club = null, team = null, buttonClasses = "", showLabel = true}: Props = $props();

  const form: Partial<ExpandedAnnouncement> = $state(
      announcement ?? {
        title: "",
        bodytext: "",
        link: "",
        link_text: "",
        author: authRecord?.id,
      },
  );

  let open = $state(false);

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: AnnouncementsResponse | null = null;

    try {
      if (form.id) {
        result = await client
            .collection("clubs")
            .update<AnnouncementsResponse>(form.id, form);
      } else {
        result = await client
            .collection("clubs")
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

<Sheet.Root bind:open={open}>
  <Sheet.Trigger class={buttonClasses}>
    {#if form.id}
      <Edit/>
      {#if showLabel}
        <span>Edit Announcement</span>
      {/if}
    {:else}
      <Plus/>
      {#if showLabel}
        <span>Create Announcement</span>
      {/if}
    {/if}
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header class="text-xl font-semibold">
      {#if form.id}
        <h2 class="h3">Edit Announcement "{form?.title}"</h2>
      {:else}
        <h2 class="h3">Create new Announcement</h2>
      {/if}
    </header>

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


      </div>

      <hr class="my-5!"/>

      <div class="flex justify-center gap-3">
        <button class="mt-2 btn preset-filled-primary-500" type="submit">
          Submit
        </button>
      </div>
    </form>

  </Sheet.Content>
</Sheet.Root>
