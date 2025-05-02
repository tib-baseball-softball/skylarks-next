<script lang="ts">
  import {Tabs} from "@skeletonlabs/skeleton-svelte";
  import type {Match} from "bsm.js";
  import CurrentMatchBlock from "$lib/components/match/CurrentMatchBlock.svelte";

  let tabSet: "previous" | "current" | "next" | string = $state("current");

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
    <Tabs value={tabSet} onValueChange={(e) => (tabSet = e.value)} listJustify="justify-center" fluid={true}>

      {#snippet list()}
        <Tabs.Control value="previous">Previous Gameday</Tabs.Control>
        <Tabs.Control value="current">Current Gameday</Tabs.Control>
        <Tabs.Control value="next">Next Gameday</Tabs.Control>
      {/snippet}

      {#snippet content()}
        
        <Tabs.Panel value="previous">
          <CurrentMatchBlock matches={matchesPrevious}/>
        </Tabs.Panel>
        <Tabs.Panel value="current">
          <CurrentMatchBlock matches={matchesCurrent}/>
        </Tabs.Panel>
        <Tabs.Panel value="next">
          <CurrentMatchBlock matches={matchesNext}/>
        </Tabs.Panel>
      {/snippet}
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
