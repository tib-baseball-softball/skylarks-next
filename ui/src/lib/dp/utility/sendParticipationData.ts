import {client} from "$lib/dp/client.svelte.js";
import type {ParticipationsCreate, ParticipationsResponse} from "$lib/dp/types/pb-types.ts";
import {Collection} from "$lib/dp/enum/Collection.ts";
import { ClientResponseError } from "pocketbase";
import { toastController } from "$lib/dp/service/ToastController.svelte";

export async function sendParticipationData(
  data: ParticipationsCreate
): Promise<void> {
  try {
    if (data.id) {
      await client.collection(Collection.Participations).update(data.id, data);
    } else {
      await client.collection(Collection.Participations).create(data);
    }
  } catch (error) {
    if (error instanceof ClientResponseError) {
      toastController.trigger({
        message:
          "Saving event participation failed. Please try again and contact support if the problem persists.",
        background: "preset-filled-error-500",
      });
    } else {
      toastController.triggerGenericErrorMessage();
    }
  }
}
