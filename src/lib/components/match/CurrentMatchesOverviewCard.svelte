<script lang="ts">
  import {Tab, TabGroup} from "@skeletonlabs/skeleton";
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
        class="card overview-card variant-soft-surface dark:border dark:border-tertiary-500-400-token"
>
  <header class="card-header">
    <h2 class="h3">Current Games</h2>
  </header>
  <section class="p-4">
    <TabGroup justify="justify-center" flex="flex-auto">
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
    </TabGroup>
  </section>
  <footer class="card-footer flex justify-end">
    <a
            href="/gamecenter"
            class="btn variant-filled-primary dark:variant-ghost-primary px-10"
    >Gamecenter</a
    >
  </footer>
</div>
