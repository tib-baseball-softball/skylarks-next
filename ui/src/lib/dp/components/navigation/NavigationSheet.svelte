<script lang="ts">
  import SidebarNavigation from "$lib/dp/components/meta/SidebarNavigation.svelte";
  import type {
    ExpandedClub,
    ExpandedTeam,
  } from "$lib/dp/types/ExpandedResponse.ts";
  import Sheet from "../modal/Sheet.svelte";

  interface Props {
    clubs: ExpandedClub[];
    teams: ExpandedTeam[];
  }

  let { clubs, teams }: Props = $props();

  let open = $state(false);
</script>

<Sheet
  bind:open
  triggerClasses="nav-trigger btn-sm"
  side="left"
  closeButtonClasses="sr-only"
>
  {#snippet triggerContent()}
    <svg viewBox="0 0 100 80" aria-hidden="true">
      <rect height="20" width="100" />
      <rect height="20" width="100" y="30" />
      <rect height="20" width="100" y="60" />
    </svg>
  {/snippet}

  {#snippet title()}
    <h2 class="sr-only">Navigation</h2>
  {/snippet}

  <div class="navigation-sheet-content">
    <a
      aria-label="to home page"
      class="anchor"
      href="/"
      onclick={() => (open = false)}
    >
      <img alt="Diamond Planner Logo" src="/icon_dp.svg" />
    </a>

    <hr />

    <SidebarNavigation bind:sheetOpen={open} {clubs} {teams} />
  </div>
</Sheet>

<style>
  :global {
    .nav-trigger {
      margin-inline-end: 2em;
      @media (min-width: 48rem) {
        display: none;
      }
    }
  }

  a {
    display: flex;
    justify-content: center;
    padding: 0.5em;

    img {
      max-width: calc(var(--spacing) * 14);
    }
  }

  hr {
    margin-block-end: 0.5em;
  }

  svg {
    --svg-size: calc(var(--spacing) * 4);
    width: var(--svg-size);
    height: var(--svg-size);
    fill: var(--color-surface-950-50);
  }
</style>
