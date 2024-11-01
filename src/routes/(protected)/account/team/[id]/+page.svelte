<script lang="ts">
    import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";
    import EventTeaser from "$lib/components/diamondplanner/event/EventTeaser.svelte";
    import Paginator from "$lib/pocketbase/Paginator.svelte";
    import {authModel} from "$lib/pocketbase/Auth";
    import {CalendarPlusOutline, UsersGroupOutline,} from "flowbite-svelte-icons";
    import {goto, invalidateAll} from "$app/navigation";
    import type {EventType} from "$lib/model/ExpandedResponse.js";
    import {type DrawerSettings, getDrawerStore, RadioGroup, RadioItem,} from "@skeletonlabs/skeleton";
    import DeleteButton from "$lib/components/utility/DeleteButton.svelte";
    import {client} from "$lib/pocketbase";
    import TeamEditButton from "$lib/components/team/TeamEditButton.svelte";

    let {data} = $props();
    const events = $derived(data.events);
    let currentPage = $derived($events.page);

    let showEvents = $state("next");
    let sorting: "asc" | "desc" = $state("asc");
    let showTypes: EventType | "any" = $state("any");

    const reloadWithQuery = () => {
        let queryString = `?timeframe=${showEvents}&page=${currentPage}&sort=${sorting}&type=${showTypes}`;

        goto(queryString, {
            noScroll: true,
        });
    };

    const drawerStore = getDrawerStore();
    const singleEventSettings: DrawerSettings = $derived({
        id: "event-form",
        position: "right",
        width: "w-[100%] sm:w-[80%] lg:w-[70%] xl:w-[50%]",
        meta: {
            event: null,
            club: data.team?.club,
            team: data.team,
        },
    });

    function teamDeleteAction(id: string) {
        goto(`/account`)
        client.collection("teams").delete(id)
        invalidateAll()
    }

    $effect.pre(() => {
        console.log(showEvents);
        reloadWithQuery();
    });
</script>

<h1 class="h2">{data.team.name} ({data.team?.expand?.club.name})</h1>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3 mb-3">
    <article class="card variant-soft-surface lg:col-span-2">
        <header class="card-header">
            <h2 class="h4 font-medium">Team Description</h2>
        </header>
        <section class="p-4">{@html data.team.description}</section>
    </article>

    <TeamTeaserCard team={data.team} link={false}/>
</div>

<h2 class="h3">Team Events</h2>

<div
        class="flex flex-wrap gap-4 lg:gap-8 variant-soft-surface justify-between px-4 py-3 rounded-token"
>
    <label
            class="flex items-center gap-2 flex-grow justify-between md:flex-grow-0"
    >
        Timeframe
        <RadioGroup>
            <RadioItem
                    checked
                    name="radio-next"
                    value="next"
                    bind:group={showEvents}
            >
                Next
            </RadioItem>
            <RadioItem name="radio-past" value="past" bind:group={showEvents}>
                Past
            </RadioItem>
        </RadioGroup>
    </label>

    <label class="label flex items-center gap-2 flex-grow justify-between md:flex-grow-0">
        Sort
        <RadioGroup>
            <RadioItem checked name="radio-asc" value="asc" bind:group={sorting}>Ascending</RadioItem>
            <RadioItem checked name="radio-desc" value="desc" bind:group={sorting}>Descending</RadioItem>
        </RadioGroup>
    </label>

    <label class="label flex items-center gap-2 flex-grow justify-between md:flex-grow-0">
        Type
        <RadioGroup padding="px-2 md:px-4 py-1">
            <RadioItem padding="px-4 py-1" checked name="radio-any" value="any" bind:group={showTypes}>All</RadioItem>
            <RadioItem checked name="radio-game" value="game" bind:group={showTypes}>Game</RadioItem>
            <RadioItem checked name="radio-practice" value="practice" bind:group={showTypes}>Practice</RadioItem>
            <RadioItem checked name="radio-misc" value="misc" bind:group={showTypes}>Other</RadioItem>
        </RadioGroup>
    </label>
