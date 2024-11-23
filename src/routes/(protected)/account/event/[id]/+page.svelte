<script lang="ts">
  import EventAttireSection from "$lib/components/diamondplanner/event/EventAttireSection.svelte";
  import EventCoreInfo from "$lib/components/diamondplanner/event/EventCoreInfo.svelte";
  import EventParticipationSection from "$lib/components/diamondplanner/event/EventParticipationSection.svelte";
  import EventTypeBadge from "$lib/components/diamondplanner/event/EventTypeBadge.svelte";
  import MatchTeaserCard from "$lib/components/match/MatchTeaserCard.svelte";
  import CloseOutline from "flowbite-svelte-icons/CloseOutline.svelte";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
  import EventPageAdminSection from "$lib/components/diamondplanner/event/EventPageAdminSection.svelte";
  import EventParticipantsOverviewSection
    from "$lib/components/diamondplanner/event/EventParticipantsOverviewSection.svelte";

  let {data} = $props();

  const event = $derived(data.event);

  const authRecord = authSettings.record as CustomAuthModel
  const canParticipate = $derived(authRecord.teams.includes($event.team))
</script>

<div class="event-container">
    <div class="flex items-center gap-3">
        <h1 class="h1" class:line-through={$event.cancelled}>{$event.title}</h1>
        <div>
            <EventTypeBadge type={$event.type}/>
        </div>
    </div>

    {#if $event.cancelled}
        <span class="badge variant-filled-error">
            <CloseOutline/>
            Cancelled
        </span>
    {/if}

    <article class="!mb-8" class:line-through={$event.cancelled}>
        <section>
            <p>{$event.desc}</p>
        </section>
    </article>

    <div class="space-y-6" class:line-through={$event.cancelled}>
        <EventCoreInfo event={$event}/>
    </div>

    {#if $event.expand.attire}
        <section class="mt-3 lg:mt-5">
            <EventAttireSection attire={$event.expand.attire}/>
        </section>
    {/if}

    {#if !$event.cancelled}
        <hr class="!my-8"/>

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
