<script lang="ts">
  import SidebarNavigation from "$lib/components/meta/SidebarNavigation.svelte";
  import type {ExpandedClub, ExpandedTeam} from "$lib/model/ExpandedResponse.ts";
  //@ts-ignore
  import * as Sheet from "$lib/components/utility/sheet/index.js";

  interface Props {
    clubs: ExpandedClub[],
    teams: ExpandedTeam[],
  }

  let {clubs, teams}: Props = $props();

  let open = $state(false);
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger aria-label="open navigation" class="md:hidden btn btn-sm mr-4">
      <span>
          <svg viewBox="0 0 100 80" class="fill-token w-4 h-4">
              <rect width="100" height="20"/>
              <rect y="30" width="100" height="20"/>
              <rect y="60" width="100" height="20"/>
          </svg>
      </span>
  </Sheet.Trigger>

  <Sheet.Content side="left" class="w-[70%]! !sm:w-[40%] p-0! navigation-sheet-content">
    <Sheet.Header></Sheet.Header>

    <a href="/" aria-label="to home page" class="flex justify-around p-2" onclick={() => open = false}>
      <img class="max-w-14" src="/berlin_skylarks_logo.svg" alt="Skylarks Team Logo">

      <h2 class="p-4 antialiased">Berlin Skylarks</h2>
    </a>

    <hr class="mb-2"/>

    <SidebarNavigation clubs={clubs} teams={teams} bind:sheetOpen={open}/>

  </Sheet.Content>
</Sheet.Root>

<style>
    :global {
        .navigation-sheet-content {
            .sheet-close-button {
                display: none;
            }
        }
    }
</style>