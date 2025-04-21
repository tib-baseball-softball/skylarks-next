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
    classes = "btn-sm btn-icon variant-ghost-error",
    buttonText = "",
  }: Props = $props();

  function deleteItem() {
    try {
      action(id);

      const toastSettingsDeletions: Toast = {
        id: crypto.randomUUID(),
        message: `${modelName} deleted successfully.`,
        background: "variant-filled-success",
      };

      toastController.trigger(toastSettingsDeletions);
      invalidate("nav:load");
    } catch {
      const toastSettingsDeletions: Toast = {
        id: crypto.randomUUID(),
        message: `Error deleting ${modelName}.`,
        background: "variant-filled-error",
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
    <button class="btn variant-ghost-surface" type="button" onclick={closeModal}>Cancel</button>
    <button class="btn variant-filled" type="button" onclick={deleteItem}>Confirm</button>
  </div>
</Dialog>