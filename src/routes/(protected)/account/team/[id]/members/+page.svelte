<script lang="ts">
  import TeamMemberDatatable from "$lib/components/datatable/TeamMemberDatatable.svelte";
  import {authSettings} from "$lib/pocketbase/index.svelte";

  let {data} = $props();

  const players = $derived(data.players);

  const showAdminSection = $derived(
      data.team.admins.includes(authSettings?.record?.id),
  );
</script>

<h1 class="h1">Team members for Team "{data.team.name}"</h1>

<TeamMemberDatatable data={$players.items} {showAdminSection}/>
