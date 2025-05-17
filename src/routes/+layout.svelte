<script lang="ts">
  import '../app.css';
  import {AppBar} from "@skeletonlabs/skeleton-svelte";
  import SidebarNavigation from "$lib/components/meta/SidebarNavigation.svelte";
  import Footer from "$lib/components/meta/Footer.svelte";
  import LoginBadge from "$lib/auth/LoginBadge.svelte";
  import {env} from "$env/dynamic/public";
  import type {LayoutData} from "../../.svelte-kit/types/src/routes/$types";
  import StaticNavigationLinks from "$lib/components/navigation/StaticNavigationLinks.svelte";
  import BottomNavigation from "$lib/components/navigation/BottomNavigation.svelte";
  import {authSettings} from "$lib/pocketbase/index.svelte.ts";
  import NavigationSheet from "$lib/components/navigation/NavigationSheet.svelte";
  import ToastContainer from "$lib/components/utility/toast/ToastContainer.svelte";

  interface Props {
    data: LayoutData;
    children?: import('svelte').Snippet;
  }

  let {data, children}: Props = $props();

  let showSidebar = $derived(data.clubs.length > 0 || data.teams.length > 0);
  let isUserAuthenticated = $derived(!!authSettings.record);
</script>

<svelte:head>
  <title>Berlin Skylarks Web App</title>
</svelte:head>

<!--Singletons-->
<ToastContainer/>

<div class="h-screen grid grid-rows-[auto_1fr_auto]">
  <!-- Header -->
  <header>

    <AppBar
            centerAlign="place-self-center"
            padding="p-3"
            background="bg-surface-500/5 dark:bg-surface-800"
    >
      {#snippet lead()}
        <div class="flex items-center justify-content-start">
          {#if isUserAuthenticated}
            <NavigationSheet clubs={data.clubs} teams={data.teams}/>
          {/if}

          <a aria-label="to home page" href="/" class="hidden md:block ms-3">
            <img class="min-w-16" src="/berlin_skylarks_logo.svg" alt="Skylarks Team Logo">
          </a>
        </div>
      {/snippet}

      {#snippet children()}
        <section class="">
          <ul class="w-full justify-center items-center hidden lg:flex py-2 gap-2 xl:gap-16">
            <StaticNavigationLinks classes="rounded-base gap-2 py-1 px-2"/>
          </ul>
        </section>
      {/snippet}

      {#snippet trail()}
        <div class="lg:me-5 flex items-center gap-5 shrink-0">

          {#if env.PUBLIC_APPLICATION_CONTEXT !== "Production"}
            <div class="preset-tonal-warning p-1.5 rounded-container border border-warning-500 text-sm">
              {env.PUBLIC_APPLICATION_CONTEXT}
            </div>
          {/if}

          {#if env.PUBLIC_AUTH_FUNCS_ENABLED === "true"}
            <LoginBadge signupAllowed={true}/>
          {/if}

        </div>
      {/snippet}

    </AppBar>

    <hr>
  </header>

  <!-- Grid Column -->

  <div class="grid grid-cols-1 md:grid-cols-[auto_1fr]">

    <!-- Sidebar (Left) -->
    {#if showSidebar}
      <aside class="bg-surface-500/5 dark:bg-surface-800 p-2 sticky top-0 col-span-1 hidden h-screen md:block w-64 lg:w-72 xl:w-80 border-r border-surface-200 dark:border-surface-100">
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
  <footer class="pb-16 lg:pb-0">
    <hr>
    <Footer></Footer>
  </footer>
</div>
