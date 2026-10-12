<script lang="ts">
  import SidebarNavigation from "#lib/dp/components/meta/SidebarNavigation.svelte";
  import type {
    ExpandedClub,
    ExpandedTeam,
  } from "#lib/dp/types/ExpandedResponse.ts";
  import Sheet from "#lib/dp/components/modal/Sheet.svelte";
  import { Menu, X } from "@lucide/svelte";

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
    <Menu />
  {/snippet}

  {#snippet title()}
    <h2 class="sr-only">Navigation</h2>
  {/snippet}

  <div class="navigation-sheet-content">
    <div class="navigation-sheet-header-box">
      <a aria-label="to home page" href="/" onclick={() => (open = false)}>
        <img alt="Diamond Planner Logo" src="/icon_dp.svg" />
      </a>
      <button
        class="btn preset-outlined-card"
        onclick={() => (open = false)}
        aria-label="close navigation drawer"
      >
        <X />
      </button>
    </div>

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

  .navigation-sheet-header-box {
    display: flex;
    justify-content: space-around;
    align-items: center;
    padding: 0.5em;

    img {
      width: calc(var(--spacing) * 14);
    }
  }

  hr {
    margin-block-end: 0.5em;
  }
</style>
