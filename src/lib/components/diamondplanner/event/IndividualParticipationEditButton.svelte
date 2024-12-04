<script lang="ts">
  import type {CustomAuthModel, ExpandedParticipation} from "$lib/model/ExpandedResponse";
  import {getModalStore, type ModalComponent, type ModalSettings} from "@skeletonlabs/skeleton";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import ParticipationForm from "$lib/components/forms/ParticipationForm.svelte";

  interface Props {
    participation: ExpandedParticipation,
    isAdmin: boolean,
    classes?: string,
  }

  let {participation, isAdmin, classes = ""}: Props = $props()

  const modalStore = getModalStore()
  const authRecord = authSettings.record as CustomAuthModel

  function triggerModal() {
    const modalComponent: ModalComponent = {
      ref: ParticipationForm,
      props: {participation: participation},
    };

    const modal: ModalSettings = {
      type: "component",
      component: modalComponent,
    };
    modalStore.trigger(modal);
  }

  const canEdit = $derived(participation.user === authRecord.id || isAdmin)
</script>

{#if canEdit}
    <button class="{classes}" onclick={triggerModal}>
        {participation?.expand?.user?.first_name}
    </button>
{:else }
    <div class="{classes} cursor-default">
        {participation?.expand?.user?.first_name}
    </div>
{/if}