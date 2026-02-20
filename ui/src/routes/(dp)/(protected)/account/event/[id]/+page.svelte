<script lang="ts">
  import type {Match} from "bsm.js";
  import {Ban, Clock} from "lucide-svelte";
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
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import type {ClubsResponse} from "$lib/dp/types/pb-types.ts";

  const {data} = $props();

  const event = $derived(data.event);

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canParticipate = $derived(authRecord.teams.includes($event.team));

  //@ts-expect-error - the multi-level expanding trips the typedef up
  const club = $derived($event?.expand?.club) as ClubsResponse;

  const matchJSON = $derived($event?.match_json) as unknown as Match;
</script>

<div class="event-page-container">
  <div class="header-row">
    <h1 class="h1" class:cancelled-text={$event.cancelled}>{$event.title}</h1>
    <div class="type-badge-wrapper">
      <EventTypeBadge type={$event.type}/>
    </div>
  </div>

  {#if $event.cancelled}
    <span class="badge cancelled-badge">
      <Ban/>
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
        <TimeSection timeValue={$event.endtime} displayText="End:" classes="col-span-2">
          {#snippet icon()}
            <Clock size="18"/>
          {/snippet}
        </TimeSection>
      {/snippet}
    </EventCoreInfo>
  </div>

  <div class="details-grid">
    {#if $event.expand.location}
      <section class="details-section">
        <h2 class="h2 details-title">Location Details</h2>

        <MatchDetailLocationCard
          field={$event.expand.location}
          classes="location-card-spacing"
          showDividers={false}
        />
      </section>
    {/if}

    {#if $event.expand.attire}
      <section class="details-section">
        <EventAttireSection attire={$event.expand.attire}/>
      </section>
    {/if}
  </div>

  {#if !$event.cancelled}
    <hr class="divider-large">

    <div class="participation-header">
      <h2 class="h4">My Participation</h2>

      {#if canParticipate}
        <EventParticipationSection event={$event} chipClasses="flex-grow"/>
      {:else}
        <div class="participation-info">
          <p>Only team members can participate in events.</p>
        </div>
      {/if}
    </div>

    <hr class="divider">

    <EventParticipantsOverviewSection event={$event}/>
  {/if}

  {#if $event.match_json}
    <section class="game-data-section">
      <h2 class="h2 game-data-title">Game Data</h2>
      <div class="game-data-grid">
        <MatchTeaserCard match={matchJSON} teamName={$event?.expand?.team?.name}/>
      </div>
    </section>
  {/if}

  <section class="comments-section">
    <div class="comments-wrapper">
      <header>
        <h2 class="h3">Comments</h2>
      </header>
      <CommentSection
        club={club}
        comments={$event?.expand?.comments_via_event ?? []}
        targetID={$event.id}
        targetType="event"
      />
    </div>
  </section>

  {#if $event.expand?.team?.admins.includes(authRecord.id) || $event?.expand?.team?.expand?.club?.admins.includes(authRecord.id)}
    <EventPageAdminSection event={$event}/>
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
        margin-bottom: calc(var(--spacing) * 8);
    }

    .details-grid {
        display: grid;
        grid-template-columns: 1fr;
        
        @media (min-width: 48rem) {
            grid-template-columns: repeat(2, minmax(0, 1fr));
            gap: calc(var(--spacing) * 4);
        }
    }

    .details-section {
        margin-top: calc(var(--spacing) * 3);
        
        @media (min-width: 64rem) {
            margin-top: calc(var(--spacing) * 5);
        }
    }

    .details-title {
        margin-bottom: calc(var(--spacing) * 3);
    }

    /* Handle the location card spacing when used as an override prop */
    :global(.location-card-spacing) {
        margin-block: calc(var(--spacing) * 4) !important;
        gap: calc(var(--spacing) * 5) !important;
    }

    .divider-large {
        margin-block: calc(var(--spacing) * 8);
    }

    .divider {
        margin-block: calc(var(--spacing) * 8);
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

    .game-data-section {
        /* any specific section styles */
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
        margin-block: calc(var(--spacing) * 6);
    }

    .comments-wrapper {
        margin-top: calc(var(--spacing) * 4);
        padding: calc(var(--spacing) * 3);
        border: 1px solid var(--color-surface-900-100);
        border-radius: var(--radius-base);
        max-width: 65ch;
        margin-inline: auto;
        
        @media (min-width: 48rem) {
            padding: calc(var(--spacing) * 4);
        }
    }
</style>
