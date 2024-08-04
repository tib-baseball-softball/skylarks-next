<script lang="ts">
    import {getDrawerStore} from "@skeletonlabs/skeleton";
    import {client} from "$lib/pocketbase";

    const drawerStore = getDrawerStore();

    function drawerClose(): void {
        drawerStore.close();
    }

    export let showLinkToMain = false
</script>

<nav class="list-nav py-1 px-6">
    <ul class="lg:flex items-end lg:flex-row gap-6">

        {#if showLinkToMain}
            <li><a href="/" on:click={drawerClose}>Start</a></li>
        {/if}

        <!--        <li><a href="/aktuelles" on:click={drawerClose}>Aktuelles</a></li>-->
        <li><a href="/gamecenter" on:click={drawerClose}>Gamecenter</a></li>
        <li><a href="/ligen" on:click={drawerClose}>Ligen</a></li>
        <li><a href="/teams" on:click={drawerClose}>Teams</a></li>
        <!--        <li><a href="/verein" on:click={drawerClose}>Verein</a></li>-->
        <!--        <li><a href="/kontakt" on:click={drawerClose}>Kontakt</a></li>-->
        
        {#if client.authStore.isValid}
            <li><a href="/account" on:click={drawerClose}>Account</a></li>
        {/if}
    </ul>
</nav>
