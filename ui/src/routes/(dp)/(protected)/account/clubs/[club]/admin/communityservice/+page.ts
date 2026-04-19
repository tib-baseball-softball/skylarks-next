import type {PageLoad} from './$types';
import {client} from "$lib/dp/client.svelte.ts";
import type {ClubCommunityServiceRow} from "$lib/dp/types/ClubCommunityServiceRow.ts";
import { Collection } from '$lib/dp/enum/Collection';
import { watchWithPagination } from '$lib/dp/records/RecordOperations';
import type { ExpandedServiceEntry } from '$lib/dp/types/ExpandedResponse';

export const load: PageLoad = (async ({fetch, params, url, depends, parent}) => {
  type Query = {
    season?: string;
  }
  const query: Query = {};

  let season = url.searchParams.get("season");
  if (!season) {
    season = new Date().getFullYear().toString();
  }
  query.season = season;

  const rows = client.send<ClubCommunityServiceRow[]>(`/api/clubs/${params.club}/admin/communityservice`, {
    fetch: fetch,
    query: query,
  });
  
  const pageNumber = Number(url.searchParams.get("page")) ?? 1;
  
  const entries = await client.collection(Collection.ServiceEntries).getFullList<ExpandedServiceEntry>(
    {
      filter: `club = "${params.club}" && strftime('%Y', service_date) = "${season}"`,
      expand: "member, board_member, club",
      sort: "+service_date",
      fetch: fetch,
    },
  )
  
  const parentData = await parent();

  depends("communityservice:admin");

  return {
    rows: rows,
    entries: entries,
    club: parentData.club,
  };
});
