<script lang="ts">
  import TeamMemberDatatable from "$lib/components/datatable/TeamMemberDatatable.svelte";
  import {authModel} from "$lib/pocketbase/Auth.svelte";

  let {data} = $props();

  const players = $derived(data.players);

  const showAdminSection = $derived(
      data.team.admins.includes($authModel?.id),
  );
</script>

<h1 class="h1">Team members for Team "{data.team.name}"</h1>

<TeamMemberDatatable data={$players.items} {showAdminSection}/>
