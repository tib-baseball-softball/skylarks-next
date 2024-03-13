<script lang="ts">
    import { getDrawerStore } from "@skeletonlabs/skeleton";

    const drawerStore = getDrawerStore();

    function drawerClose(): void {
        drawerStore.close();
    }
</script>

<nav class="list-nav py-3 px-6">
    <ul class="md:flex items-end md:flex-row gap-6">
        <li><a href="/aktuelles">Aktuelles</a></li>
        <li><a href="/blog">Blog</a></li>
        <li><a href="/blog">Blog</a></li>
        <li><a href="/blog">Blog</a></li>
        <li><a href="/contact">Contact</a></li>
    </ul>
</nav>
