<script lang="ts">
  import {Calendar, Clock, User} from "lucide-svelte";
  import PriorityBadge from "$lib/dp/components/announcements/PriorityBadge.svelte";
  import CommentSection from "$lib/dp/components/comments/CommentSection.svelte";
  import AnnouncementForm from "$lib/dp/components/forms/AnnouncementForm.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import type {PageProps} from "./$types";

  const {data}: PageProps = $props();
  const announcement = $derived(data.announcement);

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canEdit = $derived(
    $announcement.expand?.club?.admins.includes(authRecord.id) ||
    $announcement.expand?.team?.admins.includes(authRecord.id)
  );

  function deleteAction(id: string) {
    client.collection("announcements").delete(id);
    history.back();
  }

  const updated = $derived($announcement.updated);
  //@ts-expect-error - the multi-level expanding trips the typedef up
  const club = $derived($announcement?.expand?.club ?? $announcement?.expand?.team?.expand?.club);
</script>

<div class="announcement-header">
  <h1 class="h1 announcement-title">{$announcement.title}</h1>
  <PriorityBadge priority={$announcement.priority}/>
</div>

<div class="announcement-layout">
  <div class="announcement-container">
    <section class="meta-section">
      <dl class="meta-info">
        {#if $announcement.expand?.team}
          <dt>Team:</dt>
          <dd class="badge preset-tonal-primary border-primary">
            {$announcement.expand?.team?.name}
          </dd>
        {/if}

        {#if $announcement.expand?.club}
          <dt>Club:</dt>
          <dd class="badge preset-tonal-primary border-primary">
            {$announcement.expand?.club?.name}
          </dd>
        {/if}
      </dl>

      <p class="author-info">
        <User size="28"/>
        <span class="h6">{$announcement.expand?.author?.first_name} {$announcement.expand?.author?.last_name}</span>
      </p>

      <p class="date-info">
        <Calendar size="28"/>
        <time class="h6" datetime="{updated}">{new Date(updated).toLocaleDateString()}</time>
      </p>

      <p class="time-info">
        <Clock size="28"/>
        <time class="h6" datetime="{updated}">{new Date(updated).toLocaleTimeString()}</time>
      </p>
    </section>

    <section class="prose announcement-body">
      {@html $announcement.bodytext}
    </section>

    {#if $announcement.link}
      <section class="link-section">
        <header>
          <h2 class="h4">Link</h2>
        </header>

        <a class="anchor announcement-link" href={$announcement.link} target="_blank"
           title="Link associated with this announcement: {$announcement.link}">
          {#if $announcement.link_text}
            {$announcement.link_text}
          {:else }
            {$announcement.link}
          {/if}
        </a>
      </section>
    {/if}

    <section class="comments-section">
      <header>
        <h2 class="h4">Comments</h2>
      </header>
      <div class="comments-wrapper">
        <CommentSection
          club={club}
          comments={$announcement?.expand?.comments_via_announcement ?? []}
          targetID={$announcement.id}
          targetType="announcement"
        />
      </div>
    </section>

    {#if canEdit}
      <hr class="divider"/>

      <section class="admin-section card">

        <header>
          <h2 class="h4">Admin Options</h2>
        </header>
        <div class="admin-actions">
          <AnnouncementForm showLabel={true} announcement={$announcement} club={$announcement.expand?.club ?? null}
                            team={$announcement.expand?.team ?? null}
                            buttonClasses="btn preset-tonal-tertiary border-tertiary"/>

          <DeleteButton
            id={$announcement.id}
            modelName="Announcement"
            action={deleteAction}
            buttonText="Delete Announcement"
            classes="btn preset-tonal-error border-error"
            data-testid="delete-announcement-button"
          />
        </div>
      </section>
    {/if}
  </div>
</div>

<style>
  .announcement-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: calc(var(--spacing) * 4);

    @media (min-width: 64rem) {
      gap: calc(var(--spacing) * 6);
    }
  }

  .announcement-title {
    margin-block: calc(var(--spacing) * 3) !important;
  }

  .announcement-layout {
    display: flex;
    justify-content: center;
  }

  .announcement-container {
    width: 100%;
    max-width: 65ch; /* Standard prose width for readability */
  }

  .meta-section {
    margin-block: calc(var(--spacing) * 6);
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 3);
  }

  .meta-info {
    display: flex;
    gap: calc(var(--spacing) * 2);
    align-items: center;
  }

  .author-info, .date-info, .time-info {
    display: flex;
    gap: calc(var(--spacing) * 3);
    align-items: center;
  }

  .announcement-body {
    text-align: justify;
  }

  .link-section, .comments-section {
    margin-block: calc(var(--spacing) * 6);
  }

  .announcement-link {
    display: block;
    margin-top: calc(var(--spacing) * 2);
  }

  .comments-wrapper {
    margin-top: calc(var(--spacing) * 4);
    padding: calc(var(--spacing) * 3);
    border: 1px solid var(--color-surface-200);
    border-radius: var(--radius-base);

    @media (min-width: 48rem) {
      padding: calc(var(--spacing) * 4);
    }
    
    :global([data-theme='dark']) & {
        border-color: var(--color-surface-100);
    }
  }

  .divider {
    margin-block: calc(var(--spacing) * 6);
  }

  .admin-section {
    margin-block: calc(var(--spacing) * 8);
    padding: calc(var(--spacing) * 4);
    border: 1px solid var(--color-surface-200);

    :global([data-theme='dark']) & {
        border-color: var(--color-surface-100);
    }
  }

  .admin-actions {
    display: flex;
    gap: calc(var(--spacing) * 3);
    margin-top: calc(var(--spacing) * 3);
    align-items: center;
    flex-wrap: wrap;
  }

  :global(.border-primary) {
      border: 1px solid var(--color-primary-500) !important;
  }

  :global(.border-tertiary) {
      border: 1px solid var(--color-tertiary-500) !important;
  }

  :global(.border-error) {
      border: 1px solid var(--color-error-500) !important;
  }
</style>
