<script lang="ts">
  import {CalendarPlus} from "lucide-svelte";
  import {goto, invalidateAll} from "$app/navigation";
  import EventSeriesView from "$lib/dp/components/event/EventSeriesView.svelte";
  import EventForm from "$lib/dp/components/forms/EventForm.svelte";
  import TeamForm from "$lib/dp/components/forms/TeamForm.svelte";
  import TeamGamesModal from "$lib/dp/components/forms/TeamGamesModal.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedEventSeries, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";

  interface Props {
    team: ExpandedTeam;
    eventSeries: ExpandedEventSeries[];
  }

  const {team, eventSeries}: Props = $props();

  function teamDeleteAction(id: string) {
    goto(`/account`);
    client.collection("teams").delete(id);
    invalidateAll();
  }

  const model = $derived(authSettings.record as CustomAuthModel);
</script>

<h2 class="section-title">Admin Section</h2>

<div class="outer-grid">
  <div class="card admin-card preset-outlined-surface-500">
    <div>
      <header class="card-header">
        <h3 class="card-title">Games</h3>
      </header>

      <section class="card-content">
        <p class="info-text">
          Events of type "Game" don't need to be created manually.
          All data can be imported from BSM and kept in sync with
          the current state.
        </p>
      </section>
    </div>
    <footer class="card-footer">
      <div class="card-actions">

        <Dialog triggerClasses="btn preset-tonal-primary border-primary">

          {#snippet triggerContent()}
            <CalendarPlus/>
            <span>Setup Games Import</span>
          {#/snippet}

          {#snippet title()}
            <header>
              <h3 class="h3">Games Import Setup for {team.name}</h3>
            </header>
          {#/snippet}

          <TeamGamesModal {team}/>
        </Dialog>
      </div>
    </footer>
  </div>

  <div class="card admin-card preset-outlined-surface-500">
    <div>
      <header class="card-header">
        <h3 class="card-title">Event Series</h3>
      </header>

      <section class="card-content">
        <p class="info-text">
          Manage all recurring series of events (mostly of type
          "Practice") that take place on the same day every week.
        </p>
      </section>
    </div>
    <footer class="card-footer">
      <div class="card-actions">
        <EventSeriesView
          triggerVariant="tonal-primary"
          {eventSeries}
          {team}
        />
      </div>
    </footer>
  </div>

  <div class="card admin-card preset-outlined-surface-500">
    <div>
      <header class="card-header">
        <h3 class="card-title">Single Events</h3>
      </header>

      <section class="card-content">
        <p class="info-text">
          Single Events are not connected to any other events and
          don't contain additional logic. Well suited for one-off
          occurrences.
        </p>
      </section>
    </div>
    <footer class="card-footer">
      <div class="card-actions">

        <EventForm
          triggerVariant="tonal-tertiary"
          clubID={team?.club ?? ""}
          event={null}
          teamID={team.id}
        >
          {#snippet triggerContent()}
            <CalendarPlus/>
            <span>Create Single Event</span>
          {#/snippet}
        </EventForm>

      </div>
    </footer>
  </div>

  <div class="card admin-card preset-outlined-surface-500">
    <header class="card-header">
      <h3 class="card-title">Team Settings</h3>
    </header>

    <section class="card-content">
      <div class="card-actions">
        <TeamForm triggerVariant="tonal-surface" club={team?.expand?.club}
                  team={team}/>
      </div>
    </section>

    {#if team?.expand?.club?.admins.includes(model.id)}
      <footer class="card-footer danger-zone-footer">
        <div class="danger-zone">
          <header class="danger-header">Danger Zone</header>

          <DeleteButton
            id={team.id}
            modelName="Team"
            action={teamDeleteAction}
            classes="preset-tonal-error border-error mx-1"
            buttonText="Delete Team"
          />
        </div>
      </footer>
    {/if}
  </div>
</div>

<style>
  .section-title {
    font-size: var(--text-xl);
    font-weight: bold;
    margin-block: calc(var(--spacing) * 4);
  }

  .admin-card {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }

  .card-title {
    font-size: var(--text-lg);
    font-weight: 600;
  }

  .card-content {
    padding: calc(var(--spacing) * 4);
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 3);
  }

  .info-text {
    font-weight: 300;
  }

  .card-actions {
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 2);

    @media (min-width: 64rem) {
      gap: calc(var(--spacing) * 3);
    }
  }

  .danger-zone-footer {
    margin-top: calc(var(--spacing) * 2);
  }

  .danger-zone {
      display: flex;
      flex-direction: column;
      gap: calc(var(--spacing) * 2);
      border-radius: var(--radius-base);
      border: 1px solid var(--color-error-500);
      padding-inline: calc(var(--spacing) * 2);
      padding-block: calc(var(--spacing) * 3);
      
      @media (min-width: 64rem) {
          gap: calc(var(--spacing) * 3);
      }
  }

  .danger-header {
      margin-inline: calc(var(--spacing) * 2);
  }

  .outer-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 2);
    margin-block-end: calc(var(--spacing) * 6);

    @media (min-width: 48rem) {
      grid-template-columns: 1fr 1fr;
    }

    @media (min-width: 64rem) {
      grid-template-columns: 1fr 1fr 1fr;
      gap: calc(var(--spacing) * 3);
    }

    @media (min-width: 80rem) {
      grid-template-columns: 1fr 1fr 1fr 1fr;
    }
  }

  :global(.border-primary) {
      border: 1px solid var(--color-primary-500) !important;
  }

  :global(.border-tertiary) {
      border: 1px solid var(--color-tertiary-500) !important;
  }

  :global(.border-surface) {
      border: 1px solid var(--color-surface-500) !important;
  }

  :global(.border-error) {
      border: 1px solid var(--color-error-500) !important;
  }

  :global(.mx-1) {
      margin-inline: calc(var(--spacing) * 1) !important;
  }
</style>
