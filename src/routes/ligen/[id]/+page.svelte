<script lang="ts">
  import {ProgressRadial} from "@skeletonlabs/skeleton";
  import StandingsTable from "$lib/components/table/StandingsTable.svelte";
  import LeagueDetailInfoCard from "$lib/components/league/LeagueDetailInfoCard.svelte";
  import type {PageProps} from "./$types";

  let {data}: PageProps = $props();
  let expectedTable = $derived(data.table);
  let leagueGroup = $derived(data.leagueGroup);
</script>

<h1 class="h1 my-4">{leagueGroup.name} ({leagueGroup.season})</h1>

<section class="my-5">
  <h2 class="h2">Information</h2>
  <div class="grid grid-cols-1 md:grid-cols-2 gap-5">

    <LeagueDetailInfoCard {leagueGroup}/>

    <article class="card variant-soft-surface">
      <header class="card-header">
        <h3 class="h3">League Stats</h3>
      </header>

      <div class="flex flex-col justify-between items-end h-[80%]">

        <p class="p-4">
          Statistics about current league leaders in various statistical categories can be found on a
          dedicated subpage.
        </p>

        <footer class="card-footer mt-2">
          <a href="/ligen/{leagueGroup.id}/stats" class="btn variant-ghost-primary">Go to stats</a>
        </footer>
      </div>
    </article>

  </div>
</section>

<section class="my-3">

  <h2 class="h2">Tabelle</h2>
  {#await expectedTable}
    <ProgressRadial/>
  {:then table}
    {#if table}
      <StandingsTable {table}/>
    {/if}
  {:catch error}
    <p>error loading: {error.message}</p>
  {/await}

</section>

<style lang="postcss">
    h2 {
        @apply mb-3;
    }
</style>