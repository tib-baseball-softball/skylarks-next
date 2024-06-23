<script lang="ts">
    import '../app.postcss';
    import {AppBar} from "@skeletonlabs/skeleton";
    import {AppShell} from "@skeletonlabs/skeleton";
    import Navigation from "$lib/components/meta/Navigation.svelte";
    import { LightSwitch } from '@skeletonlabs/skeleton';
    import {initializeStores, Drawer, getDrawerStore} from '@skeletonlabs/skeleton';
    import Footer from "$lib/components/meta/Footer.svelte";
    import {preferences} from "$lib/stores";
    import {ExclamationCircleOutline} from "flowbite-svelte-icons";

    initializeStores();

    const drawerStore = getDrawerStore();

    function drawerOpen(): void {
        drawerStore.open({});
    }
</script>

<Drawer width="w-[70%]">
    <h2 class="p-4">Berlin Skylarks</h2>
    <hr class="mb-2"/>
    <Navigation></Navigation>
</Drawer>

<AppShell
        slotSidebarLeft="bg-surface-500/5 w-0 md:w-64"
        regionPage="relative"
        slotPageHeader="sticky top-0 z-10"
>
    <svelte:fragment slot="header">
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
                    <a href="/">
                        <img class="min-w-16" src="/berlin_skylarks_logo.svg" alt="Skylarks Team Logo">
                    </a>
                    <Navigation></Navigation>
                    <a href="https://tib1848ev.de/" target="_blank">
                        <img class="min-w-8 max-w-14" src="/tib_logo.svg" alt="TiB Logo">
                    </a>
                </section>
            </svelte:fragment>
            <svelte:fragment slot="trail">
                <div class="me-5 flex gap-5">
                    <p class="flex-shrink-0">Saison: {$preferences.selectedSeason}</p>
                    <div class="me-5">
                        <LightSwitch/>
                    </div>
                </div>
            </svelte:fragment>
        </AppBar>
    </svelte:fragment>
    <hr class="!border-t-2">

    <!-- (Default Page Content slot) -->
    <div id="pageContainer" class="flex items-center justify-center">
        <div class="flex flex-col justify-center content-center w-[92%] md:w-[85%] xl:w-[75%]">
            <slot/>

            <aside class="alert variant-ghost-error">
                <ExclamationCircleOutline size="xl"/>
                <div class="alert-message">
                    <h3 class="h3">Alpha-Version</h3>
                    <p>Hier funktioniert noch nicht alles wie gewünscht. Fehler und merkwürdiges Verhalten sind zu erwarten.</p>
                </div>
            </aside>
        </div>
    </div>

    <svelte:fragment slot="pageFooter">
        <hr class="!border-t-2">
        <Footer></Footer>
    </svelte:fragment>
</AppShell>