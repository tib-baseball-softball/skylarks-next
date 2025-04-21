<script lang="ts">
  import {page} from "$app/state";
  import {Home, Shield, Table, Ticket, Users} from "lucide-svelte";

  interface Props {
    classes?: string;
  }

  let {classes = ""}: Props = $props();
</script>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname === "/"} href="/">
    <Home/>
    <span>Start</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname.startsWith("/gamecenter")} href="/gamecenter">
    <Ticket/>
    <span>Games</span>
  </a>
</li>
<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname.startsWith("/ligen")} href="/ligen">
    <Table/>
    <span>Leagues</span>
  </a>
</li>
<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname.startsWith("/teams")} href="/teams">
    <Users/>
    <span>Teams</span>
  </a>
</li>
<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname.startsWith("/club")} href="/club">
    <Shield/>
    <span>Club</span>
  </a>
</li>

<style lang="postcss">
    a {
        display: flex;
        align-items: center;
        justify-content: center;
    }

    a:hover {
        background-color: var(--color-primary-50-950);
        color: var(--color-primary-950-50);
    }

    li {
        width: 20%;
    }
</style>