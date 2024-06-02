<script lang="ts">
    import {LogoUtility} from "$lib/service/LogoUtility";
    import {MatchUtility} from "$lib/service/MatchUtility";
    import {DateTimeUtility} from "$lib/service/DateTimeUtility";
    import type {Match} from "bsm.js"
    import {GameWinner} from "$lib/enum/GameWinner";

    export let match: Match

    const awayLogo = LogoUtility.getLogoPathForTeamName(match.away_team_name)
    const homeLogo = LogoUtility.getLogoPathForTeamName(match.home_team_name)

    const winner = MatchUtility.getWinnerForMatch(match)
    const matchDate = DateTimeUtility.parseDateFromBSMString(match.time)
</script>

<div class="card variant-ghost-primary px-2 md:px-5">
    <header class="pt-2 pb-3">
        <h2 class="h4 font-medium">Eckdaten</h2>
    </header>
    <div class="flex justify-center">
        <div class="grid grid-cols-3 sm:w-[80%] place-items-center gap-3">
            <img class="w-14 xl:w-24" src="{awayLogo}" alt="team logo for {match.away_team_name}"/>

            <div class="font-bold text-4xl lg:text-6xl xl:text-7xl">
                <span class:text-surface-500-400-token={winner === GameWinner.home}>{match.away_runs}</span>
                <span>&nbsp;-&nbsp;</span>
                <span class:text-surface-500-400-token={winner === GameWinner.away}>{match.home_runs}</span>
            </div>

            <img class="w-14 xl:w-24" src="{homeLogo}" alt="team logo for {match.home_team_name}"/>

            <div>{match.away_team_name}</div>
            <div></div>
            <div>{match.home_team_name}</div>
        </div>
    </div>
    <div class="grid grid-cols-2 place-items-stretch gap-3 md:gap-8 my-2 md:my-5 pb-2 sm:pb-4 lg:pb-5">
        <div class="card variant-ghost-primary dark:variant-soft-primary p-3">
            <p class="mb-2 font-medium">Datum</p>
            <p>{DateTimeUtility.dateTimeFormatMedium.format(matchDate)}</p>
        </div>
        <div class="card variant-ghost-primary dark:variant-soft-primary p-3">
            <p class="mb-2 font-medium">Status</p>
            <p>{match.human_state}</p>
        </div>
    </div>

</div>