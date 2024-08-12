<script lang="ts">
    import '../app.postcss';
    import {AppBar, Modal, type ModalComponent, Toast} from "@skeletonlabs/skeleton";
    import Navigation from "$lib/components/meta/Navigation.svelte";
    import {LightSwitch} from '@skeletonlabs/skeleton';
    import {initializeStores, Drawer, getDrawerStore} from '@skeletonlabs/skeleton';
    import Footer from "$lib/components/meta/Footer.svelte";
    import {ExclamationCircleOutline} from "flowbite-svelte-icons";
    import LoginBadge from "$lib/auth/LoginBadge.svelte";
    import LoginForm from "$lib/auth/LoginForm.svelte";
    import AccountModal from "$lib/auth/AccountModal.svelte";
    import {PUBLIC_AUTH_FUNCS_ENABLED} from "$env/static/public";
    import {onMount} from "svelte";

    initializeStores()

    const modalRegistry: Record<string, ModalComponent> = {
        // Set a unique modal ID, then pass the component reference
        loginForm: {ref: LoginForm},
        accountOverview: {ref: AccountModal}
    }

    const drawerStore = getDrawerStore()

    function drawerOpen(): void {
        drawerStore.open({})
    }

    // LightSwitch Workaround: https://github.com/skeletonlabs/skeleton/issues/2598
    onMount(() => {
        const e = document.documentElement.classList, t = localStorage.getItem("modeUserPrefers") === "false",
            n = !("modeUserPrefers" in localStorage), r = window.matchMedia("(prefers-color-scheme: dark)").matches;
        t || n && r ? e.add("dark") : e.remove("dark")
    })

</script>

<Drawer width="w-[70%]">
    <div class="flex justify-around p-2">
        <img class="max-w-14" src="/berlin_skylarks_logo.svg" alt="Skylarks Team Logo">

        <h2 class="p-4">Berlin Skylarks</h2>
    </div>

    <hr class="mb-2"/>

    <Navigation showLinkToMain="{true}"></Navigation>
</Drawer>

<!--Singletons-->
<Modal components={modalRegistry}/>
<Toast/>

<!--Main-->

<div class="h-screen grid grid-rows-[auto_1fr_auto]">
    <!-- Header -->
    <header>

        <AppBar
            gridColumns="grid-cols-6"
            slotDefault="place-self-center place-content-between w-full col-span-4"
            slotTrail="place-content-end"
        >
            <svelte:fragment slot="lead">
                <div class="flex items-center justify-content-around">
                    <button class="lg:hidden btn btn-sm mr-4" on:click={drawerOpen}>
                    <span>
                        <svg viewBox="0 0 100 80" class="fill-token w-4 h-4">
                            <rect width="100" height="20"/>
                            <rect y="30" width="100" height="20"/>
                            <rect y="60" width="100" height="20"/>
                        </svg>
                    </span>
                    </button>
                </div>
            </svelte:fragment>

            <svelte:fragment slot="default">
                <section class="w-full justify-between items-center hidden lg:flex py-2">

                    <a href="/" class="">
                        <img class="min-w-16" src="/berlin_skylarks_logo.svg" alt="Skylarks Team Logo">
                    </a>

                    <Navigation></Navigation>

                    <!--ugly hack alert-->
                    <div></div>
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

    </header>

    <!-- Grid Column -->

    <div class="container grid grid-cols-1 md:grid-cols-[auto_1fr_auto]">

        <!-- Sidebar (Left) -->
        <aside class="bg-surface-500/5 p-4 sticky top-0 col-span-1 hidden h-screen lg:block">
            <div>TEXTdfszjdszjzdfj</div>
        </aside>

        <!-- Main -->

        <main class="col-span-1 space-y-4 mx-4 md:mx-6 lg:mx-8 mb-4 lg:mb-6">

            <slot/>

            <aside class="alert variant-ghost-error my-5">

                <ExclamationCircleOutline size="xl"/>

                <div class="alert-message">
                    <h3 class="h3">Alpha-Version</h3>
                    <p>Hier funktioniert noch nicht alles wie gewünscht. Fehler und merkwürdiges Verhalten sind zu
                        erwarten.</p>
                </div>
            </aside>

        </main>


    </div>


    <!-- Footer -->
    <footer>
        <hr class="!border-t-2">
        <Footer></Footer>
    </footer>
</div>