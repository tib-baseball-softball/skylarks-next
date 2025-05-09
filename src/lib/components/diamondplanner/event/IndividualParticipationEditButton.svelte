<script lang="ts">
  import type {CustomAuthModel, ExpandedParticipation} from "$lib/model/ExpandedResponse";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import ParticipationForm from "$lib/components/forms/ParticipationForm.svelte";
  // for some reason this import is falsely detected as unused
  //@ts-ignore
  import {Popover} from "bits-ui";
  import {MessageCircleMore} from "lucide-svelte";
  import Dialog from "$lib/components/utility/Dialog.svelte";

  interface Props {
    participation: ExpandedParticipation,
    isAdmin: boolean,
    classes?: string,
  }

  let {participation, isAdmin, classes = ""}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  const canEdit = $derived(participation.user === authRecord.id || isAdmin);
  let isOpen = $state(false);
</script>

<Popover.Root bind:open={isOpen}>
  <Popover.Trigger>
    {#snippet child({props})}
      <button {...props} class="{classes} relative inline-block">
        {#if participation.comment}
          <span class="badge-icon preset-filled-surface-400-600">
            <MessageCircleMore size="13"/>
          </span>
        {/if}
        {participation?.expand?.user?.first_name}
      </button>
    {/snippet}
  </Popover.Trigger>

  <Popover.Content>
    <div class="card bg-surface-50 dark:bg-surface-800 border shadow-2xl p-4 text-sm max-w-80 space-y-1 text-black dark:text-white">

      <div class="item-container">
        <p class="font-light text-xs">Created:</p>
        <p>{new Date(participation.created).toLocaleString()}</p>
      </div>

      <div class="item-container">
        <p class="font-light text-xs">Last Updated:</p>
        <p>{new Date(participation.updated).toLocaleString()}</p>
      </div>

      <div class="item-container mb-2!">
        <p class="font-light text-xs">Comment/Reason:</p>
        <p>{participation.comment}</p>
      </div>

      {#if canEdit}
        <Dialog triggerClasses="btn btn-sm preset-tonal-primary border border-primary-500 mt-2">

          {#snippet triggerContent()}
            Edit
          {/snippet}

          {#snippet title()}
            <header>
              <h2 class="h3">
                Edit participation data
                for {participation?.expand?.user?.first_name} {participation?.expand?.user?.last_name}
              </h2>
            </header>
          {/snippet}

          <ParticipationForm {participation}/>
        </Dialog>
      {/if}
    </div>
    <Popover.Close/>
    <Popover.Arrow/>
  </Popover.Content>

</Popover.Root>

<style lang="postcss">
    .item-container {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        justify-content: space-between;
    }

    .badge-icon {
        top: -15px;
        right: -15px;
        z-index: 10;
        position: absolute;
    }
</style>