<script lang="ts">
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import {toastController} from "$lib/service/ToastController.svelte.ts";
  import {closeModal} from "$lib/functions/closeModal.ts";

  const authRecord = authSettings.record as CustomAuthModel;

  const form = $state({
    email: authRecord.email
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    const sent = await client.collection("users").requestEmailChange(form.email);

    if (sent) {
      closeModal()
      toastController.trigger({
        id: crypto.randomUUID(),
        message: "A verification email has been sent to your new address.",
        background: "preset-filled-success-500",
      });
    } else {
      toastController.trigger({
        id: crypto.randomUUID(),
        message: "An error occurred while sending your email change request.",
        background: "preset-filled-error-500",
      });
    }
    client.authStore.clear();
  }
</script>

<form onsubmit={submitForm} class="mt-4 space-y-3">

  <label class="label">
    E-Mail
    <input
            name="email"
            class="input"
            bind:value={form.email}
            required
            type="email"
            autocomplete="email"
    />
    <span class="mt-3 font-light text-sm">
                You will be logged out and instructed to verify your new email address.
            </span>
  </label>

  <div class="flex justify-end gap-3 mt-3">
    <button disabled={form.email === authRecord.email} type="submit" class="mt-2 btn preset-tonal-primary border border-primary-500">
      Confirm
    </button>
  </div>
</form>
