<script lang="ts">
    import TeamMemberDatatable from "$lib/components/datatable/TeamMemberDatatable.svelte";

    let { data } = $props();

    const players = $derived(data.players);
</script>

<h1 class="h1">Team members for Team "{data.team.name}"</h1>

<TeamMemberDatatable data={$players.items} />
