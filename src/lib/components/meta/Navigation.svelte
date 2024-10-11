<script lang="ts">
    import {
        Accordion,
        AccordionItem,
        getDrawerStore,
    } from "@skeletonlabs/skeleton";
    import { client } from "$lib/pocketbase";
    import {
        ChartLineUpOutline,
        ChartMixedOutline,
        HomeOutline,
        LockOutline,
        ProfileCardOutline,
        ShieldOutline,
        TableRowOutline,
        TicketOutline,
        UserCircleOutline,
        UsersGroupOutline,
        UsersOutline,
        UsersSolid,
    } from "flowbite-svelte-icons";
    import { browser } from "$app/environment";
    import type { ExpandedTeam } from "$lib/model/ExpandedResponse";
    import { authModel } from "$lib/pocketbase/Auth";

    const drawerStore = getDrawerStore();
    const model = $authModel;

    function drawerClose(): void {
        drawerStore.close();
    }

    async function getUserTeams(): Promise<ExpandedTeam[]> {
        if (browser) {
            return await client
                .collection("teams")
                .getFullList({ expand: "club" });
        }
        return [];
    }
</script>

<nav class="list-nav py-1 px-1 lg:px-4">
    <ul class="subpixel-antialiased">
        <li>
            <a href="/" onclick={drawerClose}>
                <HomeOutline size="lg" />
                <span>Start</span>
            </a>
        </li>

        <!--        <li><a href="/aktuelles" on:click={drawerClose}>Aktuelles</a></li>-->
        <li>
            <a href="/gamecenter" onclick={drawerClose}>
                <TicketOutline size="lg" />
                <span>Gamecenter</span>
            </a>
        </li>
        <li>
            <a href="/ligen" onclick={drawerClose}>
                <TableRowOutline size="lg" />
                <span>Leagues</span>
            </a>
        </li>
        <li>
            <a href="/teams" onclick={drawerClose}>
                <UsersGroupOutline size="lg" />
                <span>Teams</span>
            </a>
        </li>
        <!--        <li><a href="/verein" on:click={drawerClose}>Verein</a></li>-->
        <!--        <li><a href="/kontakt" on:click={drawerClose}>Kontakt</a></li>-->
    </ul>

    {#if client.authStore.isValid}
        <hr class="my-2" />

        <Accordion regionPanel="space-y-1">
            <AccordionItem open>
                <svelte:fragment slot="lead">
                    <ShieldOutline size="lg" />
                </svelte:fragment>

                <svelte:fragment slot="summary">
                    <span>My Account</span>
                </svelte:fragment>

                <svelte:fragment slot="content">
                    <a href="/account" onclick={drawerClose}>
                        <UserCircleOutline size="lg" />
                        <span>Dashboard</span>
                    </a>

                    <a href="/stats/{model?.id}" onclick={drawerClose}>
                        <ChartLineUpOutline size="lg" />
                        <span>Personal Stats</span>
                    </a>

                    <a href="/account/playerprofile" onclick={drawerClose}>
                        <ProfileCardOutline size="lg" />
                        <span>Player Profile</span>
                    </a>
                </svelte:fragment>
            </AccordionItem>
        </Accordion>

        <hr class="my-2" />

        {#await getUserTeams() then teams}
            <Accordion regionPanel="space-y-1">
                <AccordionItem open>
                    <svelte:fragment slot="lead">
                        <UsersSolid size="lg" />
                    </svelte:fragment>

                    <svelte:fragment slot="summary">
                        <span>My Teams</span>
                    </svelte:fragment>

                    <svelte:fragment slot="content">
                        {#each teams as team}
                            <a
                                href="/account/team/{team.id}"
                                onclick={drawerClose}
                            >
                                <UsersOutline size="lg" />
                                <span
                                    >{team.name} ({team?.expand?.club
                                        ?.acronym})</span
                                >
                            </a>

                            <a
                                class="ms-3"
                                href="/account/team/{team.id}/members"
                                onclick={drawerClose}
                            >
                                <UsersGroupOutline size="lg" />
                                <span>Players</span>
                            </a>
                        {/each}
                    </svelte:fragment>
                </AccordionItem>
            </Accordion>
        {/await}

        <hr class="my-2" />

        <Accordion regionPanel="space-y-1">
            <AccordionItem open>
                <svelte:fragment slot="lead">
                    <LockOutline size="lg" />
                </svelte:fragment>

                <svelte:fragment slot="summary">Administration</svelte:fragment>

                <svelte:fragment slot="content">
                    <a href="#" onclick={drawerClose}>
                        <ChartMixedOutline size="lg" />
                        <span>Admin Dashboard</span>
                    </a>
                </svelte:fragment>
            </AccordionItem>
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
