<script lang="ts">
import { LogoUtility } from "$lib/service/LogoUtility"
import { DateTimeUtility } from "$lib/service/DateTimeUtility"
import type { Match } from "bsm.js"
import { GameWinner } from "$lib/enum/GameWinner"
import { MatchDecorator } from "$lib/service/MatchDecorator"
import { Calendar, ClipboardList } from "lucide-svelte"

interface Props {
  match: Match
}

let { match }: Props = $props()

const awayLogo = LogoUtility.getLogoPathForTeamName(match.away_league_entry.team)
const homeLogo = LogoUtility.getLogoPathForTeamName(match.home_league_entry.team)

const matchDecorator = new MatchDecorator(match)

const winner = matchDecorator.getWinnerForMatch()
const matchDate = DateTimeUtility.parseDateFromBSMString(match.time)
</script>

<article class="card preset-tonal-primary border-primary px-2 md:px-5">
  <header class="pt-2 pb-3">
    <h2 class="h4 font-medium">Main Info</h2>
  </header>

  <div class="flex justify-center">
    <div class="grid grid-cols-3 sm:w-[80%] place-items-center gap-3">
      <img alt="team logo for {match.away_team_name}" class="w-14 xl:w-24" loading="lazy" src="{awayLogo}"/>

      <div class="font-bold text-2xl sm:text-3xl lg:text-4xl xl:text-5xl 2xl:text-6xl">
        {#if match.away_runs !== undefined && match.home_runs !== undefined}
          <span class:text-surface-600-400={winner === GameWinner.home}>{match.away_runs}</span>
          <span> - </span>
          <span class:text-surface-600-400={winner === GameWinner.away}>{match.home_runs}</span>
        {/if}
      </div>

      <img alt="team logo for {match.home_team_name}" class="w-14 xl:w-24" loading="lazy" src="{homeLogo}"/>

      <div>{match.away_team_name}</div>
      <div></div>
      <div>{match.home_team_name}</div>
    </div>
  </div>

  <div class="grid grid-cols-2 place-items-stretch gap-3 md:gap-8 my-2 md:my-5 pb-2 sm:pb-4 lg:pb-5">
    <div class="card sub-item preset-tonal-primary p-3">
      <div class="flex gap-2">
        <Calendar/>
        <p class="mb-2 font-medium">Date</p>
      </div>

      <p>{DateTimeUtility.dateTimeFormatMedium.format(matchDate)}</p>
    </div>
    <div class="card sub-item preset-tonal-primary p-3">
      <div class="flex gap-2">
        <ClipboardList/>
        <p class="mb-2 font-medium">Status</p>
      </div>
      <p>{match.human_state}</p>
    </div>
  </div>

</article>

<style>
    article {
        border: 2px solid var(--color-primary-500)
    }

    .sub-item {
        border: 2px solid var(--color-primary-500)
    }
</style>