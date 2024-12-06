<script lang="ts">
  import {getDrawerStore} from "@skeletonlabs/skeleton";
  import {CloseOutline} from "flowbite-svelte-icons";
  import type {EventseriesResponse} from "$lib/model/pb-types.ts";
  import type {ExpandedTeam} from "$lib/model/ExpandedResponse.ts";

  const drawerStore = getDrawerStore();

  const eventSeries: EventseriesResponse[] = $derived($drawerStore.meta.eventSeries ?? [])
  const team: ExpandedTeam = $derived($drawerStore.meta.team)
</script>

<article class="p-6 space-y-3">
    <div class="flex items-center gap-5">
        <button
                aria-label="cancel and close"
                class="btn variant-ghost-surface"
                onclick={drawerStore.close}
        >
            <CloseOutline/>
        </button>
        <header class="text-xl font-semibold">
            <h2 class="h3">Manage Event Series for {team?.name}</h2>
        </header>
    </div>

    <h3 class="h3 !mt-6">Active event series</h3>
    {#each eventSeries as series}
        <div>{series.title}</div>
    {/each}
</article>
