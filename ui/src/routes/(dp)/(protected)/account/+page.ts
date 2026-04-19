import type {CustomAuthModel, ExpandedAnnouncement, ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";
import {Collection} from "$lib/dp/enum/Collection.ts";
import {authSettings, client} from "$lib/dp/client.svelte.ts";
import {watchWithPagination} from "$lib/dp/records/RecordOperations.ts";

export const load = (async ({depends, url, fetch}) => {
  const model = authSettings.record as CustomAuthModel;

  const pageQuery = url.searchParams.get("page") ?? "1";
  const page = Number(pageQuery);

  const events = client.collection(Collection.Events).getList<ExpandedEvent>(
    1,
    6,
    {
      sort: "+starttime",
      filter: "starttime >= @todayStart",
      expand: "participations_via_event.user, attire, location",
      fetch: fetch,
      requestKey: `${model?.id}-events`,
    },
  );

  const announcements = await watchWithPagination<ExpandedAnnouncement>(
    Collection.Announcements,
    {
      sort: "-updated",
      fetch: fetch,
      expand: "author,club,team,comments_via_announcement.user",
      requestKey: `${model?.id}-announcements`,
    },
    page,
    3
  );

  depends("event:list");

  return {
    events: events,
    announcementStore: announcements,
  };
});
