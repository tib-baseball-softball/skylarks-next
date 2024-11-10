<script lang="ts">
  import type {AuthProviderInfo, RecordModel, RecordService} from "pocketbase";
  import {providerLogin} from "$lib/pocketbase/Auth.svelte";
  import {goto} from "$app/navigation";

  /**
   * Shows a button to log in with an external service. Uses branded image if available,
   * falls back to simple button if not.
   */

  interface Props {
    authProvider: AuthProviderInfo,
    collection: RecordService<RecordModel>
    parent: any
  }

  const {authProvider, collection, parent}: Props = $props()

  async function submitOAuthRequest(provider: AuthProviderInfo) {
    providerLogin(provider, collection)
    goto("/account", {invalidateAll: true})

    //@ts-ignore
    parent.onClose()
  }

  const isGenericProvider = $derived(
      authProvider.name !== "google" &&
      authProvider.name !== "apple" &&
      authProvider.name !== "github" &&
      authProvider.name !== "discord"
  )
</script>

<button
        class:btn={isGenericProvider}
        class:variant-ghost-tertiary={isGenericProvider}
        type="button"
        onclick={() => submitOAuthRequest(authProvider)}
        aria-label="sign in with {authProvider.displayName}"
>
    {#if (authProvider.name === "google")}
        <img class="light-logo" src="/providers/Google_light.svg" alt="Google logo" width="40">
        <img class="dark-logo" src="/providers/Google_dark.svg" alt="Google logo" width="40">

    {:else if authProvider.name === "apple"}
        <img class="light-logo" src="/providers/siwa_apple_light.svg" alt="Apple logo" width="40">
        <img class="dark-logo" src="/providers/siwa_apple_dark.svg" alt="Apple logo" width="40">

    {:else if authProvider.name === "github"}
        <img class="light-logo" src="/providers/github-mark-white.svg" alt="GitHub logo" width="40">
        <img class="dark-logo" src="/providers/github-mark.svg" alt="GitHub logo" width="40">

    {:else if authProvider.name === "discord"}
        <img src="/providers/discord_white_in_blue_circle.svg" alt="Discord logo" width="40">

    {:else }
        <span>{authProvider.displayName}</span>
    {/if}
</button>