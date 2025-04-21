<script lang="ts">
  import '../app.postcss';
  import {AppBar, autoModeWatcher,} from "@skeletonlabs/skeleton";
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
  <!-- svelte-ignore hydration_html_changed -->
  {@html '<script nonce="%sveltekit.nonce%">(' + autoModeWatcher.toString() + ')();</script>'}
</svelte:head>

<!--Singletons-->
<ToastContainer/>

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
            <NavigationSheet clubs={data.clubs} teams={data.teams}/>
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