</div>

<div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-3">
    {#each $events.items as event}
        <EventTeaser {event} link={true}/>
    {:else}
        <p>No events available with the current filters.</p>
    {/each}
</div>

<Paginator store={events} showIfSinglePage={false}/>

<hr class="!my-8"/>
<section class="space-y-2 lg:space-y-4">
    <header>
        <h2 class="h3">Links</h2>
    </header>

    <div class="flex flex-wrap items-center gap-2 lg:gap-3">
        <a
                href="/account/team/{data.team.id}/members"
                class="btn variant-ghost-tertiary"
        >
            <UsersGroupOutline/>
            <span>Player List</span>
        </a>
    </div>
</section>

{#if data.team.admins.includes($authModel?.id)}
    <hr class="my-2"/>

    <h2 class="h3">Admin Section</h2>

    <div
            class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-2 lg:gap-3"
    >
        <div class="card admin-card variant-ringed-surface">
            <div>
                <header class="card-header">
                    <h3 class="h4 font-semibold">Games</h3>
                </header>

                <section class="p-4 space-y-3">
                    <p class="font-light">
                        Events of type "Game" don't need to be created manually.
                        All data can be imported from BSM and kept in sync with
                        the current state.
                    </p>
                </section>
            </div>
            <footer class="card-footer">
                <div class="flex flex-col gap-2 lg:gap-3">
                    <button class="btn variant-ghost-primary">
                        <CalendarPlusOutline/>
                        <span>Setup Games Import</span>
                    </button>
                </div>
            </footer>
        </div>

        <div class="card admin-card variant-ringed-surface">
            <div>
                <header class="card-header">
                    <h3 class="h4 font-semibold">Practice Sets</h3>
                </header>

                <section class="p-4 space-y-3">
                    <p class="font-light">
                        Practice Sets are recurring series of events of type
                        "Practice" that take place on the same day every week.
                    </p>
                </section>
            </div>
            <footer class="card-footer">
                <div class="flex flex-col gap-2 lg:gap-3">
                    <button
                            class="btn variant-ghost-secondary dark:variant-filled-secondary dark:border"
                    >
                        <CalendarPlusOutline/>
                        <span>Create Practice Set</span>
                    </button>
                </div>
            </footer>
        </div>

        <div class="card admin-card variant-ringed-surface">
            <div>
                <header class="card-header">
                    <h3 class="h4 font-semibold">Single Events</h3>
                </header>

                <section class="p-4 space-y-3">
                    <p class="font-light">
                        Single Events are not connected to any other events and
                        don't contain additional logic. Well suited for one-off
                        occurrences.
                    </p>
                </section>
            </div>
            <footer class="card-footer">
                <div class="flex flex-col gap-2 lg:gap-3">
                    <button
                            class="btn variant-ghost-tertiary"
                            onclick={() => drawerStore.open(singleEventSettings)}
                    >
                        <CalendarPlusOutline/>
                        <span>Create Single Event</span>
                    </button>
                </div>
            </footer>
        </div>

        <div class="card admin-card variant-ringed-surface">
            <header class="card-header">
                <h3 class="h4 font-semibold">Team Settings</h3>
            </header>

            <section class="p-4 space-y-3">
                <div class="flex flex-col gap-2 lg:gap-3">
                    <TeamEditButton club={data.team?.expand?.club} team={data.team} classes="variant-ghost-surface"/>
                </div>
            </section>

            <footer class="card-footer mt-2">
                <div class="flex flex-col gap-2 lg:gap-3 rounded-token variant-ringed-error px-2 py-3">
                    <header class="mx-2">Danger Zone</header>

                    <DeleteButton
                            id={data.team.id}
                            modelName="Team"
                            action={teamDeleteAction}
                            classes="variant-ghost-error mx-1"
                            buttonText="Delete Team"
                    />
                </div>
            </footer>
        </div>
    </div>
{/if}

<style lang="postcss">
    .admin-card {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
    }
</style>
