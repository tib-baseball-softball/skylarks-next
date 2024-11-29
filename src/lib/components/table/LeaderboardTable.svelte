<script lang="ts">
  import type {LeaderboardData} from "$lib/model/LeaderboardData";
  import {LeaderboardUtility} from "$lib/service/LeaderboardUtility";

  interface Props {
    data: LeaderboardData
  }

  let {data}: Props = $props();

  const statName = $derived(LeaderboardUtility.getHumanReadableStatName(data.stats_category))
</script>

<section>
    <header class="mb-2">
        <h3 class="h3">{statName}</h3>
    </header>

    <div class="table-container dark:border-2">
        <table class="table table-compact">
            <thead>
            <tr>
                <th class="table-cell-fit">#</th>
                <th>Name</th>
                <th class="table-cell-fit">Value</th>
            </tr>
            </thead>

            <tbody>
            {#each data.data as dataset}
                <tr>
                    <td class="table-cell-fit">{dataset.rank}</td>
                    <td>{dataset.person.first_name} {dataset.person.last_name}
                    </td>
                    <td class="table-cell-fit">{dataset.value}</td>
                </tr>
            {/each}
            </tbody>
        </table>
    </div>
</section>
