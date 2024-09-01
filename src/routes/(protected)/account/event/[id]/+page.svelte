<script lang="ts">
    import EventParticipationSection from "$lib/components/diamondplanner/event/EventParticipationSection.svelte";

    let { data } = $props();
</script>

{#await data.event then event}
    <h1>{event.title}</h1>

    <p>Desc: {event.desc}</p>
    <p>BSM-ID: {event.bsm_id}</p>
    <p>Start: {event.starttime}</p>
    <p>Treffen: {event.meetingtime}</p>
    <p>Ende: {event.endtime}</p>
    <p>Typ: {event.type}</p>

    <h2>Teilnehmende</h2>

    <EventParticipationSection {event} />
{:catch error}
    <p>Fehler beim Laden: {error.message}</p>
{/await}
