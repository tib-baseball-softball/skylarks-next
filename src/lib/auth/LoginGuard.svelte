<script lang="ts">
  import {goto} from "$app/navigation";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import type {Snippet} from "svelte";

  const {
    destination,
    otherwise,
    children,
  }: {
    destination?: string;
    otherwise?: Snippet<[]>;
    children: Snippet<[]>;
  } = $props()

  $effect(() => {
    if (!!destination && client.authStore.isValid) {
      // navigate to destination if specified, and logged in
      goto(destination)
    }
  })
</script>

{#if authSettings.record}
    {@render children()}
{:else if otherwise}
    {@render otherwise()}
{/if}
