<script lang="ts">
  import {getModalStore, getToastStore, type ModalSettings} from "@skeletonlabs/skeleton";
  import {client} from "$lib/pocketbase/index.svelte";

  interface Props {
    email: string
    classes?: string,
    disabled?: boolean,
    additionalAction?: () => void
    children?: import('svelte').Snippet
  }

  let {
    email,
    disabled = false,
    additionalAction = () => {
    },
    classes = "",
    children,
  }: Props = $props();

  const modalStore = getModalStore();
  const toastStore = getToastStore()

  function triggerPasswordChangeModal() {
    additionalAction()

    const modal: ModalSettings = {
      type: 'confirm',
      title: 'Please Confirm',
      body: `Send change password request for ${email}? You will be logged out and an email with instructions will be sent to your address.`,
      response: (r: boolean) => {
        if (r) {
          try {
            client.collection("users").requestPasswordReset(email)
            toastStore.trigger({
              message: `Sent reset password request to ${email}.`,
              background: "variant-filled-success",
            })
            if (client.authStore.isValid) {
              client.authStore.clear()
            }

          } catch {
            toastStore.trigger({
              message: `Error sending reset password request.`,
              background: "variant-filled-error",
            })
          }
        }
      },
    };
    modalStore.trigger(modal);
  }
</script>

<button
        type="button"
        aria-label="change password"
        class="{classes}"
        disabled={disabled}
        onclick="{() => triggerPasswordChangeModal()}"
>
    {@render children?.()}
</button>