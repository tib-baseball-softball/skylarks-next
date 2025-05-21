<script lang="ts">
    import type {Table} from "bsm.js";
    import {env} from "$env/dynamic/public";

    interface Props {
    table: Table;
  }

  let {table}: Props = $props();
</script>

<div class="table-wrap">
    <table class="table border">
        <thead class="dark:text-surface-contrast-500!">
        <tr class="preset-ghost-surface">
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
            <tr class:text-primary-600-400={row.team_name.includes(env.PUBLIC_TEAM_NAME)}>
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