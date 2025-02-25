<script lang="ts">
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import {getToastStore} from "@skeletonlabs/skeleton";

  interface props {
    parent: any;
  }

  const authRecord = authSettings.record as CustomAuthModel;
  const toastStore = getToastStore();

  let {parent}: props = $props();

  const form = $state({
    email: authRecord.email
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();
    parent.onClose();

    const sent = await client.collection("users").requestEmailChange(form.email);

    if (sent) {
      toastStore.trigger({
        message: "A verification email has been sent to your new address.",
        background: "variant-filled-success",
      });
    } else {
      toastStore.trigger({
        message: "An error occurred while sending your email change request.",
        background: "variant-filled-error",
      });
    }
    client.authStore.clear();
  }
</script>

<div class="card p-6 max-w-xl">
  <header class="text-xl font-semibold">
    Change email address for {authRecord.first_name}
    {authRecord.last_name}
  </header>

  <form onsubmit={submitForm} class="mt-4 space-y-3">

    <label class="label">
      E-Mail
      <input
              name="email"
              class="input"
              bind:value={form.email}
              required
              type="email"
      />
      <span class="mt-3 font-light text-sm">
                You will be logged out and instructed to verify your new email address.
            </span>
    </label>

    <div class="flex justify-end gap-3 mt-3">
      <button disabled={form.email === authRecord.email} type="submit" class="mt-2 btn variant-ghost-primary">
        Confirm
      </button>
    </div>
  </form>
</div>
