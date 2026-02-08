<script lang="ts">
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import JoinTeamForm from "$lib/dp/components/forms/JoinTeamForm.svelte";

  interface Props {
    authRecord: CustomAuthModel;
    teamID: string;
  }

  let {authRecord, teamID}: Props = $props();
</script>

<p class="prose">
  You are not a member of this team, so you can only see basic information.
</p>
<p class="prose">
  If you have the team's signup key, please enter it here to join:
</p>
<div class="join-grid">
  <article class="card join-container preset-tonal-surface">
    <h3 class="h4">
      Join Team
    </h3>
    <div>
      <JoinTeamForm {authRecord} {teamID}/>
    </div>
  </article>
</div>

<style>
  .join-grid {
    display: grid;
    grid-template-columns: 1fr;
    margin-block-start: calc(var(--spacing) * 4);

    @media (min-width: 64rem) {
      grid-template-columns: 1fr 1fr;
    }
  }

  .join-container {
    padding: calc(var(--spacing) * 4);

    h3 {
      margin-block-end: calc(var(--spacing) * 4);
    }
  }
</style>
