<script lang="ts">
  import type {Club} from "bsm.js";
  import {ClipboardListOutline, LabelOutline, ShieldOutline} from "flowbite-svelte-icons";

  interface Props {
    club: Club
  }

  let {club}: Props = $props();
</script>

<h1 class="h1">Club</h1>
<div class="flex justify-center">
  <img
          src="/berlin_skylarks_logo.svg"
          alt="Skylarks Logo"
          class="w-24 h-auto mb-1"
  />
</div>

<div class="card p-3 variant-ghost shadow-md">
  <div class="container">
    <ClipboardListOutline size="lg"/>
    <span>{club.short_name}</span>
  </div>

  <hr class="my-2">

  <div class="container">
    <LabelOutline size="lg"/>
    <span>
        {club.acronym} / 0{club?.organization_id} {club.number}
      </span>
  </div>

  <hr class="my-2">

  <div class="container">
    <ShieldOutline size="lg"/>
    <span>{club.main_club}</span>
  </div>
</div>

<style lang="postcss">
    .container {
        display: flex;
        @apply gap-4;
    }
</style>