import type {PageLoad} from "../../../../.svelte-kit/types/src/routes";

export const load: PageLoad = async ({parent}) => {
  return await parent();
};
