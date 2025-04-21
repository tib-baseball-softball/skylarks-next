<script lang="ts">
  import {Tab, Tabs } from "@skeletonlabs/skeleton-svelte";
  import type {Match} from "bsm.js";
  import CurrentMatchBlock from "$lib/components/match/CurrentMatchBlock.svelte";

  let tabSet: number = $state(1);

  interface Props {
    matchesCurrent: Promise<Match[]>;
    matchesPrevious: Promise<Match[]>;
    matchesNext: Promise<Match[]>;
  }

  let {matchesCurrent, matchesPrevious, matchesNext}: Props = $props();
</script>

<div
        class="card overview-card preset-tonal-surface dark:border dark:border-tertiary-600-400"
>
  <header class="card-header">
    <h2 class="h3">Current Games</h2>
  </header>
  <section class="p-4">
    <Tabs justify="justify-center" flex="flex-auto">
      <Tab bind:group={tabSet} name="tab1" value={0}
      >Previous Gameday
      </Tab
      >
      <Tab bind:group={tabSet} name="tab2" value={1}
      >Current Gameday
      </Tab
      >
      <Tab bind:group={tabSet} name="tab3" value={2}
      >Next Gameday
      </Tab
      >
      <!-- Tab Panels --->
      <svelte:fragment slot="panel">
        {#if tabSet === 0}
          <CurrentMatchBlock matches={matchesPrevious}/>
        {:else if tabSet === 1}
          <CurrentMatchBlock matches={matchesCurrent}/>
        {:else if tabSet === 2}
          <CurrentMatchBlock matches={matchesNext}/>
        {/if}
      </svelte:fragment>
    </Tabs>
  </section>
  <footer class="card-footer flex justify-end">
    <a
            href="/gamecenter"
            class="btn preset-filled-primary-500 dark:preset-tonal-primary border border-primary-500 px-10"
    >Gamecenter</a
    >
  </footer>
</div>
