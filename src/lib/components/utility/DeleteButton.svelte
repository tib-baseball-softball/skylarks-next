<script lang="ts">
  import {getModalStore, type ModalSettings,} from "@skeletonlabs/skeleton";
  import {invalidate} from "$app/navigation";
  import {Trash} from "lucide-svelte";
  import type {Toast} from "$lib/types/Toast.ts";
  import {toastController} from "$lib/service/ToastController.svelte.ts";

  const modalStore = getModalStore();

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

  function triggerDeleteModal() {
    const modal: ModalSettings = {
      type: "confirm",
      title: "Please Confirm",
      body: `Are you sure you wish to delete this ${modelName}? In most cases, data cannot be recovered after deletion. Please be certain.`,
      response: (r: boolean) => {
        if (r) {
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
        }
      },
    };
    modalStore.trigger(modal);
  }
</script>

<button
  class="btn {classes}"
  aria-label="delete ${modelName}"
  onclick={triggerDeleteModal}
>
  <Trash />

  {#if buttonText}
    <span>{buttonText}</span>
  {/if}
</button>
