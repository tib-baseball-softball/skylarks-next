<script lang="ts">
  import GenericDetailRow from "$lib/components/utility/GenericDetailRow.svelte";
  import {Crown, Hash, Percent, Sigma, Signpost} from "lucide-svelte";
  //@ts-expect-error
  import type {TableRow} from "bsm.js/dist/model/Table";
  import type {LeagueGroup} from "bsm.js";

  interface Props {
    leagueGroup: LeagueGroup;
    tableRow: TableRow;
  }

  let {leagueGroup, tableRow}: Props = $props();
</script>

<article class="card p-3 preset-tonal dark:border dark:border-surface-500 shadow-xl">
  <GenericDetailRow
          categoryName="Acronym"
          rowValue={leagueGroup.acronym}
  >
    {#snippet icon()}
      <Signpost/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow
          categoryName="Record"
          rowValue={`${tableRow.wins_count} - ${tableRow.losses_count}`}
  >
    {#snippet icon()}
      <Sigma/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow
          categoryName="Pct."
          rowValue={tableRow.quota}
  >
    {#snippet icon()}
      <Percent/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow
          categoryName="Rank"
          rowValue={tableRow.rank}
  >
    {#snippet icon()}
      {#if tableRow.rank === "1."}
        <Crown class="stroke-primary-500"/>
      {:else }
        <Hash/>
      {/if}
    {/snippet}
  </GenericDetailRow>

</article>
