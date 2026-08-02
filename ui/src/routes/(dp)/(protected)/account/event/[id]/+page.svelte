<script lang="ts">
  import type { Match } from "bsm.js";
  import { Ban, Clock } from "@lucide/svelte";
  import CommentSection from "$lib/dp/components/comments/CommentSection.svelte";
  import EventAttireSection from "$lib/dp/components/event/EventAttireSection.svelte";
  import EventCoreInfo from "$lib/dp/components/event/EventCoreInfo.svelte";
  import EventPageAdminSection from "$lib/dp/components/event/EventPageAdminSection.svelte";
  import EventParticipantsOverviewSection from "$lib/dp/components/event/EventParticipantsOverviewSection.svelte";
  import EventParticipationSection from "$lib/dp/components/event/EventParticipationSection.svelte";
  import EventTypeBadge from "$lib/dp/components/event/EventTypeBadge.svelte";
  import TimeSection from "$lib/dp/components/event/TimeSection.svelte";
  import MatchDetailLocationCard from "$lib/dp/components/event/match/MatchDetailLocationCard.svelte";
  import MatchTeaserCard from "$lib/dp/components/event/match/MatchTeaserCard.svelte";
  import { authSettings } from "$lib/dp/client.svelte.js";
  import type { CustomAuthModel } from "$lib/dp/types/ExpandedResponse.ts";
  import type { ClubsResponse, TeamsResponse } from "$lib/dp/types/pb-types.ts";
  import EventTeamBadges from "$lib/dp/components/event/EventTeamBadges.svelte";
  import type { PageProps } from "./$types";

  const { data }: PageProps = $props();

  const event = $derived(data.event);

  const isClubWide = $derived(
    typeof $event.club === "string" && $event.club !== "",
  );

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canParticipate = $derived.by(() => {
    if (isClubWide) {
      return authRecord.club.includes($event.club);
    }
    const allApplicableTeams = new Set<string>($event.additional_teams);
    allApplicableTeams.add($event.team);

    return authRecord.teams.some((team) => allApplicableTeams.has(team));
  });

  //@ts-expect-error - the multi-level expanding trips the typedef up
  const club = $derived($event?.expand?.club) as ClubsResponse;

  const matchJSON = $derived($event?.match_json) as unknown as Match;

  /**
   * @TODO consider moving this expensive logic to backend
   */
  const canEdit = $derived.by(() => {
    if ($event.team) {
      let allPossibleAdminIDs = new Set<string>($event.expand?.team?.admins);

      for (const admin of $event?.expand?.team?.expand?.club?.admins) {
        allPossibleAdminIDs.add(admin);
      }

      for (const admin of $event?.expand?.additional_teams?.flatMap(
        (team: TeamsResponse) => team.admins,
      )) {
        allPossibleAdminIDs.add(admin);
      }
      return allPossibleAdminIDs.has(authRecord.id);
    }
    return $event?.expand?.club?.admins.includes(authRecord.id);
  });
</script>

<svelte:head>
  <title>{$event.title}</title>
  <meta
    content="Event details for {$event.title}, including time, type, and location."
    name="description"
  />
</svelte:head>

