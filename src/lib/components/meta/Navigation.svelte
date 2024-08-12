<script lang="ts">
    import {getDrawerStore} from "@skeletonlabs/skeleton";
    import {client} from "$lib/pocketbase";
    import {
        HomeOutline, ProfileCardOutline, TableRowOutline, TicketOutline, UsersGroupOutline
    } from "flowbite-svelte-icons";

    const drawerStore = getDrawerStore();

    function drawerClose(): void {
        drawerStore.close();
    }
</script>

<nav class="list-nav py-1 px-4">
    <ul class="subpixel-antialiased">

        <li>
            <a href="/" on:click={drawerClose}>
                <HomeOutline size="lg"/>
                <span>Start</span>
            </a>
        </li>

        <!--        <li><a href="/aktuelles" on:click={drawerClose}>Aktuelles</a></li>-->
        <li>
            <a href="/gamecenter" on:click={drawerClose}>
                <TicketOutline size="lg"/>
                <span>Gamecenter</span>
            </a>
        </li>
        <li>

            <a href="/ligen" on:click={drawerClose}>
                <TableRowOutline size="lg"/>
                <span>Ligen</span>
            </a>
        </li>
        <li>
            <a href="/teams" on:click={drawerClose}>
                <UsersGroupOutline size="lg"/>
                <span>Teams</span>
            </a>

        </li>
        <!--        <li><a href="/verein" on:click={drawerClose}>Verein</a></li>-->
        <!--        <li><a href="/kontakt" on:click={drawerClose}>Kontakt</a></li>-->

    </ul>

    <hr class="my-2"/>

    <ul>
        <li>
            {#if client.authStore.isValid}
                <a href="/account" on:click={drawerClose}>
                    <ProfileCardOutline size="lg"/>
                    <span>Account</span>
                </a>
            {/if}
        </li>
    </ul>
</nav>

<style lang="postcss">
    li {
        @apply my-2;
    }

    a {
        display: flex;
        @apply items-center;
    }
</style>