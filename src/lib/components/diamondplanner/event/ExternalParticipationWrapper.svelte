<script lang="ts">
  import type {UserParticipationDTO} from "$lib/model/UserParticipationDTO.ts";
  import type {ExpandedParticipation} from "$lib/model/ExpandedResponse.ts";
  import IndividualParticipationEditButton
    from "$lib/components/diamondplanner/event/IndividualParticipationEditButton.svelte";

  interface Props {
    dto: UserParticipationDTO,
    eventID: string,
    isAdmin: boolean,
    classes?: string,
  }

  let {dto, eventID, isAdmin, classes = ""}: Props = $props();

  const participation: ExpandedParticipation = {
    collectionId: "",
    collectionName: "participations",
    comment: "",
    created: "",
    event: eventID,
    id: "",
    state: "",
    updated: "",
    user: dto.id,
    expand: {
      // @ts-expect-error
      user: {
        id: dto.id,
        first_name: dto.first_name,
        last_name: dto.last_name,
      }
    }
  };
</script>

<IndividualParticipationEditButton {classes} {isAdmin} {participation}/>