<script lang="ts">
    import '../app.postcss';
    import {AppBar, Modal, type ModalComponent, Toast} from "@skeletonlabs/skeleton";
    import Navigation from "$lib/components/meta/Navigation.svelte";
    import {LightSwitch} from '@skeletonlabs/skeleton';
    import {initializeStores, Drawer, getDrawerStore} from '@skeletonlabs/skeleton';
    import Footer from "$lib/components/meta/Footer.svelte";
    import LoginBadge from "$lib/auth/LoginBadge.svelte";
    import LoginForm from "$lib/auth/LoginForm.svelte";
    import AccountModal from "$lib/auth/AccountModal.svelte";
    import {PUBLIC_AUTH_FUNCS_ENABLED} from "$env/static/public";
    import {onMount} from "svelte";
    import EventForm from '$lib/components/forms/EventForm.svelte';
    import TeamForm from '$lib/components/forms/TeamForm.svelte';
    import PlayerDataForm from '$lib/components/forms/PlayerDataForm.svelte';
    import ClubForm from "$lib/components/forms/ClubForm.svelte";

    interface Props {
        children?: import('svelte').Snippet;
    }

    let { children }: Props = $props();

    initializeStores()

    const modalRegistry: Record<string, ModalComponent> = {
        // Set a unique modal ID, then pass the component reference
        loginForm: {ref: LoginForm},
        accountOverview: {ref: AccountModal}
    }

    const drawerStore = getDrawerStore()

    function navDrawerOpen(): void {
        drawerStore.open({
          id: "nav",
          width: "w-[70%] sm:w-[40%]"
        })
    }

    // LightSwitch Workaround: https://github.com/skeletonlabs/skeleton/issues/2598
    onMount(() => {
        const e = document.documentElement.classList, t = localStorage.getItem("modeUserPrefers") === "false",
            n = !("modeUserPrefers" in localStorage), r = window.matchMedia("(prefers-color-scheme: dark)").matches;
        t || n && r ? e.add("dark") : e.remove("dark")
    })

</script>

<Drawer>
    {#if $drawerStore.id === "nav"}

    <div class="flex justify-around p-2">
        <img class="max-w-14" src="/berlin_skylarks_logo.svg" alt="Skylarks Team Logo">

        <h2 class="p-4 antialiased">Berlin Skylarks</h2>
    </div>

    <hr class="mb-2"/>

    <Navigation></Navigation>

    {:else if $drawerStore.id === "event-form"}
        <EventForm/>
    {:else if $drawerStore.id === "team-form"}
        <TeamForm/>
    {:else if $drawerStore.id === "player-data-form"}
        <PlayerDataForm/>
    {:else if $drawerStore.id === "club-form"}
        <ClubForm/>
    {/if}
</Drawer>

<!--Singletons-->
<Modal components={modalRegistry}/>
<Toast/>

<div class="h-screen grid grid-rows-[auto_1fr_auto]">
    <!-- Header -->
    <header>

        <AppBar
            gridColumns="grid-cols-6"
            slotDefault="place-self-center place-content-between w-full col-span-4"
            slotTrail="place-content-end"
            padding="p-3"
            background="bg-surface-500/5"
        >
            <svelte:fragment slot="lead">
                <div class="flex items-center justify-content-start">
                    <button aria-label="open navigation" class="md:hidden btn btn-sm mr-4" onclick={navDrawerOpen}>
                    <span>
                        <svg viewBox="0 0 100 80" class="fill-token w-4 h-4">
                            <rect width="100" height="20"/>
                            <rect y="30" width="100" height="20"/>
                            <rect y="60" width="100" height="20"/>
                        </svg>
                    </span>
                    </button>

                    <a aria-label="to home page" href="/" class="hidden md:block ms-3">
                        <img class="min-w-16" src="/berlin_skylarks_logo.svg" alt="Skylarks Team Logo">
                    </a>
                </div>
            </svelte:fragment>

            <svelte:fragment slot="default">
                <section class="w-full justify-center items-center hidden lg:flex py-2">



                </section>
            </svelte:fragment>

            <svelte:fragment slot="trail">
                <div class="lg:me-5 flex items-center gap-5 flex-shrink-0">

                    <LightSwitch/>

                    {#if PUBLIC_AUTH_FUNCS_ENABLED === "true"}
                        <LoginBadge signupAllowed={true}/>
                    {/if}

                </div>
            </svelte:fragment>

        </AppBar>

        <hr class="!border-b-2">
    </header>

    <!-- Grid Column -->

    <div class="container grid grid-cols-1 md:grid-cols-[auto_1fr_auto]">

        <!-- Sidebar (Left) -->
        <aside class="bg-surface-500/5 p-2 sticky top-0 col-span-1 hidden h-screen md:block lg:min-w-72">
            <Navigation></Navigation>
        </aside>

        <!-- Main -->

        <main class="col-span-1 space-y-4 lg:space-y-6 mx-4 md:mx-6 lg:mx-8 mb-4 lg:mb-6">

            {@render children?.()}

        </main>
    </div>

    <!-- Footer -->
    <footer>
        <hr class="!border-t-2">
        <Footer></Footer>
    </footer>
</div>
