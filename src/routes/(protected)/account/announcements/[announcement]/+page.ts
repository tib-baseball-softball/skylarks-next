import type {ExpandedAnnouncement} from '$lib/model/ExpandedResponse';
import {watchSingleRecord} from '$lib/pocketbase/RecordOperations';
import type {PageLoad} from './$types';

export const load = (async ({ params, fetch, depends }) => {
    const announcement = await watchSingleRecord<ExpandedAnnouncement>("announcements", params.announcement, {
        fetch: fetch,
        expand: "author,club,team,comments_via_announcement.user"
    });
    depends("announcement:single", "comments:list");

    return {
        announcement: announcement
    };
}) satisfies PageLoad;