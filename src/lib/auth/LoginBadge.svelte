<script lang="ts">
  import {client} from "../pocketbase/index.svelte";
  import {Avatar} from "@skeletonlabs/skeleton-svelte";
  import {browser} from "$app/environment";
  import {invalidateAll} from "$app/navigation";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import {LucideLogIn} from "lucide-svelte";
  import {toastController} from "$lib/service/ToastController.svelte.ts";
  import Dialog from "$lib/components/utility/Dialog.svelte";
  import AccountModal from "$lib/auth/AccountModal.svelte";

  const {signupAllowed = true} = $props();

  client.authStore.onChange((_token, model) => {
    // do not do any auth stuff on the server
    if (browser) {
      if (model) {
        const {email} = model;
        toastController.trigger({
          id: crypto.randomUUID(),
          message: `Signed in as ${email}`,
          background: "preset-filled-success-500",
        });
        invalidateAll();
      } else {
        toastController.trigger({
          id: crypto.randomUUID(),
          message: "Logout successful",
          background: "preset-filled",
        });
        invalidateAll();
      }
    }
  });

</script>

{#if authSettings.record}

  <Dialog triggerClasses="badge">

    {#snippet triggerContent()}
      {#if authSettings.record?.avatar}
        <Avatar
                name={authSettings.record?.first_name + authSettings.record?.last_name}
                src={client.files.getURL(authSettings.record, authSettings.record?.avatar)}
                size="w-14"
        />

      {:else}

        <Avatar
                name={authSettings.record?.first_name + authSettings.record?.last_name}
                background="preset-filled-primary-500"
                size="w-14"
                classes="fill-white"
        />
      {/if}
    {/snippet}

    {#snippet title()}
      <h2>Account Overview</h2>
    {/snippet}

    <AccountModal/>
  </Dialog>

{:else}

  <a
          class="btn preset-tonal-primary border border-primary-500 flex items-center gap-2"
          href="/login"
          title={signupAllowed ? "Login / Register" : "Login"}
  >
    <LucideLogIn/>
    Login
  </a>

{/if}
