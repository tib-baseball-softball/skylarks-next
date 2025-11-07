import type { PageLoad } from "./$types"

export const load = (async ({ data }) => {
  return {
    streamed: {
      matchesCurrent: data?.matchesCurrent,
      matchesPrevious: data?.matchesPrevious,
      matchesNext: data?.matchesNext,
    },
  }
}) satisfies PageLoad
