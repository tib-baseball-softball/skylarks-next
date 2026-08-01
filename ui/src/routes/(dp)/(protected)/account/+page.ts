import type {
  CustomAuthModel,
  ExpandedAnnouncement,
  ExpandedEvent,
} from "$lib/dp/types/ExpandedResponse.ts";
import { Collection } from "$lib/dp/enum/Collection.ts";
import { authSettings } from "$lib/dp/client.svelte.ts";
import { watchWithPagination } from "$lib/dp/records/RecordOperations.ts";

export const load = async ({ depends, url, fetch }) => {
  const model = authSettings.record as CustomAuthModel;

  const pageQuery = url.searchParams.get("page") ?? "1";
  const page = Number(pageQuery);

  let teamFilter = ""
  if (model.teams) {
    teamFilter += "(";
    model.teams?.forEach((team, index) => {
      teamFilter += `(team = "${team}" || additional_teams.id ?= "${team}")`;
      if (index < model.teams.length - 1) {
        teamFilter += " || ";
      }
    });
    teamFilter += ")";
  }

  let clubFilter = ""
  if (model.club) {
    clubFilter += "(";
    model.club?.forEach((club, index) => {
      clubFilter += `club = "${club}"`;
      if (index < model.club.length - 1) {
        clubFilter += " || ";
      }
    });
    clubFilter += ")";
  }

  // shortened regular auth rule - user dashboard shows only events as member, not admin
  // there will be a dedicated administration dashboard
  let filterString = `starttime >= @todayStart && (${teamFilter} || ${clubFilter})`;

  const eventStore = await watchWithPagination<ExpandedEvent>(
    Collection.Events,
    {
      sort: "+starttime",
      filter: filterString,
      expand:
        "participations_via_event.user, attire, club, team, additional_teams, location",
      fetch: fetch,
      requestKey: `${model?.id}-events`,
    },
    1,
    6,
  );

  const announcements = watchWithPagination<ExpandedAnnouncement>(
    Collection.Announcements,
    {
      sort: "-created",
      fetch: fetch,
      expand: "author,club,team,comments_via_announcement.user",
      requestKey: `${model?.id}-announcements`,
    },
    page,
    3,
  );

  depends("event:list");

  return {
    eventStore: eventStore,
    announcementStore: announcements,
  };
};
