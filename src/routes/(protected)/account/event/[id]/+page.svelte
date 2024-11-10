<script lang="ts">
  import EventAttireSection from "$lib/components/diamondplanner/event/EventAttireSection.svelte";
  import EventCoreInfo from "$lib/components/diamondplanner/event/EventCoreInfo.svelte";
  import EventParticipationSection from "$lib/components/diamondplanner/event/EventParticipationSection.svelte";
  import EventTypeBadge from "$lib/components/diamondplanner/event/EventTypeBadge.svelte";
  import MatchTeaserCard from "$lib/components/match/MatchTeaserCard.svelte";
  import {
    CalendarEditOutline,
    CalendarPlusOutline,
    CheckOutline,
    EditOutline,
    InfoCircleOutline,
    QuestionCircleOutline,
    TrashBinOutline,
  } from "flowbite-svelte-icons";
  import CloseOutline from "flowbite-svelte-icons/CloseOutline.svelte";
  import {type DrawerSettings, getDrawerStore,} from "@skeletonlabs/skeleton";
  import {authSettings} from "$lib/pocketbase/index.svelte";

  const drawerStore = getDrawerStore();

  let {data} = $props();

  const event = $derived(data.event);

  const settings: DrawerSettings = $derived({
    id: "event-form",
    position: "right",
    width: "w-[100%] sm:w-[80%] lg:w-[70%] xl:w-[50%]",
    meta: {
      event: $event,
      club: $event?.expand?.team?.club,
      team: $event?.expand?.team?.id,
    },
  });
  // TODO: get this data via backend hook

  let participations = $derived(
      $event?.expand?.participations_via_event ?? [],
  );

  let participationsIn = $derived(
      participations?.filter((participation) => {
        return participation.state === "in";
      }),
  );
  let participationsMaybe = $derived(
      participations?.filter((participation) => {
        return participation.state === "maybe";
      }),
  );
  let participationsOut = $derived(
      participations?.filter((participation) => {
        return participation.state === "out";
      }),
  );
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
        <section>
            <EventAttireSection attire={$event.expand.attire}/>
        </section>
    {/if}

    {#if !$event.cancelled}
        <hr class="!my-8"/>

        <div class="flex justify-between items-center">
            <h2 class="h3">My Participation</h2>
            <EventParticipationSection event={$event}/>
        </div>

        <hr class="my-8"/>

        <h2 class="h2">Participants</h2>
        <section class="grid grid-cols-1 md:grid-cols-3 gap-2 lg:gap-3">
            <div class="card variant-ghost-success flex-grow">
                <header class="card-header flex items-center gap-2">
                    <span><CheckOutline size="lg"/></span>
                    <h3 class="h4">In</h3>
                </header>

                <section class="p-4 flex flex-wrap gap-1 lg:gap-2">
                    {#each participationsIn as inResponse}
                        <div class="chip variant-ghost-success">
                            {inResponse?.expand?.user?.first_name}
                        </div>
                    {/each}
                </section>
            </div>

            <div class="card variant-ghost-warning flex-grow">
                <header class="card-header flex items-center gap-2">
                    <span><QuestionCircleOutline size="lg"/></span>
                    <h3 class="h4">Maybe</h3>
                </header>
                <section class="p-4 flex flex-wrap gap-1 lg:gap-2">
                    {#each participationsMaybe as maybeResponse}
                        <div class="chip variant-ghost-warning">
                            {maybeResponse?.expand?.user?.first_name}
                        </div>
                    {/each}
                </section>
            </div>

            <div class="card variant-ghost-error flex-grow">
                <header class="card-header flex items-center gap-2">
                    <span><CloseOutline size="lg"/></span>
                    <h3 class="h4">Out</h3>
                </header>
                <section class="p-4 flex flex-wrap gap-1 lg:gap-2">
                    {#each participationsOut as outResponse}
                        <div class="chip variant-ghost-error">
                            {outResponse?.expand?.user?.first_name}
                        </div>
                    {/each}
                </section>
            </div>
        </section>
    {/if}

    {#if $event.match_json}
        <section>
            <h2 class="h2 mb-3">Game Data</h2>
            <div class="grid grid-cols-1 md:grid-cols-2">
                <MatchTeaserCard match={$event.match_json}/>
            </div>
        </section>
    {/if}

    {#if $event.expand?.team?.admins.includes(authSettings?.record?.id)}
        <hr class="my-2"/>

        <h2 class="h2">Admin Section</h2>

        <div
                class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2 lg:gap-3"
        >
            <div class="card variant-ringed-surface">
                <header class="card-header">
                    <h3 class="h4 font-semibold">Event Details</h3>
                </header>

                <section class="p-4 space-y-2">
                    <div class="flex items-center gap-3">
                        <InfoCircleOutline/>
                        <div>
                            <p class="text-sm font-light">BSM event ID</p>
                            <p>{$event.bsm_id}</p>
                        </div>
                    </div>

                    <div class="flex items-center gap-3">
                        <CalendarPlusOutline/>
                        <div>
                            <p class="text-sm font-light">Created</p>
                            <p>{new Date($event.created).toLocaleString()}</p>
                        </div>
                    </div>

                    <div class="flex items-center gap-3">
                        <CalendarEditOutline/>
                        <div>
                            <p class="text-sm font-light">Last Updated</p>
                            <p>{new Date($event.updated).toLocaleString()}</p>
                        </div>
                    </div>
                </section>
            </div>

            <div class="card variant-ringed-surface">
                <header class="card-header">
                    <h3 class="h4 font-semibold">Add Guest player</h3>
                </header>

                <section class="p-4 space-y-2">
                    <form
                            onsubmit={() => console.log("guest player submitted")}
                    >
                        <label class="label">
                            Name
                            <input
                                    type="text"
                                    class="input mt-1"
                                    placeholder="Enter the guest player's name"
                            />
                        </label>
                        <button type="submit" class="mt-3 btn variant-soft"
                        >Submit
                        </button
                        >
                    </form>
                </section>
            </div>

            <div class="card variant-ringed-surface">
                <header class="card-header">
                    <h3 class="h4 font-semibold">Actions</h3>
                </header>

                <section class="p-4 space-y-3">
                    <div class="flex flex-col gap-2 lg:gap-3">
                        <button
                                class="btn variant-ghost-surface"
                                onclick={() => drawerStore.open(settings)}
                        >
                            <EditOutline/>
                            <span>Edit Event</span>
                        </button>

                        <button class="btn variant-ghost-error">
                            <TrashBinOutline/>
                            <span>Delete Event</span>
                        </button>
                    </div>
                </section>
            </div>
        </div>
    {/if}
</div>

<style lang="postcss">
    .event-container {
        @apply space-y-4 lg:space-y-6 xl:space-y-7;
    }
</style>
