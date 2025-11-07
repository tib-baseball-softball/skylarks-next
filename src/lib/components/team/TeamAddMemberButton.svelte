<script lang="ts">
import type { ExpandedTeam } from "$lib/model/ExpandedResponse.ts"
import type { ClubsResponse } from "$lib/model/pb-types.ts"
import Dialog from "$lib/components/utility/Dialog.svelte"
import TeamMembersForm from "$lib/components/forms/TeamMembersForm.svelte"
import { Plus } from "lucide-svelte"

interface Props {
  club: ClubsResponse
  team: ExpandedTeam
}

let { club, team }: Props = $props()
</script>

<Dialog triggerClasses="btn preset-filled-primary-500">

  {#snippet triggerContent()}
    <Plus/>
    <span>Add Members</span>
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
</Dialog>