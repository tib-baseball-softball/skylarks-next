<script lang="ts">
  import DeleteButton from "$lib/components/utility/DeleteButton.svelte";
  import {
    type DrawerSettings,
    getDrawerStore,
    getModalStore,
    type ModalComponent,
    type ModalSettings
  } from "@skeletonlabs/skeleton";
  import {goto, invalidateAll} from "$app/navigation";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/model/ExpandedResponse";
  import TeamGamesModal from "$lib/components/forms/TeamGamesModal.svelte";
  import type {EventseriesResponse} from "$lib/model/pb-types.ts";
  import {CalendarPlus} from "lucide-svelte";
  import EventSeriesView from "$lib/components/diamondplanner/event/EventSeriesView.svelte";
  import TeamForm from "$lib/components/forms/TeamForm.svelte";

  interface Props {
    team: ExpandedTeam,
    eventSeries: EventseriesResponse[],
  }

  let {team, eventSeries}: Props = $props();

  const drawerStore = getDrawerStore();
  const modalStore = getModalStore();

  const singleEventSettings: DrawerSettings = $derived({
    id: "event-form",
    position: "right",
    width: "w-[100%] sm:w-[80%] lg:w-[70%] xl:w-[50%]",
    meta: {
      event: null,
      club: team?.club,
      team: team,
    },
  });

  function teamDeleteAction(id: string) {
    goto(`/account`);
    client.collection("teams").delete(id);
    invalidateAll();
  }

  function triggerModal() {
    const modalComponent: ModalComponent = {
      ref: TeamGamesModal,
      props: {team: team},
    };

    const modal: ModalSettings = {
      type: "component",
      component: modalComponent,
    };
    modalStore.trigger(modal);
  }

  const model = $derived(authSettings.record as CustomAuthModel);
</script>

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
        <button class="btn variant-ghost-primary" onclick={() => triggerModal()}>
          <CalendarPlus/>
          <span>Setup Games Import</span>
        </button>
      </div>
    </footer>
  </div>

  <div class="card admin-card variant-ringed-surface">
    <div>
      <header class="card-header">
        <h3 class="h4 font-semibold">Event Series</h3>
      </header>

      <section class="p-4 space-y-3">
        <p class="font-light">
          Manage all recurring series of events (mostly of type
          "Practice") that take place on the same day every week.
        </p>
      </section>
    </div>
    <footer class="card-footer">
      <div class="flex flex-col gap-2 lg:gap-3">
        <EventSeriesView
                {team}
                {eventSeries}
                buttonClasses="btn variant-ghost-secondary dark:variant-filled-secondary dark:border"
        />
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
          <CalendarPlus/>
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
        <TeamForm club={team?.expand?.club} team={team} buttonClasses="btn variant-ghost-surface"/>
      </div>
    </section>

    {#if team?.expand?.club?.admins.includes(model.id)}
      <footer class="card-footer mt-2">
        <div class="flex flex-col gap-2 lg:gap-3 rounded-token variant-ringed-error px-2 py-3">
          <header class="mx-2">Danger Zone</header>

          <DeleteButton
                  id={team.id}
                  modelName="Team"
                  action={teamDeleteAction}
                  classes="variant-ghost-error mx-1"
                  buttonText="Delete Team"
          />
        </div>
      </footer>
    {/if}
  </div>
</div>

<style lang="postcss">
    .admin-card {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
    }
</style>