<script lang="ts">
  import TeamMemberDatatable from "$lib/components/datatable/TeamMemberDatatable.svelte";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import TeamAddMemberButton from "$lib/components/team/TeamAddMemberButton.svelte";

  let {data} = $props();

  const players = $derived(data.players);

  const showAdminSection = $derived.by(() => {
    if (!authSettings?.record?.id) {
      return false
    }
    if (data.team?.expand?.club?.admins.includes(authSettings?.record.id)) {
      return true
    }
    return data.team.admins.includes(authSettings?.record?.id);
  });
</script>

<svelte:head>
  <title>Members of {data.team.name}</title>
  <meta name="description" content="Team member list for {data.team.name}."/>
</svelte:head>

<h1 class="h1">Team members for Team "{data.team.name}"</h1>

<TeamMemberDatatable data={$players.items} team={data.team} {showAdminSection}/>

<TeamAddMemberButton team={data.team} club={data.team?.expand?.club}/>