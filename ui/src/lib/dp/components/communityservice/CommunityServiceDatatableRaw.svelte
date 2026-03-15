<script lang="ts">
  import {TableHandler} from "@vincjo/datatables";
  import type {ExpandedClub, ExpandedServiceEntry} from "$lib/dp/types/ExpandedResponse.ts";
  import GenericDatatable from "$lib/dp/components/datatable/GenericDatatable.svelte";
  import type {FilterSortOptions} from "$lib/dp/types/FilterSortOptions.ts";
  import CommunityServiceTableContentRaw from "./CommunityServiceTableContentRaw.svelte";

  interface Props {
    data: ExpandedServiceEntry[];
    club: ExpandedClub;
    rowsPerPage?: number;
  }

  const filterSortOptions: FilterSortOptions = {
    service_date: {displayName: "Date", value: ""},
    name: {displayName: "Name", value: ""},
    minutes: {displayName: "Amount", value: ""},
    title: {displayName: "Title", value: ""},
    description: {displayName: "Description", value: ""},
    admin_comments: {displayName: "Admin Comment", value: ""},
    board_member: {displayName: "Board Member", value: ""},
  };

  let {
    data,
    club,
    rowsPerPage = 75,
  }: Props = $props();

  const handler = $derived(
    new TableHandler<ExpandedServiceEntry>(data, {
      rowsPerPage: rowsPerPage,
    })
  );
</script>

<GenericDatatable
  {handler}
  {filterSortOptions}
  showInHeader="all"
  caption="Community Service Data for {club.name}"
>
  {#snippet additionalSort()}
    <th>Actions</th>
  {/snippet}
  
  {#snippet additionalFilters()}
    <th></th>
  {/snippet}
  
  {#snippet contentRows()}
    <CommunityServiceTableContentRaw {handler}/>
  {/snippet}
</GenericDatatable>
