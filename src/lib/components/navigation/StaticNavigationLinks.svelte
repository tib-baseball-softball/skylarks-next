<script lang="ts">
  import {getDrawerStore} from "@skeletonlabs/skeleton";
  import {page} from "$app/state";
  import {Home, Shield, Table, Ticket, Users} from "lucide-svelte";

  interface Props {
    classes?: string;
  }

  let {classes = ""}: Props = $props();

  const drawerStore = getDrawerStore();

  function drawerClose(): void {
    drawerStore.close();
  }
</script>

<li>
  <a class="{classes}" class:variant-filled-primary={page.url.pathname === "/"} href="/" onclick={drawerClose}>
    <Home/>
    <span>Start</span>
  </a>
</li>

<li>
  <a class="{classes}" class:variant-filled-primary={page.url.pathname.startsWith("/gamecenter")} href="/gamecenter"
     onclick={drawerClose}>
    <Ticket/>
    <span>Games</span>
  </a>
</li>
<li>
  <a class="{classes}" class:variant-filled-primary={page.url.pathname.startsWith("/ligen")} href="/ligen"
     onclick={drawerClose}>
    <Table/>
    <span>Leagues</span>
  </a>
</li>
<li>
  <a class="{classes}" class:variant-filled-primary={page.url.pathname.startsWith("/teams")} href="/teams"
     onclick={drawerClose}>
    <Users/>
    <span>Teams</span>
  </a>
</li>
<li>
  <a class="{classes}" class:variant-filled-primary={page.url.pathname.startsWith("/club")} href="/club"
     onclick={drawerClose}>
    <Shield/>
    <span>Club</span>
  </a>
</li>

<style lang="postcss">
    a {
        display: flex;
        align-items: center;
        justify-content: center;
        @apply hover:variant-glass-primary;
    }

    li {
        width: 20%;
    }
</style>