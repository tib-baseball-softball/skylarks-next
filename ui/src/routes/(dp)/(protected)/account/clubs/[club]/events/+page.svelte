<script lang="ts">
  import EventFilters from "$lib/dp/components/event/EventFilters.svelte";
  import EventGrid from "$lib/dp/components/event/EventGrid.svelte";
  import EventForm from "$lib/dp/components/forms/EventForm.svelte";
  import { CalendarPlus } from "@lucide/svelte";
  import type { PageProps } from "./$types";
  import { authSettings } from "$lib/dp/client.svelte";
  import Paginator from "$lib/dp/utility/Paginator.svelte";

  const { data }: PageProps = $props();
  const eventStore = $derived(data.eventStore);
</script>

<svelte:head>
  <title>Events for {data.club.name}</title>
  <meta
    content="Club-wide events for {data.club.name} with participation status."
    name="description"
  />
</svelte:head>

<h1 class="h1 page-title">Events for {data.club.name}</h1>

<section>
  <div class="space">
    <EventFilters />
  </div>

  <div class="space">
    <EventGrid events={$eventStore?.items ?? []} />
  </div>

  <Paginator showIfSinglePage={true} store={eventStore} />

  {#if data.club?.admins.includes(authSettings.record?.id)}
    <div class="space">
      <EventForm
        mode="clubEvent"
        clubID={data.club.id}
        event={null}
        teamID=""
        triggerVariant="filled-primary"
      >
        {#snippet triggerContent()}
          <CalendarPlus />
          <span>Create Club Event</span>
        {/snippet}
      </EventForm>
    </div>
  {/if}
</section>

<style>
  .space {
    margin-block: calc(var(--spacing) * 6);
  }
</style>