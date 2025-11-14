<script lang="ts">
  import {House, Shield, Star, Table, Ticket, Users} from "lucide-svelte";
  import {page} from "$app/state";
  import {preferences} from "$lib/tib/globals.svelte.ts";

  interface Props {
    classes?: string;
  }

  const {classes = ""}: Props = $props();
</script>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname === "/"}
     href="/">
    <House/>
    <span>Start</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname === "/favorite"}
     href="/favorite?season={preferences.current.selectedSeason}">
    <Star/>
    <span>Favorite</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname.startsWith("/gamecenter")}
     href="/gamecenter">
    <Ticket/>
    <span>Games</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname.startsWith("/ligen")}
     href="/ligen">
    <Table/>
    <span>Leagues</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname.startsWith("/teams")}
     href="/teams">
    <Users/>
    <span>Teams</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname.startsWith("/club")}
     href="/club">
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