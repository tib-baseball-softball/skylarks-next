<script lang="ts">
    import type {Match} from "../../../../../bsm.js/src/model/Match";
    import {LogoUtility} from "$lib/service/LogoUtility";
    import {MatchUtility} from "$lib/service/MatchUtility";
    import {GameWinner} from "$lib/enum/GameWinner";

    export let match: Match

    const awayLogo = LogoUtility.getLogoPathForTeamName(match.away_team_name)
    const homeLogo = LogoUtility.getLogoPathForTeamName(match.home_team_name)

    const winner = MatchUtility.getWinnerForMatch(match)
</script>

<a href="gamecenter/game-detail/{match.id}"
   class="grid grid-cols-3 px-3 py-1 gap-x-2 border-b border-surface-400-500-token justify-around items-center">
    <div class="flex justify-end items-center gap-2">
        <div>{match.away_league_entry.team?.short_name}</div>
        <img src="{awayLogo}" alt="team logo for {match.away_team_name}" width="35"/>
    </div>

    <div class="font-extrabold flex justify-center text-lg">
        <span class:text-surface-600-300-token={winner === GameWinner.home}>{match.away_runs ?? " "}</span>
        <span>&nbsp;:&nbsp;</span>
        <span class:text-surface-600-300-token={winner === GameWinner.away}>{match.home_runs ?? " "}</span>
    </div>

    <div class="flex justify-start items-center gap-2">
        <img src="{homeLogo}" alt="team logo for {match.away_team_name}" width="35"/>
        <div>{match.home_league_entry.team?.short_name}</div>
    </div>
</a>