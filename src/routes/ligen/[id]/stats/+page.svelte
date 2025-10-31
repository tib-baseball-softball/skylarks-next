<script lang="ts">
  // @ts-ignore
  import {Tabs} from "bits-ui";
  import {StatsType} from "bsm.js";
  import LeaderboardTable from "$lib/components/table/LeaderboardTable.svelte";
  import {goto} from "$app/navigation";
  import ProgressRing from "$lib/components/utility/ProgressRing.svelte";

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

<h1 class="h1 mb-6!">Leaderboards for {data.leagueGroup.name} ({data.leagueGroup.season})</h1>

<Tabs.Root bind:value={type} class="flex">
  <Tabs.List class="tabs-list border mb-1 preset-tonal-surface">
    <Tabs.Trigger class="tabs-trigger btn flex-grow" value={StatsType.batting}>Batting</Tabs.Trigger>
    <Tabs.Trigger class="tabs-trigger btn flex-grow" value={StatsType.pitching}>Pitching</Tabs.Trigger>
    <Tabs.Trigger class="tabs-trigger btn flex-grow" value={StatsType.fielding}>Fielding</Tabs.Trigger>
  </Tabs.List>
</Tabs.Root>

{#await data.leaderboardData}
  <ProgressRing/>
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
