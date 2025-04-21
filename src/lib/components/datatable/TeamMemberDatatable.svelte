<script lang="ts">
  import Search from "$lib/components/datatable/Search.svelte";
  import RowsPerPage from "$lib/components/datatable/RowsPerPage.svelte";
  import RowCount from "$lib/components/datatable/RowCount.svelte";
  import Pagination from "$lib/components/datatable/Pagination.svelte";
  import {DataHandler} from "@vincjo/datatables";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/model/ExpandedResponse";
  import TeamMembersTableContent from "./TeamMembersTableContent.svelte";

  interface Props {
    data: CustomAuthModel[];
    team: ExpandedTeam,
    rowsPerPage?: number;
    showAdminSection?: boolean;
  }

  let {data, team, rowsPerPage = 25, showAdminSection = false}: Props = $props();

  const handler = $derived(
      new DataHandler<CustomAuthModel>(data, {
        rowsPerPage: rowsPerPage,
      }),
  );
</script>

<div class=" overflow-x-auto space-y-4 table-container">
  <!-- Header -->
  <header class="flex justify-between gap-4">
    <Search {handler}/>

    <RowsPerPage {handler}/>
  </header>

  <!-- Table -->
  <table class="table  table-compact w-full table-auto">
    <TeamMembersTableContent {handler} {team} {showAdminSection}/>

    <tfoot></tfoot>
  </table>

  <!-- Footer -->
  <footer class="flex justify-between">
    <RowCount {handler}/>
    <Pagination {handler}/>
  </footer>
</div>
