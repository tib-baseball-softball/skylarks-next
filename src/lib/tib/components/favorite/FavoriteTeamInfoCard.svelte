<script lang="ts">
  import type {Team} from "bsm.js";
  import GenericDetailRow from "$lib/dp/components/utils/GenericDetailRow.svelte";
  import {Calendar, Star, TableProperties} from "lucide-svelte";

  interface Props {
    clubTeam: Team;
  }

  let {clubTeam}: Props = $props();
</script>

<article class="card p-3 preset-tonal dark:border dark:border-surface-500 shadow-xl">
  <GenericDetailRow
    categoryName="Name"
    rowValue={`${clubTeam.name} (${clubTeam.short_name})`}
  >
    {#snippet icon()}
      <Star class="fill-primary-500"/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow
    categoryName="League"
    rowValue={clubTeam.league_entries.at(0)?.league.name ?? ""}
  >
    {#snippet icon()}
      <TableProperties/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow
    categoryName="Season"
    rowValue={clubTeam.league_entries.at(0)?.league.season ?? ""}
  >
    {#snippet icon()}
      <Calendar/>
    {/snippet}
  </GenericDetailRow>
</article>
