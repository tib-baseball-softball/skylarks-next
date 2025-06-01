<script lang="ts">
  // @ts-ignore
  import {Tabs} from "bits-ui";
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

{#snippet gameDayWord()}
  <span class="hidden lg:inline">Gameday</span>
{/snippet}

<div
        class="card overview-card preset-tonal-surface dark:border dark:border-tertiary-600-400"
>
  <header class="card-header">
    <h2 class="h3">Current Games</h2>
  </header>
  <section class="p-4">
    <Tabs.Root
            bind:value={tabSet}
            class=""
    >
      <Tabs.List
              class="tabs-list border mb-1 preset-tonal-surface"
      >
        <Tabs.Trigger
                value="previous"
                class="tabs-trigger"
        >
          Previous
          {@render gameDayWord()}
        </Tabs.Trigger>
        <Tabs.Trigger
                value="current"
                class="tabs-trigger"
        >Current
          {@render gameDayWord()}
        </Tabs.Trigger>
        <Tabs.Trigger
                value="next"
                class="tabs-trigger"
        >Next
          {@render gameDayWord()}
        </Tabs.Trigger>
      </Tabs.List>

      <Tabs.Content value="previous" class="pt-4">
        <CurrentMatchBlock matches={matchesPrevious}/>
      </Tabs.Content>
      <Tabs.Content value="current" class="pt-4">
        <CurrentMatchBlock matches={matchesCurrent}/>
      </Tabs.Content>
      <Tabs.Content value="next" class="pt-4">
        <CurrentMatchBlock matches={matchesNext}/>
      </Tabs.Content>
    </Tabs.Root>
  </section>
  <footer class="card-footer flex justify-end">
    <a
            href="/gamecenter"
            class="btn preset-filled-primary-500 dark:preset-tonal-primary border border-primary-500 px-10"
    >Gamecenter</a
    >
  </footer>
</div>
