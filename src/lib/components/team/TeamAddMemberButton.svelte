<script lang="ts">
  import type {ExpandedTeam} from "$lib/model/ExpandedResponse.ts";
  import type {ClubsResponse} from "$lib/model/pb-types.ts";
  import {PlusOutline} from "flowbite-svelte-icons";
  import ProjectDialog from "$lib/components/utility/ProjectDialog.svelte";
  import TeamMembersForm from "$lib/components/forms/TeamMembersForm.svelte";

  interface Props {
    club: ClubsResponse,
    team: ExpandedTeam,
  }

  let {club, team}: Props = $props()
</script>

<ProjectDialog buttonText="Add Members">

  {#snippet triggerIcon()}
    <PlusOutline/>
  {/snippet}

  {#snippet title()}
    Add Members to team "{team.name}"
  {/snippet}

  {#snippet description()}
    <div class="text-sm">
      Select all {club.name} members that should be added to the team.
    </div>
  {/snippet}

  <TeamMembersForm {team} {club}/>
</ProjectDialog>