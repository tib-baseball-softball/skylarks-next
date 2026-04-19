<script lang="ts">
  // @ts-ignore
  // noinspection ES6UnusedImports
  import {Tabs} from "bits-ui";
  import {Users} from "lucide-svelte";
  import {goto} from "$app/navigation";
  import AnnouncementSectionContent from "$lib/dp/components/announcements/AnnouncementSectionContent.svelte";
  import TeamAdminSection from "$lib/dp/components/team/TeamAdminSection.svelte";
  import TeamTeaserCard from "$lib/dp/components/team/TeamTeaserCard.svelte";
  import AnnouncementForm from "$lib/dp/components/forms/AnnouncementForm.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, EventType,} from "$lib/dp/types/ExpandedResponse.js";
  import Paginator from "$lib/dp/utility/Paginator.svelte";
  import type {PageProps} from "./$types";
  import JoinTeamSection from "$lib/dp/components/team/JoinTeamSection.svelte";
  import ICalSection from "$lib/dp/components/settings/ICalSection.svelte";
  import EventGrid from "$lib/dp/components/event/EventGrid.svelte";

  const {data}: PageProps = $props();
  const events = $derived(data.events);
  const currentPage = $derived($events.page);
  const announcementStore = $derived(data.announcementStore);
  const model = $derived(authSettings.record) as CustomAuthModel;

  let showEvents = $state("next");
  let sorting: "asc" | "desc" | string = $state("asc");
  let showTypes: EventType | "any" | string = $state("any");

  const reloadWithQuery = () => {
    const queryString = `?timeframe=${showEvents}&page=${currentPage}&sort=${sorting}&type=${showTypes}`;

    goto(queryString, {
      noScroll: true,
    });
  };

  $effect.pre(() => {
    console.log(showEvents);
    reloadWithQuery();
  });

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canEdit = $derived(
    data.team?.admins.includes(authRecord?.id) ||
    data.team?.expand?.club?.admins.includes(authRecord?.id),
  );
  const isMember = $derived(authRecord?.teams.includes(data.team.id));
</script>

<h1 class="h1 page-title">{data.team.name} ({data.team?.expand?.club.name})</h1>

<div class="team-overview-grid">
  <article class="card team-description-card preset-tonal-surface">
    <header class="card-header">
      <h2 class="h4 card-title">Team Description</h2>
    </header>
    <section class="card-content">{@html data.team.description}</section>
  </article>

  <TeamTeaserCard link={false} team={data.team}/>
</div>

