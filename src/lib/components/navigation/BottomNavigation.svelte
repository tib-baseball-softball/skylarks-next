<script lang="ts">
  import StaticNavigationLinks from "$lib/components/navigation/StaticNavigationLinks.svelte";
</script>

<nav
    class="lg:hidden bottom-nav bg-surface-50 dark:bg-surface-800 border-t border-b border-surface-500 dark:border-surface-100 z-2"
>
  <ul class="flex">
    <StaticNavigationLinks classes="flex-col py-1.5 text-sm"/>
  </ul>
</nav>

<style lang="postcss">
    .bottom-nav {
        position: fixed;
        bottom: 0;
        left: 0;
        width: 100%;
        padding-bottom: env(safe-area-inset-bottom);
    }
</style>
