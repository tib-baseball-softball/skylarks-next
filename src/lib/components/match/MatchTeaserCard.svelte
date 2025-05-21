<script lang="ts">
    import type { Match } from "bsm.js";
    import { LogoUtility } from "$lib/service/LogoUtility";
    import MatchTeaserCardRow from "$lib/components/match/MatchTeaserCardRow.svelte";
    import { GameWinner } from "$lib/enum/GameWinner";
    import { DateTimeUtility } from "$lib/service/DateTimeUtility";
    import { MatchDecorator } from "$lib/service/MatchDecorator";
    import GameResultIndicator from "$lib/components/match/GameResultIndicator.svelte";

    interface Props {
        match: Match;
    }

    let { match }: Props = $props();

    const awayLogo = LogoUtility.getLogoPathForTeamName(
        match?.away_league_entry?.team,
    );
    const homeLogo = LogoUtility.getLogoPathForTeamName(
        match?.home_league_entry?.team,
    );

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
            <GameResultIndicator {match} />
        </div>
    </header>

    <hr class="mt-2 mx-3" />

    <section class="p-4 flex flex-col gap-2">
        <MatchTeaserCardRow
            logo={awayLogo}
            teamName={match.away_team_name}
            score={match.away_runs}
            subduedScore={winner === GameWinner.home}
        />
        <MatchTeaserCardRow
            logo={homeLogo}
            teamName={match.home_team_name}
            score={match.home_runs}
            subduedScore={winner === GameWinner.away}
        />
    </section>
</a>
