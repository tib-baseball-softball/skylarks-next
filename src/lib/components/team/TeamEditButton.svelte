<script lang="ts">
    import {type DrawerSettings, getDrawerStore} from "@skeletonlabs/skeleton";
    import type {ExpandedTeam} from "$lib/model/ExpandedResponse";
    import type {ClubsResponse} from "$lib/model/pb-types";
    import {Edit} from "lucide-svelte";

    interface Props {
    club: ClubsResponse,
    team: ExpandedTeam,
    classes: string,
    showLabel?: boolean,
  }

  let {club, team, classes, showLabel = true}: Props = $props();

  const drawerStore = getDrawerStore();

  const teamSettings: DrawerSettings = $derived({
    id: "team-form",
    position: "right",
    width: "w-[100%] sm:w-[80%] lg:w-[70%] xl:w-[50%]",
    meta: {
      club: club,
      team: team,
    },
  });
</script>

<button
        class="btn {classes}"
        onclick={() => drawerStore.open(teamSettings)}
        aria-label="edit team {team.name}"
>
  <Edit/>
  {#if showLabel}
    <span>Edit Team</span>
  {/if}
</button>