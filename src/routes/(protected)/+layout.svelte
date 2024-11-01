<script lang="ts">
  import LoginGuard from "$lib/auth/LoginGuard.svelte";

  const {children} = $props()
</script>

<LoginGuard>

    {@render children()}

    {#snippet otherwise()}
        <article class="mt-3">
            <h1 class="h2">Internal Area</h1>
            <p>Please log in!</p>
        </article>
    {/snippet}

</LoginGuard>