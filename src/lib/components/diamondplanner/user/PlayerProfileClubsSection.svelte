<script lang="ts">
  import type {ClubsResponse} from "$lib/model/pb-types";
  import {
    ArrowLeftToBracketOutline,
    InfoCircleOutline,
    PlusOutline,
    ShieldOutline,
    TagOutline
  } from "flowbite-svelte-icons";
  import {type DrawerSettings, getDrawerStore} from "@skeletonlabs/skeleton";
  import {authModel} from "$lib/pocketbase/Auth";
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";

  interface Props {
    clubs: ClubsResponse[]
  }

  const model = $authModel as CustomAuthModel
  let {clubs}: Props = $props()

  const drawerStore = getDrawerStore()
  let clubAddEditSettings: DrawerSettings = $derived({
    id: "club-form",
    position: "right",
    width: "w-[100%] sm:w-[80%] lg:w-[70%] xl:w-[50%]",
    meta: {
      club: null
    },
  })

  function openDrawer(club?: ClubsResponse) {
    clubAddEditSettings.meta.club = club
    drawerStore.open(clubAddEditSettings)
  }
</script>

{#each clubs as club}
    <div class="card variant-glass-primary shadow-lg">
        <header class="card-header">
            <h2 class="h4 font-semibold">Club</h2>
        </header>

        <section class="p-4 space-y-2">
            <div class="flex items-center gap-3">
                <ShieldOutline/>
                <div>
                    <p>{club.name}</p>
                    <p class="text-sm font-light">Name</p>
                </div>
            </div>

            <div class="flex items-center gap-3">
                <TagOutline/>
                <div>
                    <p>{club.acronym}</p>
                    <p class="text-sm font-light">Acronym</p>
                </div>
            </div>

            <div class="flex items-center gap-3">
                <InfoCircleOutline/>
                <div>
                    <p>{club.bsm_id}</p>
                    <p class="text-sm font-light">BSM-ID</p>
                </div>
            </div>
        </section>

        {#if club.admins.includes(model.id)}
            <footer class="card-footer flex justify-end">
                <button class="btn variant-ghost-secondary" onclick={() => openDrawer(club)}>
                    Edit Club
                </button>
            </footer>
        {/if}
    </div>
{/each}

{#if !clubs}
    <div class="card variant-glass-primary shadow-lg">
        <header class="card-header">
            <h2 class="h4 font-semibold">Club</h2>
        </header>

        <section class="p-4 space-y-2">
            <div class="flex items-center gap-3">
                <p>You are not a member of any clubs yet.</p>
            </div>
        </section>

        <footer class="card-footer">
            <div class="flex flex-wrap items-center gap-3">
                <button class="btn variant-ghost-primary" onclick={() => openDrawer()}>
                    <PlusOutline/>
                    <span>Create a Club</span>
                </button>
                <button class="btn variant-ghost-primary">
                    <ArrowLeftToBracketOutline/>
                    <span>Join a Club</span>
                </button>
            </div>
        </footer>
    </div>
{/if}