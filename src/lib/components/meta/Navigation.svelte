<script lang="ts">
    import { getDrawerStore } from "@skeletonlabs/skeleton";

    const drawerStore = getDrawerStore();

    function drawerClose(): void {
        drawerStore.close();
    }
</script>

<nav class="list-nav py-1 px-6">
    <ul class="lg:flex items-end lg:flex-row gap-6">
        <li><a href="/aktuelles">Aktuelles</a></li>
        <li><a href="/gamecenter">Gamecenter</a></li>
        <li><a href="/ligen">Ligen</a></li>
        <li><a href="/teams">Teams</a></li>
        <li><a href="/verein">Verein</a></li>
        <li><a href="/kontakt">Kontakt</a></li>
    </ul>
</nav>
