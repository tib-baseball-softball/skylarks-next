<script lang="ts">
    import type {AdditionalStat} from "bsm.js";

    export let entries: AdditionalStat[]
</script>

<span class="entries">
    {#each entries as entry}
        <span>{entry.person.last_name} {entry.person.first_name.charAt(0)}.
            {#if entry.sum > 1}
                ({entry.sum})
            {/if}
        </span>
    {/each}
</span>

<style>
    .entries span:not(:last-child)::after {
        content: ", ";
    }
</style>