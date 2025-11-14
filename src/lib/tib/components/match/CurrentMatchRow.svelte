<script lang="ts">
  import type {Match} from "bsm.js";
  import {env} from "$env/dynamic/public";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";
  import {GameWinner} from "$lib/tib/enum/GameWinner.ts";
  import {MatchState} from "$lib/tib/enum/MatchState.ts";
  import {LogoUtility} from "$lib/tib/service/LogoUtility.ts";
  import {MatchDecorator} from "$lib/tib/service/MatchDecorator.ts";

  interface Props {
    match: Match;
  }

  const {match}: Props = $props();

  const awayLogo = LogoUtility.getLogoPathForTeamName(match.away_league_entry.team);
  const homeLogo = LogoUtility.getLogoPathForTeamName(match.home_league_entry.team);

  const matchDecorator = new MatchDecorator(match);
  const winner = matchDecorator.getWinnerForMatch();
  const matchState = matchDecorator.getMatchState(env.PUBLIC_TEAM_NAME);
  const matchDate = DateTimeUtility.parseDateFromBSMString(match.time);
</script>

<a class="grid grid-cols-3 px-3 py-1 gap-x-2 border-b border-surface-500 justify-around items-center"
   href="gamecenter/game-detail/{match.id}">
  <div class="flex justify-end items-center gap-2">
    <div>{match.away_league_entry.team?.short_name}</div>
    <img alt="team logo for {match.away_team_name}" loading="lazy" src="{awayLogo}" width="35"/>
  </div>

  {#if matchState === MatchState.notYetPlayed}

    <div class="flex justify-center items-center">
            <span class="preset-tonal-surface dark:preset-filled-surface-500 py-0.5 px-2 rounded-sm">
                {DateTimeUtility.timeFormatShort.format(matchDate)}
            </span>
    </div>

  {:else if matchState === MatchState.cancelled}

    <div class="flex justify-center items-center">
            <span class="preset-tonal-surface dark:preset-filled-surface-500 py-0.5 px-2 rounded-sm">
                PPD
            </span>
    </div>

  {:else}

    <div class="font-extrabold flex justify-center text-lg">
      <span class:text-surface-700-300={winner === GameWinner.home}>{match.away_runs ?? " "}</span>
      <span> : </span>
      <span class:text-surface-700-300={winner === GameWinner.away}>{match.home_runs ?? " "}</span>
    </div>

  {/if}

  <div class="flex justify-start items-center gap-2">
    <img alt="team logo for {match.away_team_name}" loading="lazy" src="{homeLogo}" width="35"/>
    <div>{match.home_league_entry.team?.short_name}</div>
  </div>
</a>