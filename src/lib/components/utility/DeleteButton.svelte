<script lang="ts">
  import {invalidate} from "$app/navigation";
  import {Trash} from "lucide-svelte";
  import type {Toast} from "$lib/types/Toast.ts";
  import {toastController} from "$lib/service/ToastController.svelte.ts";
  import Dialog from "$lib/components/utility/Dialog.svelte";
  import {closeModal} from "$lib/functions/closeModal.ts";

  interface Props {
    id: string;
    modelName: string;
    action: (id: string) => void;
    classes?: string;
    buttonText?: string;
  }

  let {
    id,
    modelName,
    action,
    classes = "btn-sm btn-icon preset-tonal-error border border-error-500",
    buttonText = "",
  }: Props = $props();

  function deleteItem() {
    try {
      action(id);

      const toastSettingsDeletions: Toast = {
        message: `${modelName} deleted successfully.`,
        background: "preset-filled-success-500",
      };

      toastController.trigger(toastSettingsDeletions);
      invalidate("nav:load");
    } catch {
      const toastSettingsDeletions: Toast = {
        message: `Error deleting ${modelName}.`,
        background: "preset-filled-error-500",
      };

      toastController.trigger(toastSettingsDeletions);
    }
    closeModal();
  }
</script>

<Dialog triggerClasses="btn {classes}" closeButtonClasses="sr-only">
  {#snippet triggerContent()}
    <Trash/>
    {#if buttonText}
      <span>{buttonText}</span>
    {/if}
  {/snippet}

  {#snippet title()}
    <span>Please Confirm</span>
  {/snippet}

  {#snippet description()}
    <span>Are you sure you wish to delete this {modelName}? In most cases, data cannot be recovered after deletion. Please be certain.</span>
  {/snippet}

  <div class="flex justify-end gap-2 mt-1">
    <button class="btn preset-tonal-surface border border-surface-500" type="button" onclick={closeModal}>Cancel</button>
    <button class="btn preset-filled" type="button" onclick={deleteItem}>Confirm</button>
  </div>
</Dialog>
