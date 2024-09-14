<script lang="ts">
  import { goto } from "$app/navigation";
  import { client } from "$lib/pocketbase";
  import type { Snippet } from "svelte";
  import { authModel } from "$lib/pocketbase/Auth";

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

  const authorized = $derived($authModel)
</script>

{#if authorized}
  {@render children()}
{:else if otherwise}
  {@render otherwise()}
{/if}
