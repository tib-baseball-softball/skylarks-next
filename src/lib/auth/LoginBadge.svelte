<script lang="ts">
  import {client} from "../pocketbase/index.svelte";
  import {Avatar} from "@skeletonlabs/skeleton-svelte";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import {LucideLogIn} from "lucide-svelte";
  import Dialog from "$lib/components/utility/Dialog.svelte";
  import AccountModal from "$lib/auth/AccountModal.svelte";

  const {signupAllowed = true} = $props();

</script>

{#if authSettings.record && client.authStore.isValid}

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
