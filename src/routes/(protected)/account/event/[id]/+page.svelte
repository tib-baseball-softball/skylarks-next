<script lang="ts">
    import {ProgressRadial} from "@skeletonlabs/skeleton";

    let {data} = $props()
</script>

{#await data.event}
    <ProgressRadial/>
{:then event}
    <h1>{event.title}</h1>

    <p>Desc: {event.desc}</p>
    <p>BSM-ID: {event.bsm_id}</p>
    <p>Start: {event.starttime}</p>
    <p>Treffen: {event.meetingtime}</p>
    <p>Ende: {event.endtime}</p>
    <p>Typ: {event.type}</p>

    <h2>Teilnehmende</h2>

    <ul>
        {#each event.expand?.participants as participant}
            <li>{participant.first_name} {participant.last_name}</li>
        {/each}
    </ul>

{:catch error}
    <p>Fehler beim Laden: {error.message}</p>
{/await}