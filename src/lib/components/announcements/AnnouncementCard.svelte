<script lang="ts">
  import type {CustomAuthModel, ExpandedAnnouncement,} from "$lib/model/ExpandedResponse";
  import AnnouncementCoreContent from "./AnnouncementCoreContent.svelte";
  import {authSettings, client} from "$lib/pocketbase/index.svelte.ts";
  import AnnouncementForm from "$lib/components/forms/AnnouncementForm.svelte";
  import DeleteButton from "$lib/components/utility/DeleteButton.svelte";

  interface Props {
    announcement: ExpandedAnnouncement;
  }

  let {announcement}: Props = $props();

  function deleteAction(id: string) {
    client.collection("announcements").delete(id);
  }

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canEdit = $derived(announcement.expand?.club?.admins.includes(authRecord.id) || announcement.expand?.team?.admins.includes(authRecord.id));
</script>

<article class="card p-4 preset-tonal-surface shadow-xl h-full">
  <div class="flex flex-col justify-between h-full">
    <div>
      <AnnouncementCoreContent {announcement} textClasses="line-clamp-5"/>
    </div>

    <div>
      <hr class="mt-6 mb-2"/>

      <footer class="mt-3 flex gap-2 justify-between">

        {#if canEdit}
          <div class="flex gap-2">
            <AnnouncementForm showLabel={false} announcement={announcement} club={announcement.expand?.club ?? null}
                              team={announcement.expand?.team ?? null}
                              buttonClasses="btn btn-icon btn-sm preset-tonal-tertiary border border-tertiary-500"
            />

            <DeleteButton
                    id={announcement.id}
                    modelName="Announcement"
                    action={deleteAction}
                    classes="btn btn-icon btn-sm preset-tonal-error border border-error-500"
            />
          </div>
        {:else }
          <div aria-hidden="true"></div>
        {/if}
        <a
                href="/account/announcements/{announcement.id}"
                class="btn btn-sm preset-filled-primary-500">Read more</a
        >
      </footer>
    </div>
  </div>
</article>
