<script lang="ts">
    import LoginGuard from "$lib/auth/LoginGuard.svelte";
    import {LucideLogIn} from "lucide-svelte";

    const {children} = $props()
</script>

<LoginGuard>

    {@render children()}

    {#snippet otherwise()}
        <article>
            <h1 class="h2">Internal Area</h1>
            <p>Please log in!</p>

            <a title="Go to Login Page" class="btn preset-filled-primary-500" href="/login">
                <LucideLogIn/>
                Login Page
            </a>
        </article>
    {/snippet}

</LoginGuard>

<style>
    article {
        margin-top: calc(var(--spacing) * 2);

        p, a {
            margin-top: 1rem;
            margin-bottom: 1rem;
        }
    }
</style>