{#if isMember || canEdit}
  <section class="announcements-section">
    <header>
      <h2 class="h2">Announcements</h2>
    </header>

    <AnnouncementSectionContent store={announcementStore}/>

    {#if canEdit}
      <div class="announcement-actions">
        <AnnouncementForm
          announcement={null}
          team={data.team}
          club={null}
          triggerVariant="filled-primary"
          showLabel={true}
        />
      </div>
    {/if}
  </section>

  <section class="events-section">
    <header>
      <h2 class="h2">Team Events</h2>
    </header>

    <div class="filters-bar preset-tonal-surface">
      <label class="filter-label">
        <span>Timeframe</span>
        <Tabs.Root bind:value={showEvents}>
          <Tabs.List class="tabs-list event-segment-container">
            <Tabs.Trigger
              class="tabs-trigger btn timeframe-trigger-next"
              data-testid="segment-item"
              value="next"
            >Next
            </Tabs.Trigger>
            <Tabs.Trigger
              class="tabs-trigger btn timeframe-trigger-past"
              data-testid="segment-item"
              value="past"
            >Past
            </Tabs.Trigger>
          </Tabs.List>
        </Tabs.Root>
      </label>

      <label class="filter-label">
        <span>Sort</span>
        <Tabs.Root bind:value={sorting}>
          <Tabs.List class="tabs-list event-segment-container">
            <Tabs.Trigger
              class="tabs-trigger btn"
              data-testid="segment-item"
              value="asc">Ascending
            </Tabs.Trigger
            >
            <Tabs.Trigger
              class="tabs-trigger btn"
              data-testid="segment-item"
              value="desc">Descending
            </Tabs.Trigger
            >
          </Tabs.List>
        </Tabs.Root>
      </label>

      <label class="filter-label">
        <span>Type</span>
        <Tabs.Root bind:value={showTypes}>
          <Tabs.List class="tabs-list event-segment-container type-tabs">
            <Tabs.Trigger
              class="tabs-trigger btn"
              data-testid="segment-item"
              value="any">All
            </Tabs.Trigger
            >
            <Tabs.Trigger
              class="tabs-trigger btn"
              data-testid="segment-item"
              value="game">Game
            </Tabs.Trigger
            >
            <Tabs.Trigger
              class="tabs-trigger btn"
              data-testid="segment-item"
              value="practice">Practice
            </Tabs.Trigger
            >
            <Tabs.Trigger
              class="tabs-trigger btn"
              data-testid="segment-item"
              value="misc">Other
            </Tabs.Trigger
            >
          </Tabs.List>
        </Tabs.Root>
      </label>
    </div>

    {#if !isMember}
      <p class="hint">Only team members can participate in events.</p>
    {/if}

    <EventGrid events={$events?.items}/>

  </section>

  <Paginator showIfSinglePage={false} store={events}/>

  <hr class="section-divider"/>

  <section class="links-section">
    <header>
      <h2 class="h3">Links</h2>
    </header>

    <div class="links-container">
      <a
        class="btn preset-tonal-tertiary border-tertiary"
        href="/account/team/{data.team.id}/members"
      >
        <Users/>
        <span>Player List</span>
      </a>
    </div>
  </section>

  <hr class="section-divider"/>

  <article class="cal-card card preset-outlined-surface-500">
    <ICalSection link={`${model?.ical_link}?team=${data.team.id}`}>
      {#snippet header()}
        <span>Calendar Import</span>
      {/snippet}

      {#snippet subheader()}
        <p>
          This link includes all events for this team,
          going back one year.
        </p>
      {/snippet}
    </ICalSection>
  </article>

  {#if canEdit}
    <hr class="admin-divider"/>

    <TeamAdminSection team={data.team} eventSeries={data.eventSeries}/>
  {/if}
{:else}
  <JoinTeamSection {authRecord} teamID={data.team.id}/>
{/if}

<style>
  .page-title {
    margin-block: calc(var(--spacing) * 3) !important;
  }

  .team-overview-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 3);
    margin-block: calc(var(--spacing) * 6) !important;

    @media (min-width: 48rem) {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }

    @media (min-width: 64rem) {
      grid-template-columns: repeat(3, minmax(0, 1fr));
    }
  }

  .team-description-card {
    @media (min-width: 64rem) {
      grid-column: span 2 / span 2;
    }
  }

  .card-title {
    font-weight: 500;
  }

  .card-content {
    padding: calc(var(--spacing) * 4);
  }

  .announcements-section {
    margin-block: calc(var(--spacing) * 8) !important;
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 4);

    h2 {
      margin-bottom: calc(var(--spacing) * 3);
    }
  }

  .filters-bar {
    display: flex;
    flex-wrap: wrap;
    gap: calc(var(--spacing) * 4);
    justify-content: space-between;
    padding-inline: calc(var(--spacing) * 4);
    padding-block: calc(var(--spacing) * 3);
    border-radius: var(--radius-base);
    font-size: var(--text-sm);

    @media (min-width: 64rem) {
      font-size: var(--text-base);
    }
  }

  .cal-card {
    margin-block-end: calc(var(--spacing) * 6);
  }

  .section-divider {
    margin-block: calc(var(--spacing) * 8) !important;
  }

  .links-section {
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 2);

    @media (min-width: 64rem) {
      gap: calc(var(--spacing) * 4);
    }
  }

  .links-container {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: calc(var(--spacing) * 2);

    @media (min-width: 64rem) {
      gap: calc(var(--spacing) * 3);
    }
  }

  .admin-divider {
    margin-block: calc(var(--spacing) * 2);
  }

  :global {
    .event-segment-container {
      border: 1px solid var(--color-surface-600-400);
      padding: calc(var(--spacing) * 1) !important;

      .tabs-trigger {
        padding: 0.25rem 0.6rem;

        &.timeframe-trigger-next:active {
          background-color: var(--color-error-300-700);
          color: var(--color-error-contrast-300-700);
        }

        &.timeframe-trigger-past:active {
          background-color: var(--color-surface-950-50);
          color: var(--color-surface-50-950);
        }
      }

      &.type-tabs {
        gap: calc(var(--spacing) * 1);
      }
    }

    .border-tertiary {
      border: 1px solid var(--color-tertiary-500) !important;
    }
  }

  .hint {
    margin-block: calc(var(--spacing) * 4);
  }

  .filter-label {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 2);
    justify-content: space-between;
    flex-grow: 1;

    @media (min-width: 48rem) {
      flex-grow: 0;
    }

    @media (min-width: 80rem) {
      justify-content: flex-start;
    }
  }

  header {
    margin-block: calc(var(--spacing) * 3);
  }

  section {
    margin-block: calc(var(--spacing) * 4);
  }
</style>
