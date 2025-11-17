import type {PageLoad} from "../../../.svelte-kit/types/src/routes";

export const load = (async ({data}) => {
  return {
    streamed: {
      matchesCurrent: data?.matchesCurrent,
      matchesPrevious: data?.matchesPrevious,
      matchesNext: data?.matchesNext,
    },
  };
}) satisfies PageLoad;
