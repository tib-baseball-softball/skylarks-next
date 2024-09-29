<script lang="ts">
    import {Accordion, AccordionItem, getDrawerStore} from "@skeletonlabs/skeleton";
    import {client} from "$lib/pocketbase";
    import {
        ChartMixedOutline,
        HomeOutline, LockOutline, ProfileCardOutline, TableRowOutline, TicketOutline, UsersGroupOutline, UsersOutline
    } from "flowbite-svelte-icons";
    import {browser} from "$app/environment";
    import type { ExpandedTeam } from "$lib/model/ExpandedResponse";

    const drawerStore = getDrawerStore();

    function drawerClose(): void {
        drawerStore.close();
    }

    async function getUserTeams(): Promise<ExpandedTeam[]> {
        if (browser) {
            return await client.collection("teams").getFullList({expand: "club"})
        }
        return []
    }
</script>

<nav class="list-nav py-1 px-1 lg:px-4">
    <ul class="subpixel-antialiased">

        <li>
            <a href="/" onclick={drawerClose}>
                <HomeOutline size="lg"/>
                <span>Start</span>
            </a>
        </li>

        <!--        <li><a href="/aktuelles" on:click={drawerClose}>Aktuelles</a></li>-->
        <li>
            <a href="/gamecenter" onclick={drawerClose}>
                <TicketOutline size="lg"/>
                <span>Gamecenter</span>
            </a>
        </li>
        <li>

            <a href="/ligen" onclick={drawerClose}>
                <TableRowOutline size="lg"/>
                <span>Ligen</span>
            </a>
        </li>
        <li>
            <a href="/teams" onclick={drawerClose}>
                <UsersGroupOutline size="lg"/>
                <span>Teams</span>
            </a>

        </li>
        <!--        <li><a href="/verein" on:click={drawerClose}>Verein</a></li>-->
        <!--        <li><a href="/kontakt" on:click={drawerClose}>Kontakt</a></li>-->

    </ul>

    {#if client.authStore.isValid}
        <hr class="my-2"/>

        <Accordion
            regionPanel="space-y-1"
        >
            <AccordionItem open>
                <svelte:fragment slot="lead">
                    <LockOutline size="lg"/>
                </svelte:fragment>

                <svelte:fragment slot="summary">
                    <span>Kaderplanung</span>
                </svelte:fragment>

                <svelte:fragment slot="content">

                    <a href="/account" onclick={drawerClose}>
                        <ProfileCardOutline size="lg"/>
                        <span>Meine Seite</span>
                    </a>

                    {#await getUserTeams() then teams}
                        {#each teams as team}
                            <a href="/account/team/{team.id}" onclick={drawerClose}>
                                <UsersOutline size="lg"/>
                                <span>{team.name} ({team?.expand?.club?.acronym})</span>
                            </a>
                        {/each}
                    {/await}

                    <hr/>

                    <a href="#" onclick={drawerClose}>
                        <ChartMixedOutline size="lg"/>
                        <span>Admin Dashboard</span>
                    </a>

                </svelte:fragment>
            </AccordionItem>
            <!-- ... -->
        </Accordion>


    {/if}
</nav>

<style lang="postcss">
    li {
        @apply my-2;
    }

    a {
        display: flex;
        @apply items-center;
    }

    span {
        @apply text-nowrap;
    }
</style>