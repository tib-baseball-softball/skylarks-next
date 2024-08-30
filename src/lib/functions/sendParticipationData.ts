import type {ParticipationsCreate, ParticipationsResponse, ParticipationsUpdate} from "$lib/model/pb-types"
import { client } from "$lib/pocketbase"

export async function sendParticipationData(state: EventParticipationState, data: ParticipationsCreate): Promise<ParticipationsResponse> {
    data.state = state
    if (data.id) {
        return await client.collection("participations").update(data.id, data)
    } else {
        return await client.collection("participations").create(data)
    }
}