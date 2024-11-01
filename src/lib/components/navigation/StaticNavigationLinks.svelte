<script lang="ts">
  import {HomeOutline, ShieldOutline, TableRowOutline, TicketOutline, UsersGroupOutline} from "flowbite-svelte-icons";
  import {getDrawerStore} from "@skeletonlabs/skeleton";

  const drawerStore = getDrawerStore();

  function drawerClose(): void {
    drawerStore.close();
  }
</script>

<li>
    <a href="/" onclick={drawerClose}>
        <HomeOutline size="lg"/>
        <span>Start</span>
    </a>
</li>

<!--        <li><a href="/aktuelles" on:click={drawerClose}>Aktuelles</a></li>-->
<li>
    <a href="/gamecenter" onclick={drawerClose}>
        <TicketOutline size="lg"/>
        <span>Gamecenter</span>
    </a>
</li>
<li>
    <a href="/ligen" onclick={drawerClose}>
        <TableRowOutline size="lg"/>
        <span>Leagues</span>
    </a>
</li>
<li>
    <a href="/teams" onclick={drawerClose}>
        <UsersGroupOutline size="lg"/>
        <span>Teams</span>
    </a>
</li>
<li>
    <a href="/club" onclick={drawerClose}>
        <ShieldOutline size="lg"/>
        <span>Club</span>
    </a>
</li>
<!--        <li><a href="/kontakt" on:click={drawerClose}>Contact</a></li>-->

<style lang="postcss">
    a {
        display: flex;
        @apply gap-2;
        align-items: center;
        @apply hover:variant-glass-primary;
        padding: 0.5rem 1rem;
        @apply rounded-token;
    }
</style>