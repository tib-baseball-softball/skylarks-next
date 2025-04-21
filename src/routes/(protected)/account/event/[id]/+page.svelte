<script lang="ts">
  import EventAttireSection from "$lib/components/diamondplanner/event/EventAttireSection.svelte";
  import EventCoreInfo from "$lib/components/diamondplanner/event/EventCoreInfo.svelte";
  import EventParticipationSection from "$lib/components/diamondplanner/event/EventParticipationSection.svelte";
  import EventTypeBadge from "$lib/components/diamondplanner/event/EventTypeBadge.svelte";
  import MatchTeaserCard from "$lib/components/match/MatchTeaserCard.svelte";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
  import EventPageAdminSection from "$lib/components/diamondplanner/event/EventPageAdminSection.svelte";
  import EventParticipantsOverviewSection
    from "$lib/components/diamondplanner/event/EventParticipantsOverviewSection.svelte";
  import {Ban, Clock} from "lucide-svelte";
  import MatchDetailLocationCard from "$lib/components/match/MatchDetailLocationCard.svelte";
  import TimeSection from "$lib/components/diamondplanner/event/TimeSection.svelte";

  let {data} = $props();

  const event = $derived(data.event);

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canParticipate = $derived(authRecord.teams.includes($event.team));
</script>

<div class="event-container">
  <div class="flex items-center gap-3">
    <h1 class="h1" class:line-through={$event.cancelled}>{$event.title}</h1>
    <div>
      <EventTypeBadge type={$event.type}/>
    </div>
  </div>

  {#if $event.cancelled}
    <span class="badge preset-filled-error-500 gap-1">
      <Ban/>
      Cancelled
    </span>
  {/if}

  <article class="mb-8!" class:line-through={$event.cancelled}>
    <section>
      <p>{$event.desc}</p>
    </section>
  </article>

  <div class="space-y-6" class:line-through={$event.cancelled}>
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

  <div class="grid grid-cols-1 md:grid-cols-2">
    {#if $event.expand.location}
      <section class="mt-3 lg:mt-5">
        <h2 class="h2">Location Details</h2>

        <MatchDetailLocationCard
                field={$event.expand.location}
                classes="my-4 space-y-5"
                showDividers={false}
        />
      </section>
    {/if}

    {#if $event.expand.attire}
      <section class="mt-3 lg:mt-5">
        <EventAttireSection attire={$event.expand.attire}/>
      </section>
    {/if}
  </div>

  {#if !$event.cancelled}
    <hr class="my-8!"/>

    <div class="flex justify-between items-center">
      <h2 class="h3">My Participation</h2>

      {#if canParticipate}
        <EventParticipationSection event={$event}/>
      {:else}
        <div class="flex justify-end">
          <p>Only team members can participate in events.</p>
        </div>
      {/if}
    </div>

    <hr class="my-8"/>

    <EventParticipantsOverviewSection event={$event}/>
  {/if}

  {#if $event.match_json}
    <section>
      <h2 class="h2 mb-3">Game Data</h2>
      <div class="grid grid-cols-1 md:grid-cols-2">
        <MatchTeaserCard match={$event.match_json}/>
      </div>
    </section>
  {/if}

  {#if $event.expand?.team?.admins.includes(authRecord.id) || $event?.expand?.team?.expand?.club?.admins.includes(authRecord.id)}
    <EventPageAdminSection event={$event}/>
  {/if}
</div>

<style lang="postcss">
    .event-container {
        @apply space-y-4 lg:space-y-6 xl:space-y-7;
    }
</style>
