<script lang="ts">
  import type {Match} from "bsm.js";
  import GameResultIndicator from "$lib/tib/components/match/GameResultIndicator.svelte";
  import MatchTeaserCardRow from "$lib/tib/components/match/MatchTeaserCardRow.svelte";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";
  import {GameWinner} from "$lib/tib/enum/GameWinner.ts";
  import {LogoUtility} from "$lib/tib/service/LogoUtility.ts";
  import {MatchDecorator} from "$lib/tib/service/MatchDecorator.ts";

  interface Props {
    match: Match;
  }

  const {match}: Props = $props();

  const awayLogo = LogoUtility.getLogoPathForTeamName(match?.away_league_entry?.team);
  const homeLogo = LogoUtility.getLogoPathForTeamName(match?.home_league_entry?.team);

  const matchDecorator = new MatchDecorator(match);
  const winner = matchDecorator.getWinnerForMatch();
  const matchDate = DateTimeUtility.parseDateFromBSMString(match.time);
</script>

<a
        class=" block card card-hover text-sm {matchDecorator.isPlayoffGame() ===
    true
        ? 'preset-tonal-primary border border-primary-500'
        : 'preset-tonal-surface border border-surface-500'}"
        href="/gamecenter/game-detail/{match.id}"
>
  <header class="card-header">
    <div class="flex justify-between items-center">
      <div>
        <h3 class="font-medium">{match.league?.name}</h3>
        <p class="font-light">
          {DateTimeUtility.dateTimeFormatMedium.format(matchDate)}
        </p>
      </div>
      <GameResultIndicator {match}/>
    </div>
  </header>

  <hr class="mt-2 mx-3"/>

  <section class="p-4 flex flex-col gap-2">
    <MatchTeaserCardRow
            logo={awayLogo}
            score={match.away_runs}
            subduedScore={winner === GameWinner.home}
            teamName={match.away_team_name}
    />
    <MatchTeaserCardRow
            logo={homeLogo}
            score={match.home_runs}
            subduedScore={winner === GameWinner.away}
            teamName={match.home_team_name}
    />
  </section>
</a>
