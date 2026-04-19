<script lang="ts">
  import IndividualParticipationEditButton from "$lib/dp/components/event/IndividualParticipationEditButton.svelte";
  import type {ExpandedParticipation} from "$lib/dp/types/ExpandedResponse.ts";
  import type {UserParticipationDTO} from "$lib/dp/types/UserParticipationDTO.ts";
  import {Collection} from "$lib/dp/enum/Collection.ts";

  interface Props {
    dto: UserParticipationDTO;
    eventID: string;
    isAdmin: boolean;
    classes?: string;
  }

  let {dto, eventID, isAdmin, classes = ""}: Props = $props();

  //@ts-expect-error - casting DTO to proper type
  const participation: ExpandedParticipation = $derived({
    collectionId: "",
    collectionName: Collection.Participations,
    comment: "",
    created: "",
    event: eventID,
    id: "",
    state: "",
    updated: "",
    user: dto.id,
    expand: {
      user: {
        id: dto.id,
        first_name: dto.first_name,
        last_name: dto.last_name,
        display_name: dto.display_name,
      },
    },
  });
</script>

<IndividualParticipationEditButton {classes} {isAdmin} {participation}/>
