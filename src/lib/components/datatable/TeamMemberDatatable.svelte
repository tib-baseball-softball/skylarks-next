<script lang="ts">
  import {DataHandler} from "@vincjo/datatables";
  import Pagination from "$lib/components/datatable/Pagination.svelte";
  import RowCount from "$lib/components/datatable/RowCount.svelte";
  import RowsPerPage from "$lib/components/datatable/RowsPerPage.svelte";
  import Search from "$lib/components/datatable/Search.svelte";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import TeamMembersTableContent from "./TeamMembersTableContent.svelte";

  interface Props {
    data: CustomAuthModel[];
    team: ExpandedTeam;
    rowsPerPage?: number;
    showAdminSection?: boolean;
  }

  let {data, team, rowsPerPage = 25, showAdminSection = false}: Props = $props();

  const handler = $derived(
      new DataHandler<CustomAuthModel>(data, {
        rowsPerPage: rowsPerPage,
      })
  );
</script>

<div class=" overflow-x-auto space-y-4 table-wrap">
  <!-- Header -->
  <header class="md:flex space-y-2 md:space-y-0 justify-between gap-4">
    <Search {handler}/>

    <RowsPerPage {handler}/>
  </header>

  <!-- Table -->
  <table class="table table-compact w-full">
    <TeamMembersTableContent {handler} {showAdminSection} {team}/>

    <tfoot></tfoot>
  </table>

  <!-- Footer -->
  <footer class="flex justify-between">
    <RowCount {handler}/>
    <Pagination {handler}/>
  </footer>
</div>
