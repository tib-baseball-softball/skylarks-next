<script lang="ts">
  import '../app.postcss';
  import {
    AppBar,
    autoModeWatcher,
    Drawer,
    getDrawerStore,
    initializeStores,
    Modal,
    type ModalComponent,
    Toast,
  } from "@skeletonlabs/skeleton";
  import SidebarNavigation from "$lib/components/meta/SidebarNavigation.svelte";
  import Footer from "$lib/components/meta/Footer.svelte";
  import LoginBadge from "$lib/auth/LoginBadge.svelte";
  import LoginForm from "$lib/auth/LoginForm.svelte";
  import AccountModal from "$lib/auth/AccountModal.svelte";
  import {env} from "$env/dynamic/public";
  import EventForm from '$lib/components/forms/EventForm.svelte';
  import TeamForm from '$lib/components/forms/TeamForm.svelte';
  import PlayerDataForm from '$lib/components/forms/PlayerDataForm.svelte';
  import ClubForm from "$lib/components/forms/ClubForm.svelte";
  import type {LayoutData} from "../../.svelte-kit/types/src/routes/$types";
  import StaticNavigationLinks from "$lib/components/navigation/StaticNavigationLinks.svelte";
  import EventSeriesView from "$lib/components/diamondplanner/event/EventSeriesView.svelte";
  import BottomNavigation from "$lib/components/navigation/BottomNavigation.svelte";
  import {authSettings} from "$lib/pocketbase/index.svelte.ts";

  interface Props {
    data: LayoutData;
    children?: import('svelte').Snippet;
  }

  let {data, children}: Props = $props();

  initializeStores();

  const modalRegistry: Record<string, ModalComponent> = {
    // Set a unique modal ID, then pass the component reference
    loginForm: {ref: LoginForm},
    accountOverview: {ref: AccountModal}
  };

  const drawerStore = getDrawerStore();

  function navDrawerOpen(): void {
    drawerStore.open({
      id: "nav",
      width: "w-[70%] sm:w-[40%]"
    });
  }

  let showSidebar = $derived(data.clubs.length > 0 || data.teams.length > 0);
  let isUserAuthenticated = $derived(!!authSettings.record);
</script>

<svelte:head>
  <title>Berlin Skylarks Web App</title>
  <!-- svelte-ignore hydration_html_changed -->
  {@html '<script>(' + autoModeWatcher.toString() + ')();</script>'}
</svelte:head>

<Drawer>
  {#if $drawerStore.id === "nav"}

    <div class="flex justify-around p-2">
      <img class="max-w-14" src="/berlin_skylarks_logo.svg" alt="Skylarks Team Logo">

      <h2 class="p-4 antialiased">Berlin Skylarks</h2>
    </div>

    <hr class="mb-2"/>

    <SidebarNavigation clubs={data.clubs} teams={data.teams}/>

  {:else if $drawerStore.id === "event-form"}
    <EventForm/>
  {:else if $drawerStore.id === "team-form"}
    <TeamForm/>
  {:else if $drawerStore.id === "player-data-form"}
    <PlayerDataForm/>
  {:else if $drawerStore.id === "club-form"}
    <ClubForm/>
  {:else if $drawerStore.id === "eventseries-view"}
    <EventSeriesView/>
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
          {#if isUserAuthenticated}
            <button aria-label="open navigation" class="md:hidden btn btn-sm mr-4" onclick={navDrawerOpen}>
                    <span>
                        <svg viewBox="0 0 100 80" class="fill-token w-4 h-4">
                            <rect width="100" height="20"/>
                            <rect y="30" width="100" height="20"/>
                            <rect y="60" width="100" height="20"/>
                        </svg>
                    </span>
            </button>
          {/if}

          <a aria-label="to home page" href="/" class="hidden md:block ms-3">
            <img class="min-w-16" src="/berlin_skylarks_logo.svg" alt="Skylarks Team Logo">
          </a>
        </div>
      </svelte:fragment>

      <svelte:fragment slot="default">
        <section class="">
          <ul class="w-full justify-center items-center hidden lg:flex py-2 gap-2 xl:gap-16">
            <StaticNavigationLinks classes="rounded-token gap-2 py-1 px-2"/>
          </ul>
        </section>
      </svelte:fragment>

      <svelte:fragment slot="trail">
        <div class="lg:me-5 flex items-center gap-5 flex-shrink-0">

          {#if env.PUBLIC_AUTH_FUNCS_ENABLED === "true"}
            <LoginBadge signupAllowed={true}/>
          {/if}

        </div>
      </svelte:fragment>

    </AppBar>

    <hr class="!border-b-2">
  </header>

  <!-- Grid Column -->

  <div class="grid grid-cols-1 md:grid-cols-[auto_1fr]">

    <!-- Sidebar (Left) -->
    {#if showSidebar}
      <aside class="bg-surface-500/5 p-2 sticky top-0 col-span-1 hidden h-screen md:block w-64 lg:w-72 xl:w-80">
        <SidebarNavigation clubs={data.clubs} teams={data.teams}/>
      </aside>
    {:else}
      <!-- hack: render empty div to not mess up the grid -->
      <div aria-hidden="true" class="hidden md:block"></div>
    {/if}

    <!-- Main -->

    <main class="col-span-1 space-y-4 lg:space-y-6 w-[93%] md:w-[90%] lg:w-[85%] justify-self-center">
      {@render children?.()}
      <noscript>This website is actually a JavaScript application with lots of functionality. You need to enable
        JavaScript for it to work.
      </noscript>
    </main>
  </div>

  <!-- Bottom Nav (Fixed to bottom of page) for mobile -->

  <BottomNavigation/>

  <!-- Footer -->
  <footer>
    <hr class="!border-t-2">
    <Footer></Footer>
  </footer>
</div>
