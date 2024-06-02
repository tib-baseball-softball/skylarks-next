<script lang="ts">
    import type {Table} from "bsm.js";

    export let table: Table
</script>

<div class="table-container">
    <table class="table">
        <thead>
        <tr>
            <th>#</th>
            <th>Team</th>
            <th>W</th>
            <th>L</th>
            <th>Win %</th>
            <th>GB</th>
            <th>Serie</th>
        </tr>
        </thead>

        <tbody>
        {#each table?.rows ?? [] as row}
            <tr>
                <td>{row.rank}</td>
                <td>{row.team_name}</td>
                <td>{row.wins_count}</td>
                <td>{row.losses_count}</td>
                <td>{row.quota}</td>
                <td>{row.games_behind}</td>
                <td>{row.streak}</td>
            </tr>
        {/each}
        </tbody>
    </table>
</div>