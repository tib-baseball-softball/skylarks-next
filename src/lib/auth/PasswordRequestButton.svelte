<script lang="ts">
  import {closeModal} from "$lib/functions/closeModal.js";
  import {client} from "$lib/pocketbase/index.svelte";
  import {toastController} from "$lib/service/ToastController.svelte.ts";
  import type {Snippet} from "svelte";
  import Dialog from "$lib/components/utility/Dialog.svelte";

  interface Props {
    email: string
    classes?: string,
    disabled?: boolean,
    additionalAction?: () => void
    children?: Snippet
  }

  let {
    email,
    disabled = false,
    additionalAction = () => {
    },
    classes = "",
    children,
  }: Props = $props();

  async function triggerPasswordChange() {
    additionalAction();

    try {
      const success = await client.collection("users").requestPasswordReset(email);

      if (success) {
        toastController.trigger({
          id: crypto.randomUUID(),
          message: `If we have your email (${email}) on file, you will receive a reset password request in your inbox.`,
          background: "variant-filled-success",
        });
      }

      if (client.authStore.isValid) {
        client.authStore.clear();
      }

    } catch {
      toastController.trigger({
        id: crypto.randomUUID(),
        message: `Error sending reset password request.`,
        background: "variant-filled-error",
      });
    }
    closeModal();
  }
</script>

<Dialog
        triggerClasses={classes}
        disabled={disabled}
        closeButtonClasses="sr-only"
>
  {#snippet triggerContent()}
    {@render children?.()}
  {/snippet}

  {#snippet title()}
    <span>Please Confirm</span>
  {/snippet}

  {#snippet description()}
    <span>
      Send change password request for <code class="code">{email ? email : "---"}</code>?
      You will be logged out and an email with instructions will be sent to your address.
    </span>
  {/snippet}

  <div class="flex justify-end gap-2 mt-1">
    <button class="btn variant-ghost-surface" type="button" onclick={closeModal}>Cancel</button>
    <button class="btn variant-filled" type="button" onclick={triggerPasswordChange}>Confirm</button>
  </div>
</Dialog>