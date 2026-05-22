import type {CustomAuthModel, ExpandedAnnouncement, ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";
import {Collection} from "$lib/dp/enum/Collection.ts";
import {authSettings} from "$lib/dp/client.svelte.ts";
import {watchWithPagination} from "$lib/dp/records/RecordOperations.ts";

export const load = (async ({depends, url, fetch}) => {
  const model = authSettings.record as CustomAuthModel;

  const pageQuery = url.searchParams.get("page") ?? "1";
  const page = Number(pageQuery);

  let filterString = "starttime >= @todayStart";

  if (model?.teams) {
    filterString += " && (";
    model.teams?.forEach((team, index) => {
      filterString += `(team = "${team}")`;
      if (index < model.teams.length - 1) {
        filterString += " || ";
      }
    });
    filterString += ")";
  }


  const eventStore = await watchWithPagination<ExpandedEvent>(
    Collection.Events,
    {
      sort: "+starttime",
      filter: filterString,
      expand: "participations_via_event.user, attire, location",
      fetch: fetch,
      requestKey: `${model?.id}-events`,
    },
    1,
    6,
  );

  const announcements = watchWithPagination<ExpandedAnnouncement>(
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
    eventStore: eventStore,
    announcementStore: announcements,
  };
});
