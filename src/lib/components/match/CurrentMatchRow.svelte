<script lang="ts">
    import type {Match} from "../../../../../bsm.js/src/model/Match";
    import {LogoUtility} from "$lib/service/LogoUtility";
    import {MatchUtility} from "$lib/service/MatchUtility";
    import {GameWinner} from "$lib/enum/GameWinner";
    import {MatchState} from "$lib/enum/MatchState";
    import {DateTimeUtility} from "$lib/service/DateTimeUtility";

    export let match: Match

    const awayLogo = LogoUtility.getLogoPathForTeamName(match.away_league_entry.team)
    const homeLogo = LogoUtility.getLogoPathForTeamName(match.home_league_entry.team)

    const winner = MatchUtility.getWinnerForMatch(match)
    const matchState = MatchUtility.getMatchState(match)
    const matchDate = DateTimeUtility.parseDateFromBSMString(match.time)
</script>

<a href="gamecenter/game-detail/{match.id}"
   class="grid grid-cols-3 px-3 py-1 gap-x-2 border-b border-surface-400-500-token justify-around items-center">
    <div class="flex justify-end items-center gap-2">
        <div>{match.away_league_entry.team?.short_name}</div>
        <img src="{awayLogo}" alt="team logo for {match.away_team_name}" width="35" loading="lazy"/>
    </div>

    {#if matchState === MatchState.notYetPlayed}

        <div class="flex justify-center items-center">
            <span class="variant-soft-surface dark:variant-filled-surface py-0.5 px-2 rounded">
                {DateTimeUtility.timeFormatShort.format(matchDate)}
            </span>
        </div>

    {:else if matchState === MatchState.cancelled}

        <div class="flex justify-center items-center">
            <span class="variant-soft-surface dark:variant-filled-surface py-0.5 px-2 rounded">
                PPD
            </span>
        </div>

    {:else}

        <div class="font-extrabold flex justify-center text-lg">
            <span class:text-surface-600-300-token={winner === GameWinner.home}>{match.away_runs ?? " "}</span>
            <span>&nbsp;:&nbsp;</span>
            <span class:text-surface-600-300-token={winner === GameWinner.away}>{match.home_runs ?? " "}</span>
        </div>

    {/if}

    <div class="flex justify-start items-center gap-2">
        <img src="{homeLogo}" alt="team logo for {match.away_team_name}" width="35" loading="lazy"/>
        <div>{match.home_league_entry.team?.short_name}</div>
    </div>
</a>