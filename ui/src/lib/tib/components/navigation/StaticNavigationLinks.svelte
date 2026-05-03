<script lang="ts">
  import {House, Settings, Shield, User, Users} from "lucide-svelte";
  import {page} from "$app/state";

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
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname === "/account"}
     href="/account">
    <User/>
    <span>Account</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname === "/account/clubs"}
     href="/account/clubs">
    <Shield/>
    <span>Clubs</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname === "/account/team"}
     href="/account/team">
    <Users/>
    <span>Teams</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname === "/account/settings"}
     href="/account/settings">
    <Settings/>
    <span>Settings</span>
  </a>
</li>

<style>
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
