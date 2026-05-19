import {client} from "$lib/dp/client.svelte.js";
import type {ParticipationsCreate, ParticipationsResponse} from "$lib/dp/types/pb-types.ts";
import {Collection} from "$lib/dp/enum/Collection.ts";

export async function sendParticipationData(
  data: ParticipationsCreate
): Promise<ParticipationsResponse> {
  if (data.id) {
    return await client.collection(Collection.Participations).update(data.id, data);
  } else {
    return await client.collection(Collection.Participations).create(data);
  }
}
