<script lang="ts">
    import type {CustomAuthModel, ExpandedClub} from "$lib/model/ExpandedResponse";
    import {type DrawerSettings, getDrawerStore} from "@skeletonlabs/skeleton";
    import {authSettings} from "$lib/pocketbase/index.svelte";
    import {ClipboardPen, Info, Shield, Tag} from "lucide-svelte";

    interface Props {
    club: ExpandedClub;
  }

  let {club}: Props = $props();

  const drawerStore = getDrawerStore();

  const model = authSettings.record as CustomAuthModel;

  let clubAddEditSettings: DrawerSettings = $derived({
    id: "club-form",
    position: "right",
    width: "w-[100%] sm:w-[80%] lg:w-[70%] xl:w-[50%]",
    meta: {
      club: null,
      admins: null,
    },
  });

  function openDrawer(club?: ExpandedClub) {
    clubAddEditSettings.meta.club = club;
    clubAddEditSettings.meta.admins = club?.expand?.admins;
    drawerStore.open(clubAddEditSettings);
  }
</script>

<div class="card variant-glass-primary shadow-lg">
  <header class="card-header">
    <h2 class="h4 font-semibold">Club</h2>
  </header>

  <section class="p-4 space-y-2">
    <div class="flex items-center gap-3">
      <Shield/>
      <div>
        <p>{club.name}</p>
        <p class="text-sm font-light">Name</p>
      </div>
    </div>

    <div class="flex items-center gap-3">
      <Tag/>
      <div>
        <p>{club.acronym}</p>
        <p class="text-sm font-light">Acronym</p>
      </div>
    </div>

    <div class="flex items-center gap-3">
      <Info/>
      <div>
        <p>{club.bsm_id}</p>
        <p class="text-sm font-light">BSM-ID</p>
      </div>
    </div>
  </section>

  {#if club.admins.includes(model.id)}
    <footer class="card-footer flex justify-end">
      <button class="btn variant-ghost-secondary" onclick={() => openDrawer(club)}>
        <ClipboardPen/>
        <span>Edit Club</span>
      </button>
    </footer>
  {/if}
</div>