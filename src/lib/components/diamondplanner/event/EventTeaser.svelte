<script lang="ts">
    import EventTypeBadge from "$lib/components/diamondplanner/event/EventTypeBadge.svelte";
    import type {CustomAuthModel, ExpandedEvent} from "$lib/model/ExpandedResponse";
    import EventCoreInfo from "./EventCoreInfo.svelte";
    import EventParticipationSection from "./EventParticipationSection.svelte";
    import {authSettings} from "$lib/pocketbase/index.svelte";
    import {Ban} from "lucide-svelte";

    interface props {
    event: ExpandedEvent;
    link: boolean;
  }

  const {event, link}: props = $props();

  const authRecord = authSettings.record as CustomAuthModel;

  // club admins can see the team and change settings, but not participate
  const canParticipate = $derived(authRecord.teams.includes(event.team));
</script>

<article
        class="card variant-soft-surface shadow-xl text-sm h-full"
        class:card-hover={link}
>
  <a href="/account/event/{event.id}" class:line-through={event.cancelled}>
    <header class="card-header">
      <h2>
        <EventTypeBadge type={event.type}/>
        <span class="ms-1 font-bold h5">{event?.title}</span>
      </h2>
    </header>

    <EventCoreInfo {event} classes={"p-4"}/>
  </a>

  <footer class="card-footer">
    <hr class="my-2"/>
    {#if event.cancelled}
      <div class="flex justify-end">
                <span class="badge variant-filled-error">
                    <Ban/>
                    Cancelled
                </span>
      </div>
    {:else}
      {#if canParticipate}
        <EventParticipationSection {event}/>
      {:else}
        <div class="flex justify-end">
          <p>Only team members can participate in events.</p>
        </div>
      {/if}
    {/if}
  </footer>
</article>
