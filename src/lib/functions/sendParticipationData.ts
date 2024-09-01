import type { ParticipationsCreate, ParticipationsResponse, ParticipationsUpdate } from "$lib/model/pb-types"
import { client } from "$lib/pocketbase"

export async function sendParticipationData(data: ParticipationsCreate): Promise<ParticipationsResponse> {
  if (data.id) {
    return await client.collection("participations").update(data.id, data)
  } else {
    return await client.collection("participations").create(data)
  }
}
