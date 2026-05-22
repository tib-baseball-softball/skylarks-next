import type {ExpandedParticipation} from "$lib/dp/types/ExpandedResponse.ts";
import {Collection} from "$lib/dp/enum/Collection.ts";

/**
 * Represents user participation in an event.
 * Reduced response used for ghost responses in every event (users who have not yet responded).
 */
export type UserParticipationDTO = {
  id: string
  first_name: string
  last_name: string
  display_name: string
}

/**
 * Creates an expanded participation object from a UserParticipationDTO.
 * Used for drag and drop operations with ghost responses.
 *
 * @param dto the UserParticipationDTO to expand
 * @param eventID the ID of the event the participation is for
 */
export function participationDTOToExpandedParticipation(dto: UserParticipationDTO, eventID: string): ExpandedParticipation {
  return {
    id: "",
    collectionId: "",
    collectionName: Collection.Participations,
    comment: "",
    created: "",
    event: eventID,
    state: "",
    updated: "",
    user: dto.id
  };
}
