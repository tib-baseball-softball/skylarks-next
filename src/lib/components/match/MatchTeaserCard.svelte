<script lang="ts">
    import type {Match} from "../../../../../bsm.js/src/model/Match";

    export let match: Match
</script>

<a class="block card card-hover variant-ghost-surface" href="gamecenter/game-detail/{match.id}">
    <header class="card-header">
        <h3>{match.league.name}</h3>
        <p>{match.time}</p>
    </header>
    <hr class="mt-2 mx-3"/>
    <section class="p-4">
        <div class="flex justify-between items-center">
            <img src="/berlin_skylarks_logo.svg" alt="team logo" width="50">
            <span>{match.away_team_name}</span>
            <span>{match.away_runs ?? ""}</span>
        </div>
        <div class="flex justify-between items-center">
            <img src="/berlin_skylarks_logo.svg" alt="team logo" width="50">
            <span>{match.home_team_name}</span>
            <span>{match.home_runs ?? ""}</span>
        </div>
    </section>
</a>