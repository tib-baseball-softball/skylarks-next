<script lang="ts">
  import type {ClubCommunityServiceRow} from "$lib/dp/types/ClubCommunityServiceRow.ts";
  import type {ExpandedClub, ExpandedEventSeries, ExpandedServiceEntry} from "$lib/dp/types/ExpandedResponse.ts";
  import CommunityServiceDatatable from "$lib/dp/components/communityservice/CommunityServiceDatatable.svelte";
  import CommunityServiceDatatableRaw from "./CommunityServiceDatatableRaw.svelte";

  interface Props {
    rows: ClubCommunityServiceRow[];
    rawEntries: ExpandedServiceEntry[];
    club: ExpandedClub;
  }

  let {rows,rawEntries, club}: Props = $props();
</script>

<div class="wrapper">
  <h2 class="h2">Aggregated Data (per user)</h2>
  <CommunityServiceDatatable data={rows} {club}/>
</div>

<div class="wrapper">
  <h2 class="h2">Raw data (single entries)</h2>
  <CommunityServiceDatatableRaw data={rawEntries} {club}/>
</div>

<style>
  .wrapper {
    margin-block: calc(var(--spacing) * 6);
  }
  
  .h2 {
    margin-block-end: calc(var(--spacing) * 4);
  }
</style>
