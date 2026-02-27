<script lang="ts">
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  let form = $derived.by(() => {
    const formData = $state({
      email: authRecord.email,
    });
    return formData;
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    const sent = await client.collection("users").requestEmailChange(form.email);

    if (sent) {
      closeModal();
      toastController.trigger({
        message: "A verification email has been sent to your new address.",
        background: "preset-filled-success-500",
      });
    } else {
      toastController.trigger({
        message: "An error occurred while sending your email change request.",
        background: "preset-filled-error-500",
      });
    }
    client.authStore.clear();
  }
</script>

<form class="form-root" onsubmit={submitForm}>

  <label class="label">
    E-Mail
    <input
      autocomplete="email"
      bind:value={form.email}
      class="input"
      name="email"
      required
      type="email"
    />
    <span class="help-text">
                You will be logged out and instructed to verify your new email address.
            </span>
  </label>

  <div class="actions">
    <button class="btn preset-tonal-primary submit-btn" disabled={form.email === authRecord.email}
            type="submit">
      Confirm
    </button>
  </div>
</form>

<style>
  .form-root {
    margin-top: calc(var(--spacing) * 4);
  }

  .form-root > * + * {
    margin-top: calc(var(--spacing) * 3);
  }

  .help-text {
    margin-top: calc(var(--spacing) * 3);
    font-weight: var(--font-weight-light);
    font-size: var(--text-sm);
  }

  .actions {
    display: flex;
    justify-content: flex-end;
    gap: calc(var(--spacing) * 3);
    margin-top: calc(var(--spacing) * 3);
  }

  .submit-btn {
    margin-top: calc(var(--spacing) * 2);
    border: 1px solid var(--color-primary-500);
  }
</style>
