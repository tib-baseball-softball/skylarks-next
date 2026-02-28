<script lang="ts">
  import {invalidate} from "$app/navigation";
  import ProgressRing from "$lib/dp/components/utils/ProgressRing.svelte";
  import {client} from "$lib/dp/client.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import type {GamesCount} from "$lib/dp/types/GamesCount.ts";
  import type {LeaguegroupsResponse} from "$lib/dp/types/pb-types.ts";
  import type {Toast} from "$lib/dp/types/Toast.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";
  import {range} from "$lib/dp/utility/range.ts";
  import { Collection } from "$lib/dp/enum/Collection";

  interface Props {
    team: ExpandedTeam;
  }

  const {team}: Props = $props();

  const toastSettingsGeneralError: Toast = {
    message: "An error occurred while saving the event.",
    background: "preset-filled-error-500",
  };

  const currentYear = new Date().getFullYear();
  const cutoffYear = currentYear - 4;
  let season = $state(currentYear);
  const possibleSeasons = range(cutoffYear, currentYear);
  let leagueGroups: LeaguegroupsResponse[] = $state([]);

  function formFromProps(data: ExpandedTeam) {
    return data ?? {
      id: "",
      bsm_league_group: 0,
    };
  }

  let form = $derived.by(() => {
    const formData = $state(formFromProps(team));
    return formData;
  });

  async function getCurrentGamesCount(): Promise<number> {
    const response = await client.send<GamesCount>(`/api/dp/teams/${team.id}/gamecount?season=${season}`, {});
    return response.count ?? 0;
  }

  async function loadClubLeagueGroups() {
    const RELOAD_DELAY = 15000; // 15 seconds
    const loadLeagueGroups = client.collection(Collection.LeagueGroups).getFullList<LeaguegroupsResponse>({
      filter: `season = '${season}' && clubs ?~ '${team.club}'`,
      query: {
        // adding those parameters triggers a hook to load league groups from BSM if response is empty
        season: season,
        club: team.club,
      },
    });

    leagueGroups = await loadLeagueGroups;

    // give the backend hook some time to get the requested LeagueGroups and try again
    if (leagueGroups.length === 0) {
      setTimeout(async () => {
        leagueGroups = await loadLeagueGroups;
      }, RELOAD_DELAY);
    }
  }

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ExpandedTeam | null = null;

    try {
      if (form.id) {
        result = await client.collection("teams").update<ExpandedTeam>(form.id, form);
      }
    } catch {
      toastController.trigger(toastSettingsGeneralError);
    }

    if (result) {
      toastController.trigger({
        message: "League has been successfully changed.",
        background: "preset-filled-success-500",
      });
      closeModal();
      await invalidate("teams:list");
      await invalidate("nav:load");
      await invalidate("clubs:list");
    }
  }
</script>

<form class="team-games-form" onsubmit={submitForm}>
  <div class="form-grid">
    <h3 class="section-title">Current Values</h3>
    <input
      autocomplete="off"
      bind:value={form.id}
      class="input col-full"
      name="id"
      readonly
      type="hidden"
    />

    <div>Current League ID:</div>
    <div class="badge preset-tonal-primary info-badge">
      {team.bsm_league_group !== 0 ? team.bsm_league_group : "None selected"}
    </div>

    {#await getCurrentGamesCount() then count}
      <div>BSM-imported games in database:</div>
      <div class="badge preset-tonal-primary info-badge">
        {count}
      </div>
    {/await}

    <hr class="divider">

    <label class="label form-label">

      <span class="label-text">Show leagues for</span>
      <select bind:value={season}
              class="select"
              name="season"
      >
        {#each possibleSeasons as season}
          <option value={season}>{season}</option>
        {/each}
      </select>
    </label>

    {#await loadClubLeagueGroups()}
      <ProgressRing/>
      <div class="placeholder col-full"></div>
    {:then result}
      <label class="label form-label">
        <span class="label-text">League</span>

        <select class="select" bind:value={form.bsm_league_group}>
          <option value="{0}">None</option>
          {#each leagueGroups as leagueGroup}
            <option value="{leagueGroup.bsm_id}">{leagueGroup.name} ({leagueGroup.acronym}, {leagueGroup.season})</option>
          {/each}
        </select>
      </label>

    {:catch error}
      <p class="col-full">error loading matches: {error.message}</p>
    {/await}

    <hr class="divider">

    <p class="footer-info col-full">
      To import games automatically, Diamond Planner needs to know the league this team is currently
      playing in.
      The league is imported from BSM and changes every season (even if the actual league name is still
      the same).
    </p>

    <div class="actions col-full">
      <button class="btn preset-filled-primary-500" type="submit">
        Submit
      </button>
    </div>
  </div>
</form>

<style>
  .team-games-form {
    margin-top: calc(var(--spacing) * 4);
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 3);
  }

  .form-grid {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: calc(var(--spacing) * 2);

    @media (min-width: 48rem) {
      gap: calc(var(--spacing) * 3);
    }

    @media (min-width: 64rem) {
      gap: calc(var(--spacing) * 4);
    }
  }

  .col-full {
    grid-column: span 2 / span 2;
  }

  .section-title {
    grid-column: span 2 / span 2;
    font-weight: bold;
  }

  .info-badge {
    border: 1px solid var(--color-primary-500);
    font-size: var(--text-lg);
  }

  .divider {
    grid-column: span 2 / span 2;
    margin-block: calc(var(--spacing) * 2);
  }

  .form-label {
    grid-column: span 2 / span 2;

    @media (min-width: 48rem) {
      grid-column: span 1 / span 1;
    }

    .label-text {
      display: block;
    }
  }

  .footer-info {
    font-weight: 300;
    font-size: var(--text-sm);
  }

  .actions {
    display: flex;
    justify-content: center;
    gap: calc(var(--spacing) * 3);

    button {
      margin-top: calc(var(--spacing) * 2);
    }
  }
</style>
