<script lang="ts">
  import { Ban } from "@lucide/svelte";
  import EventTypeBadge from "$lib/dp/components/event/EventTypeBadge.svelte";
  import type {
    CustomAuthModel,
    ExpandedEvent,
  } from "$lib/dp/types/ExpandedResponse.ts";
  import EventCoreInfo from "$lib/dp/components/event/EventCoreInfo.svelte";
  import EventParticipationSection from "./EventParticipationSection.svelte";
  import EventTeamBadges from "$lib/dp/components/event/EventTeamBadges.svelte";
  import { authSettings } from "$lib/dp/client.svelte";

  interface props {
    event: ExpandedEvent;
    link: boolean;
  }

  const { event, link }: props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canParticipate = $derived.by(() => {
    const allApplicableTeams = new Set<string>(event.additional_teams);
    allApplicableTeams.add(event.team);

    return authRecord.teams.some((team) => allApplicableTeams.has(team));
  });
</script>

<article class={[link && "card-hover", "card preset-outlined-card teaser"]}>
  <a class={[event.cancelled && "cancelled"]} href="/account/event/{event.id}">
    <header class="card-header">
      <h2 class="title">
        <EventTypeBadge type={event.type} />
        <span class="title-text h5">{event?.title}</span>
      </h2>

      <EventTeamBadges {event} />
    </header>

    <div class="core-info-wrapper">
      <EventCoreInfo {event} />
    </div>
  </a>

  <hr />
  <footer class="card-footer">
    {#if event.cancelled}
      <div class="cancelled-container">
        <span class="badge preset-filled-error-500 cancelled-badge">
          <Ban size="18" />
          Cancelled
        </span>
      </div>
    {:else}
      <EventParticipationSection {event} {canParticipate} />
    {/if}
  </footer>
</article>

<style>
  article {
    display: grid;
    grid-template-rows: subgrid;
    grid-row: span 3;
  }

  .teaser {
    font-size: var(--text-sm);
    height: 100%;
  }

  .title-text {
    margin-inline-start: calc(var(--spacing) * 1);
    font-weight: var(--font-weight-bold);
    vertical-align: sub;
  }

  .cancelled {
    .title-text,
    .core-info-wrapper {
      text-decoration: line-through;
    }
  }

  .cancelled-container {
    display: flex;
    justify-content: end;
  }

  .cancelled-badge {
    gap: calc(var(--spacing) * 1);
  }

  .core-info-wrapper {
    padding-inline: calc(var(--spacing) * 4);
  }

  a {
    display: flex;
    flex-direction: column;
    width: 100%;
    justify-content: space-between;
  }

  .title {
    padding-block-end: calc(var(--spacing) * 3);
  }
</style>
