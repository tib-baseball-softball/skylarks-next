<script lang="ts">
  import type {ExpandedTeam} from "$lib/model/ExpandedResponse";
  import {client} from "$lib/pocketbase/index.svelte";
  import {getToastStore, ProgressBar, type ToastSettings,} from "@skeletonlabs/skeleton";
  import {CloseOutline} from "flowbite-svelte-icons";
  import type {LeaguegroupsResponse} from "$lib/model/pb-types";
  import type {GamesCount} from "$lib/model/GamesCount";
  import {range} from "$lib/functions/range";
  import {invalidate} from "$app/navigation";

  interface Props {
    team: ExpandedTeam,
    parent: any,
  }

  const {team, parent}: Props = $props();

  async function closeModal() {
    // @ts-ignore => this is why this is wrapped in another function
    parent.onClose();
  }

  const toastSettingsGeneralError: ToastSettings = {
    message: "An error occurred while saving the event.",
    background: "variant-filled-error",
  };
  const toastStore = getToastStore();

  const currentYear = new Date().getFullYear()
  const cutoffYear = currentYear - 4
  let season = $state(currentYear)
  const possibleSeasons = range(cutoffYear, currentYear)
  let leagueGroups: LeaguegroupsResponse[] = $state([])

  const form = $state(team ?? {
    id: "",
    bsm_league_group: 0,
  });

  async function getCurrentGamesCount(): Promise<number> {
    const response = await client.send<GamesCount>(`/api/gamecount/${team.id}?season=${season}`, {})
    return response.count ?? 0
  }

  async function loadClubLeagueGroups() {
    const RELOAD_DELAY = 15000 // 15 seconds
    const loadLeagueGroups = client
        .collection("leaguegroups")
        .getFullList<LeaguegroupsResponse>({
          filter: `season = '${season}' && clubs ?~ '${team.club}'`,
          query: {
            // adding those parameters triggers a hook to load league groups from BSM if response is empty
            season: season,
            club: team.club
          }
        });

    leagueGroups = await loadLeagueGroups

    // give the backend hook some time to get the requested LeagueGroups and try again
    if (leagueGroups.length === 0) {
      setTimeout(async () => {
            leagueGroups = await loadLeagueGroups
          }, RELOAD_DELAY
      )
    }
  }

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ExpandedTeam | null = null;

    try {
      if (form.id) {
        result = await client
            .collection("teams")
            .update<ExpandedTeam>(form.id, form);
      }
    } catch {
      toastStore.trigger(toastSettingsGeneralError);
    }

    if (result) {
      toastStore.trigger({
        message: "League has been successfully changed.",
        background: "variant-filled-success",
      });
      parent.onClose()
      invalidate("teams:list")
      invalidate("nav:load")
      invalidate("clubs:list")
    }
  }
</script>

<div class="w-modal">
    <article class="card p-4">

        <div class="flex items-center gap-5">
            <button
                    aria-label="cancel and close"
                    class="btn variant-ghost-surface"
                    onclick={() => closeModal()}
            >
                <CloseOutline/>
            </button>
            <header class="text-xl font-semibold">
                <h2 class="h3">Games Import Setup for {team.name}</h2>
            </header>
        </div>

        <form onsubmit={submitForm} class="mt-4 space-y-3">
            <div class="grid grid-cols-2 gap-2 md:gap-3 lg:gap-4">
                <h3 class="col-span-2 font-bold">Current Values</h3>
                <input
                        name="id"
                        autocomplete="off"
                        class="input col-span-2"
                        type="hidden"
                        readonly
                        bind:value={form.id}
                />

                <div>Current League ID:</div>
                <div class="badge variant-soft-primary text-lg">
                    {team.bsm_league_group !== 0 ? team.bsm_league_group : "None selected"}
                </div>

                {#await getCurrentGamesCount() then count}
                    <div>BSM-imported games in database:</div>
                    <div class="badge variant-soft-primary text-lg">
                        {count}
                    </div>
                {/await}

                <hr class="my-2 col-span-2">

                <label class="label col-span-2 md:col-span-1">

                    <span class="block">Show leagues for</span>
                    <select class="select"
                            name="season"
                            bind:value={season}
                    >
                        {#each possibleSeasons as season}
                            <option value={season}>{season}</option>
                        {/each}
                    </select>
                </label>

                {#await loadClubLeagueGroups()}
                    <ProgressBar/>
                    <div class="placeholder col-span-2"></div>
                {:then result}
                    <label class="label col-span-2 md:col-span-1">
                        <span class="block">League</span>

                        <select class="select" bind:value={form.bsm_league_group}>
                            <option value="{0}">None</option>
                            {#each leagueGroups as leagueGroup}
                                <option value="{leagueGroup.bsm_id}">{leagueGroup.name} ({leagueGroup.bsm_id})</option>
                            {/each}
                        </select>
                    </label>

                {:catch error}
                    <p class="col-span-2">error loading matches: {error.message}</p>
                {/await}

                <hr class="my-2 col-span-2">

                <p class="font-light col-span-2 text-sm">
                    To import games automatically, Diamond Planner needs to know the league this team is currently
                    playing in.
                    The league is imported from BSM and changes every season (even if the actual league name is still
                    the same).
                </p>

                <div class="flex justify-center gap-3 col-span-2">
                    <button type="button" class="mt-2 btn variant-ghost-surface !border-surface-50 border">
                        Help
                    </button>

                    <button type="submit" class="mt-2 btn variant-ghost-primary">
                        Submit
                    </button>
                </div>
            </div>
        </form>

    </article>
</div>
