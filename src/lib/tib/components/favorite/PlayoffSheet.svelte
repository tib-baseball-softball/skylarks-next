<script lang="ts">
  import {Trophy} from "lucide-svelte";
  import MatchTeaserCard from "$lib/tib/components/match/MatchTeaserCard.svelte";
  //@ts-expect-error
  import * as Sheet from "$lib/dp/components/modal/sheet";
  import type {PlayoffSeries} from "$lib/tib/types/PlayoffSeries.ts";

  interface Props {
    playoffSeries: PlayoffSeries;
  }

  const {playoffSeries}: Props = $props();

  const open = $state(false);
  const teams = $derived(Object.values(playoffSeries.teams));
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger class="btn preset-filled-primary-500 p-2 my-2">
    <Trophy/>
    <span>Show Playoff Series</span>
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header>
      <h2 class="h3 mb-3">Playoff Series</h2>
    </Sheet.Header>

    <div class="mx-2 my-4">
      <article class="card preset-outlined-primary p-4">
        <div class="flex flex-col items-center gap-3">
          {#if playoffSeries.status === "leading"}
            <p>{playoffSeries.leading_team.name} lead series</p>
            <p class="series-score">{playoffSeries.leading_team.wins} - {playoffSeries.trailing_team.wins}</p>
          {:else if playoffSeries.status === "won"}
            <p>{playoffSeries.leading_team.name} win series</p>
            <p class="series-score">{playoffSeries.leading_team.wins} - {playoffSeries.trailing_team.wins}</p>
          {:else if playoffSeries.status === "tied"}
            <p>Series tied</p>
            <p>{teams.at(0)?.wins} - {teams.at(1)?.wins}</p>
          {/if}
          <p class="font-light text-sm">Length of series: Best of {playoffSeries.series_length}</p>
        </div>

      </article>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 m-2">
      {#each playoffSeries.series_games as match}
        <MatchTeaserCard {match}/>
      {/each}
    </div>
  </Sheet.Content>
</Sheet.Root>

<style>
    .series-score {
        font-size: 2rem;
        font-weight: bold;
    }
</style>