<div class="event-page-container">
  <div>
    <div class="header-row">
      <h1 class={["h1", $event.cancelled && "cancelled-text"]}>
        {$event.title}
      </h1>
      <div class="type-badge-wrapper">
        <EventTypeBadge type={$event.type} />
      </div>
    </div>

    <EventTeamBadges event={$event} {isClubWide} />
  </div>

  {#if $event.cancelled}
    <span class="badge cancelled-badge">
      <Ban />
      Cancelled
    </span>
  {/if}

  <article class="description-section" class:cancelled-text={$event.cancelled}>
    <section>
      <p>{$event.desc}</p>
    </section>
  </article>

  <div class="core-info-section" class:cancelled-text={$event.cancelled}>
    <EventCoreInfo event={$event}>
      {#snippet additionalTimeSection()}
        <TimeSection
          timeValue={$event.endtime}
          displayText="End:"
          classes="col-span-2"
        >
          {#snippet icon()}
            <Clock size="18" />
          {/snippet}
        </TimeSection>
      {/snippet}
    </EventCoreInfo>
  </div>

  {#if !$event.cancelled}
    <div class="participation-header">
      <h2 class="h4">My Participation</h2>

      {#if canParticipate}
        <EventParticipationSection
          event={$event}
          growChips={true}
          {canParticipate}
        />
      {:else}
        <div class="participation-info">
          <p>Only team members can participate in events.</p>
        </div>
      {/if}
    </div>

    <hr class="divider" />

    <EventParticipantsOverviewSection event={$event} />
  {/if}

  <hr class="divider" />
  <div class="details-grid">
    {#if $event.expand.location}
      <section class="details-section">
        <h2 class="h2 details-title">Location Details</h2>

        <MatchDetailLocationCard
          --location-padding="0"
          --location-spacing="4"
          field={$event.expand.location}
          showDividers={false}
        />
      </section>
    {/if}

    {#if $event.expand.attire}
      <section class="details-section">
        <EventAttireSection attire={$event.expand.attire} />
      </section>
    {/if}
  </div>

  <hr class="divider" />

  {#if $event.match_json}
    <section class="game-data-section">
      <h2 class="h2 game-data-title">Game Data</h2>
      <div class="game-data-grid">
        <MatchTeaserCard
          match={matchJSON}
          teamName={$event?.expand?.team?.name}
        />
      </div>
    </section>
  {/if}

  <section class="comments-section">
    <div class="comments-wrapper preset-outlined-card card">
      <header>
        <h2 class="h3">Comments</h2>
      </header>
      <CommentSection
        {club}
        comments={$event?.expand?.comments_via_event ?? []}
        targetID={$event.id}
        targetType="event"
      />
    </div>
  </section>

  {#if canEdit}
    <EventPageAdminSection event={$event} />
  {/if}
</div>

<style>
  .event-page-container {
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 4);

    @media (min-width: 64rem) {
      gap: calc(var(--spacing) * 6);
    }

    @media (min-width: 80rem) {
      gap: calc(var(--spacing) * 7);
    }
  }

  .header-row {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 3);
  }

  .cancelled-text {
    text-decoration: line-through;
  }

  .cancelled-badge {
    background-color: var(--color-error-500);
    color: var(--color-error-contrast-500);
    gap: calc(var(--spacing) * 1);
    width: fit-content;
  }

  .description-section {
    margin-bottom: calc(var(--spacing) * 2);
  }

  .details-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 6);

    @media (min-width: 48rem) {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }
  }

  .details-section {
    margin-top: calc(var(--spacing) * 2);

    @media (min-width: 64rem) {
      margin-top: calc(var(--spacing) * 4);
    }
  }

  .details-title {
    margin-bottom: calc(var(--spacing) * 3);
  }

  .divider {
    margin-block: calc(var(--spacing) * 4);
  }

  .participation-header {
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 2.5);

    @media (min-width: 48rem) {
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
      gap: 0;
    }
  }

  .participation-info {
    display: flex;
    justify-content: flex-end;
  }

  .game-data-title {
    margin-bottom: calc(var(--spacing) * 3);
  }

  .game-data-grid {
    display: grid;
    grid-template-columns: 1fr;

    @media (min-width: 48rem) {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }
  }

  .comments-section {
    margin-block: calc(var(--spacing) * 4);
  }

  .comments-wrapper {
    margin-top: calc(var(--spacing) * 4);
    padding: calc(var(--spacing) * 3);
    max-width: 65ch;
    margin-inline: auto;

    @media (min-width: 48rem) {
      padding: calc(var(--spacing) * 4);
    }
  }
</style>
