<script lang="ts">
  import {HomeOutline, ShieldOutline, TableRowOutline, TicketOutline, UsersGroupOutline} from "flowbite-svelte-icons";
  import {getDrawerStore} from "@skeletonlabs/skeleton";
  import {page} from "$app/state";

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
    <HomeOutline size="lg"/>
    <span>Start</span>
  </a>
</li>

<li>
  <a class="{classes}" class:variant-filled-primary={page.url.pathname.startsWith("/gamecenter")} href="/gamecenter"
     onclick={drawerClose}>
    <TicketOutline size="lg"/>
    <span>Games</span>
  </a>
</li>
<li>
  <a class="{classes}" class:variant-filled-primary={page.url.pathname.startsWith("/ligen")} href="/ligen"
     onclick={drawerClose}>
    <TableRowOutline size="lg"/>
    <span>Leagues</span>
  </a>
</li>
<li>
  <a class="{classes}" class:variant-filled-primary={page.url.pathname.startsWith("/teams")} href="/teams"
     onclick={drawerClose}>
    <UsersGroupOutline size="lg"/>
    <span>Teams</span>
  </a>
</li>
<li>
  <a class="{classes}" class:variant-filled-primary={page.url.pathname.startsWith("/club")} href="/club"
     onclick={drawerClose}>
    <ShieldOutline size="lg"/>
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