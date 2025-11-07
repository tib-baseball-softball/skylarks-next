<script lang="ts">
import "../app.css"
import AppBar from "$lib/components/meta/AppBar.svelte"
import Footer from "$lib/components/meta/Footer.svelte"
import SidebarNavigation from "$lib/components/meta/SidebarNavigation.svelte"
import TopAppBarTrailing from "$lib/components/meta/TopAppBarTrailing.svelte"
import BottomNavigation from "$lib/components/navigation/BottomNavigation.svelte"
import NavigationSheet from "$lib/components/navigation/NavigationSheet.svelte"
import StaticNavigationLinks from "$lib/components/navigation/StaticNavigationLinks.svelte"
import ToastContainer from "$lib/components/utility/toast/ToastContainer.svelte"
import { authSettings } from "$lib/pocketbase/index.svelte.ts"
import type { LayoutProps } from "../../.svelte-kit/types/src/routes/$types"

const { data, children }: LayoutProps = $props()

const isUserAuthenticated = $derived(!!authSettings.record)
const showSidebar = $derived(
  (data.clubs.length > 0 || data.teams.length > 0) && isUserAuthenticated
)
</script>

<svelte:head>
  <title>Berlin Skylarks Web App</title>
</svelte:head>

<!--Singletons-->
<ToastContainer/>

<div class="root-layout">
  <!-- Header -->
  <div>

    <AppBar>
      {#snippet lead()}
        <div class="app-bar-lead">
          {#if isUserAuthenticated}
            <NavigationSheet clubs={data.clubs} teams={data.teams}/>
          {/if}

          <a aria-label="to home page" href="/" class="logo-link">
            <img class="team-logo" src="/berlin_skylarks_logo.svg" alt="Skylarks Team Logo">
          </a>
        </div>
      {/snippet}

      {#snippet children()}
        <section>
          <ul class="app-bar-link-list ">
            <StaticNavigationLinks classes="rounded-base gap-2 py-1 px-2"/>
          </ul>
        </section>
      {/snippet}

      {#snippet trail()}
        <div class="app-bar-trail">

          <TopAppBarTrailing/>

        </div>
      {/snippet}

    </AppBar>

    <hr>
  </div>

  <!-- Grid Column -->

  <div class="sidebar-grid">

    <!-- Sidebar (Left) -->
    {#if showSidebar}
      <aside class="">
        <SidebarNavigation clubs={data.clubs} teams={data.teams}/>
      </aside>
    {:else}
      <!-- hack: render empty div to not mess up the grid -->
      <div class="shame-div" aria-hidden="true"></div>
    {/if}

    <!-- Main -->

    <main class="">
      {@render children?.()}
      <noscript>This website is actually a JavaScript application with lots of functionality. You need to enable
        JavaScript for it to work.
      </noscript>
    </main>
  </div>

  <!-- Bottom Nav (Fixed to bottom of page) for mobile -->

  <BottomNavigation/>

  <!-- Footer -->
  <footer class="app-footer">
    <hr>
    <Footer></Footer>
  </footer>
</div>

<style>
    .root-layout {
        height: 100%;
        display: grid;
        grid-auto-rows: auto 1fr auto;
    }

    .app-bar-lead {
        display: flex;
        justify-content: start;
        align-items: center;
    }

    .app-bar-link-list {
        width: 100%;
        justify-content: center;
        align-items: center;
        display: none;
        padding-block: calc(var(--spacing) * 2);
        gap: calc(var(--spacing) * 2);

        @media (min-width: 48rem) {
            display: flex;
        }

        @media (min-width: 80rem) {
            gap: calc(var(--spacing) * 16);
        }
    }

    .app-bar-trail {
        display: flex;
        align-items: center;
        gap: calc(var(--spacing) * 5);
        flex-shrink: 0;

        @media (min-width: 64rem) {
            margin-inline-end: calc(var(--spacing) * 5);
        }
    }

    .sidebar-grid {
        display: grid;
        grid-template-columns: repeat(1, minmax(0, 1fr));

        @media (min-width: 48rem) {
            grid-template-columns: auto 1fr;
        }
    }

    aside {
        position: sticky;
        top: 0;
        display: none;
        height: 100vh;
        width: calc(var(--spacing) * 64);
        border-right: 1px solid light-dark(var(--color-surface-200), var(--color-surface-100));
        padding: calc(var(--spacing) * 2);
        background-color: var(--nav-item-background);
        grid-column: span 1/span 1;

        @media (min-width: 48rem) {
            display: block;
        }

        @media (min-width: 64rem) {
            width: calc(var(--spacing) * 72);
        }

        @media (min-width: 80rem) {
            width: calc(var(--spacing) * 80);
        }
    }

    main {
        max-width: 1200px;
        width: 93%;
        justify-self: center;
        margin-bottom: 2em;
        grid-column: span 1/span 1;

        @media (min-width: 48rem) {
            width: 90%;
        }

        @media (min-width: 64rem) {
            width: 85%;
        }
    }

    .app-footer {
        padding-bottom: calc(var(--spacing) * 16);

        @media (min-width: 64rem) {
            padding-bottom: 0;
        }
    }

    .logo-link {
        display: none;
        margin-inline-start: 1.25em;

        @media (min-width: 48em) {
            display: block;
        }
    }

    .team-logo {
        min-width: calc(var(--spacing) * 16);
    }

    [aria-hidden="true"] {
        display: none;

        @media (min-width: 48em) {
            display: block;
        }
    }
</style>