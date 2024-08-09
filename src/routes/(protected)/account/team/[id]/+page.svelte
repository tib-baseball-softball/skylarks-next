<script lang="ts">
    import {ProgressRadial} from "@skeletonlabs/skeleton";

    let {data} = $props()
</script>

<h1 class="h2">{data.team.name}</h1>

<p>BSM-Liga: {data.team.bsm_league_group}</p>

<h2 class="h3">Aktuelle Events</h2>
{#await data.events}
    <ProgressRadial/>
{:then events}
    <ul>
        {#each events.items as event}
        <li>
            <a class="anchor" href="/account/event/{event.id}">{event.title}</a>
        </li>
        {/each}
    </ul>
{:catch error}
    <p>Fehler beim Laden: {error.message}</p>
{/await}