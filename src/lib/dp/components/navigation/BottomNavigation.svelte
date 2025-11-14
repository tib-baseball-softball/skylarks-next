<script lang="ts">
  import StaticNavigationLinks from "$lib/dp/components/navigation/StaticNavigationLinks.svelte";
</script>

<nav
        class="bottom-nav"
>
  <ul>
    <StaticNavigationLinks classes="flex-col py-1.5 text-sm"/>
  </ul>
</nav>

<style lang="postcss">
    ul {
        display: flex;
    }

    .bottom-nav {
        position: fixed;
        bottom: 0;
        left: 0;
        width: 100%;
        padding-bottom: env(safe-area-inset-bottom);
        background-color: var(--nav-item-background);
        z-index: 2;
        border-top: 1px solid light-dark(var(--color-surface-500), var(--color-surface-100));
        border-bottom: 1px solid light-dark(var(--color-surface-500), var(--color-surface-100));

        @media (min-width: 48rem) {
            display: none;
        }
    }
</style>
