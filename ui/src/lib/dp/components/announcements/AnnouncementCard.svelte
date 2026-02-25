<script lang="ts">
  import {MessageCircle} from "lucide-svelte";
  import AnnouncementForm from "$lib/dp/components/forms/AnnouncementForm.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.ts";
  import type {CustomAuthModel, ExpandedAnnouncement} from "$lib/dp/types/ExpandedResponse.ts";
  import AnnouncementCoreContent from "./AnnouncementCoreContent.svelte";

  interface Props {
    announcement: ExpandedAnnouncement;
  }

  const {announcement}: Props = $props();
  const commentCount = $derived(announcement?.expand?.comments_via_announcement?.length ?? 0);

  function deleteAction(id: string) {
    client.collection("announcements").delete(id);
  }

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canEdit = $derived(
    announcement.expand?.club?.admins.includes(authRecord.id) ||
    announcement.expand?.team?.admins.includes(authRecord.id)
  );
</script>

<article class="root card preset-tonal-surface shadow-xl">
  <div class="inner-container">
    <div>
      <AnnouncementCoreContent {announcement} textClasses="line-clamp-5"/>
    </div>

    <div>
      <hr class="separator"/>

      <footer class="actions">

        {#if canEdit}
          <div class="admin-actions">
            <AnnouncementForm showLabel={false} announcement={announcement} club={announcement.expand?.club ?? null}
                              team={announcement.expand?.team ?? null}
                              triggerVariant="tonal-tertiary"
                              triggerSize="sm"
                              triggerIcon={true}
            />

            <DeleteButton
              id={announcement.id}
              modelName="Announcement"
              action={deleteAction}
              classes="btn btn-icon btn-sm preset-tonal-error border border-error-500"
              data-testid="delete-announcement-button"
              iconSize={16}
            />
          </div>
        {:else }
          <div aria-hidden="true"></div>
        {/if}

        <div class="meta-actions">
          {#if commentCount > 0}
            <div class="badge-wrapper">
              <span class="sr-only">Number of Comments:</span>
              <span class="badge-icon comment-badge preset-filled-primary-500">
                {commentCount}
              </span>
              <MessageCircle aria-hidden="true"/>
            </div>
          {/if}

          <a
            class="btn btn-sm preset-filled-primary-500"
            href="/account/announcements/{announcement.id}">Read more</a
          >
        </div>
      </footer>
    </div>
  </div>
</article>

<style>
  .root {
    padding: calc(var(--spacing) * 4);
    height: 100%;
  }

  .inner-container {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 100%;
  }

  .separator {
    margin-top: calc(var(--spacing) * 6);
    margin-bottom: calc(var(--spacing) * 2);
  }

  .actions {
    margin-top: calc(var(--spacing) * 3);
    display: flex;
    gap: calc(var(--spacing) * 2);
    justify-content: space-between;
  }

  .admin-actions {
    display: flex;
    gap: calc(var(--spacing) * 2);
  }

  .meta-actions {
    display: flex;
    gap: calc(var(--spacing) * 4);
  }

  .badge-icon {
    top: -9px;
    right: -6px;
    padding: 0.2rem;
  }

  .badge-wrapper {
    position: relative;
    display: inline-block;
  }

  .comment-badge {
    position: absolute;
    z-index: 10;
  }
</style>
