<script lang="ts">
    import LoginGuard from "$lib/auth/LoginGuard.svelte";

    const {children} = $props()
</script>

<LoginGuard>

    {@render children()}

    {#snippet otherwise()}
        <article class="prose mt-3">
            <h1 class="">Interner Bereich</h1>
            <p>Bitte einloggen!</p>
        </article>
    {/snippet}

</LoginGuard>