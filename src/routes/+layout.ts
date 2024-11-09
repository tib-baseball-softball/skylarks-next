import type {LayoutLoad} from "../../.svelte-kit/types/src/routes/$types";
import type {CustomAuthModel, ExpandedClub, ExpandedTeam} from "$lib/model/ExpandedResponse";
import {browser} from "$app/environment";
import {client} from "$lib/pocketbase";
import {authModel} from "$lib/pocketbase/Auth.svelte";
import {get} from "svelte/store";

/**
 * This loads club and team data for the extended DiamondPlanner nav bar.
 *
 * Should be executed only client-side and only if logged in.
 */
export const load = (async ({fetch, depends}) => {
  let clubs: ExpandedClub[] = []
  let teams: ExpandedTeam[] = []

  if (browser && client.authStore.isValid) {
    /**
     * setting requestKey to `null` means "do not autocancel" => nav is always up to date
     */
    teams = await client
        .collection("teams")
        .getFullList<ExpandedTeam>({
          expand: "club",
          fetch: fetch,
          requestKey: null,
          sort: "+name",
        });

    const model = get(authModel) as CustomAuthModel
    clubs = await client
        .collection("clubs")
        .getFullList<ExpandedClub>({
          filter: `"${model?.club}" ?~ id`,
          fetch: fetch,
          requestKey: null,
          expand: "admins",
        })
  }
  depends("nav:load")

  return {
    clubs: clubs,
    teams: teams,
  }
}) satisfies LayoutLoad