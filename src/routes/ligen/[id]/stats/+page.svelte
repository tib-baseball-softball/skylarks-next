<script lang="ts">
  import {Progress, Segment } from "@skeletonlabs/skeleton-svelte";
  import {StatsType} from "bsm.js";
  import LeaderboardTable from "$lib/components/table/LeaderboardTable.svelte";
  import {goto} from "$app/navigation";

  let {data} = $props();

  let type: StatsType = $state(StatsType.batting);

  const reloadWithQuery = () => {
    let queryString = `?statsType=${type.toString()}`;

    goto(queryString, {
      noScroll: true,
    });
  };

  $effect.pre(() => {
    console.log(type);
    reloadWithQuery();
  });
</script>

<h1 class="h1">Leaderboards for {data.leagueGroup.name} ({data.leagueGroup.season})</h1>

<Segment name="stats-type" value={type} onValueChange={(e) => (type = e.value ?? StatsType.batting)} display="flex">
    <Segment.Item value={StatsType.batting} classes="flex-grow">Batting</Segment.Item>
    <Segment.Item value={StatsType.pitching} classes="flex-grow">Pitching</Segment.Item>
    <Segment.Item value={StatsType.fielding} classes="flex-grow">Fielding</Segment.Item>
</Segment>

{#await data.leaderboardData}
    <Progress/>
    <div class="placeholder col-span-2"></div>
{:then leaderboardData}
    <header class="space-y-3">
        <h2 class="h2">{leaderboardData.stats_type}</h2>
    </header>

    <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6 xl:gap-8 2xl:gap-10">
        {#each leaderboardData.data as data}
            <LeaderboardTable {data}/>
        {/each}
    </div>
{:catch error}
    <p class="col-span-2">error loading matches: {error.message}</p>
{/await}
