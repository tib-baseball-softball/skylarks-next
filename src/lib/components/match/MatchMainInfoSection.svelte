<script lang="ts">
    import {LogoUtility} from "$lib/service/LogoUtility";
    import {DateTimeUtility} from "$lib/service/DateTimeUtility";
    import type {Match} from "bsm.js";
    import {GameWinner} from "$lib/enum/GameWinner";
    import {MatchDecorator} from "$lib/service/MatchDecorator";
    import {Calendar, ClipboardList} from "lucide-svelte";

    interface Props {
    match: Match;
  }

  let {match}: Props = $props();

  const awayLogo = LogoUtility.getLogoPathForTeamName(match.away_league_entry.team);
  const homeLogo = LogoUtility.getLogoPathForTeamName(match.home_league_entry.team);

  const matchDecorator = new MatchDecorator(match);

  const winner = matchDecorator.getWinnerForMatch();
  const matchDate = DateTimeUtility.parseDateFromBSMString(match.time);
</script>

<div class="card variant-ghost-primary px-2 md:px-5">
  <header class="pt-2 pb-3">
    <h2 class="h4 font-medium">Eckdaten</h2>
  </header>

  <div class="flex justify-center">
    <div class="grid grid-cols-3 sm:w-[80%] place-items-center gap-3">
      <img class="w-14 xl:w-24" src="{awayLogo}" alt="team logo for {match.away_team_name}" loading="lazy"/>

      <div class="font-bold text-3xl lg:text-6xl xl:text-7xl">
        {#if match.away_runs !== undefined && match.home_runs !== undefined}
          <span class:text-surface-500-400-token={winner === GameWinner.home}>{match.away_runs}</span>
          <span>&nbsp;-&nbsp;</span>
          <span class:text-surface-500-400-token={winner === GameWinner.away}>{match.home_runs}</span>
        {/if}
      </div>

      <img class="w-14 xl:w-24" src="{homeLogo}" alt="team logo for {match.home_team_name}" loading="lazy"/>

      <div>{match.away_team_name}</div>
      <div></div>
      <div>{match.home_team_name}</div>
    </div>
  </div>

  <div class="grid grid-cols-2 place-items-stretch gap-3 md:gap-8 my-2 md:my-5 pb-2 sm:pb-4 lg:pb-5">
    <div class="card variant-ghost-primary dark:variant-soft-primary p-3">
      <div class="flex gap-2">
        <Calendar/>
        <p class="mb-2 font-medium">Datum</p>
      </div>

      <p>{DateTimeUtility.dateTimeFormatMedium.format(matchDate)}</p>
    </div>
    <div class="card variant-ghost-primary dark:variant-soft-primary p-3">
      <div class="flex gap-2">
        <ClipboardList/>
        <p class="mb-2 font-medium">Status</p>
      </div>
      <p>{match.human_state}</p>
    </div>
  </div>

</div>