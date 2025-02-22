<script lang="ts">
  import type {CustomAuthModel, ExpandedParticipation} from "$lib/model/ExpandedResponse";
  import {getModalStore, type ModalComponent, type ModalSettings} from "@skeletonlabs/skeleton";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import ParticipationForm from "$lib/components/forms/ParticipationForm.svelte";
  // for some reason this import is falsely detected as unused
  //@ts-ignore
  import {Popover} from "bits-ui";

  interface Props {
    participation: ExpandedParticipation,
    isAdmin: boolean,
    classes?: string,
  }

  let {participation, isAdmin, classes = ""}: Props = $props();

  const modalStore = getModalStore();
  const authRecord = authSettings.record as CustomAuthModel;

  function triggerModal() {
    isOpen = false;
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

  const canEdit = $derived(participation.user === authRecord.id || isAdmin);
  let isOpen = $state(false);
</script>

<Popover.Root bind:open={isOpen}>
  <Popover.Trigger class={classes}>
    {participation?.expand?.user?.first_name}
  </Popover.Trigger>

  <Popover.Content>
    <div class="card variant-glass p-4 text-sm max-w-80 space-y-1">

      <div class="item-container">
        <p class="font-light text-xs">Created:</p>
        <p>{new Date(participation.created).toLocaleString()}</p>
      </div>

      <div class="item-container">
        <p class="font-light text-xs">Last Updated:</p>
        <p>{new Date(participation.updated).toLocaleString()}</p>
      </div>

      <div class="item-container !mb-2">
        <p class="font-light text-xs">Comment/Reason:</p>
        <p>{participation.comment}</p>
      </div>

      {#if canEdit}
        <button class="btn btn-sm variant-ghost-primary mt-2" onclick={triggerModal}>
          Edit
        </button>
      {/if}
    </div>
    <Popover.Close/>
    <Popover.Arrow/>
  </Popover.Content>

</Popover.Root>

<style>
    .item-container {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        justify-content: space-between;
    }
</style>