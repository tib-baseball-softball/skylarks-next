<script lang="ts">
  import type {ClubFunction} from "bsm.js";
  import ClubOfficialRow from "$lib/components/club/officials/ClubOfficialRow.svelte";

  interface Props {
    clubOfficials: ClubFunction[];
  }

  let {clubOfficials = []}: Props = $props();
</script>

<article class="mb-4!">
  <section>
    <ul class="card p-3 preset-tonal dark:border dark:border-surface-500 shadow-xl">

      {#each clubOfficials as clubOfficial, index}
        <li class="p-0.5">
          <a href="/club/officials/{clubOfficial.id}">
            <ClubOfficialRow {clubOfficial}/>
          </a>
        </li>

        {#if index + 1 < clubOfficials.length}
          <hr class="my-2">
        {/if}
      {:else}
        <li>
          <p>No club officials available.</p>
        </li>
      {/each}
    </ul>
  </section>
</article>