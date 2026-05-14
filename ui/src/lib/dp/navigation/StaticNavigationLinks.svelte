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
    <span class="link-label">Start</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname === "/account"}
     href="/account">
    <User/>
    <span class="link-label">Account</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname === "/account/clubs"}
     href="/account/clubs">
    <Shield/>
    <span class="link-label">Clubs</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname === "/account/team"}
     href="/account/team">
    <Users/>
    <span class="link-label">Teams</span>
  </a>
</li>

<li>
  <a class="{classes}" class:preset-filled-primary-500={page.url.pathname === "/account/settings"}
     href="/account/settings">
    <Settings/>
    <span class="link-label">Settings</span>
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

  .link-label {
    max-width: calc(var(--spacing) * 16);
    overflow: clip;
    text-overflow: ellipsis;
    font-size: var(--text-sm);
  }
</style>
