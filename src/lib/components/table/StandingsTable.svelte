<script lang="ts">
    import type {Table} from "bsm.js";
    import {PUBLIC_TEAM_NAME} from "$env/static/public";

    interface Props {
        table: Table;
    }

    let { table }: Props = $props();
</script>

<div class="table-container dark:border-2">
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
            <tr class:text-primary-500-400-token={row.team_name.includes(PUBLIC_TEAM_NAME)}>
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