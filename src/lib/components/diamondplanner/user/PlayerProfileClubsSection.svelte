<script lang="ts">
  import {ArrowLeftToBracketOutline, PlusOutline} from "flowbite-svelte-icons";
  import {type DrawerSettings, getDrawerStore} from "@skeletonlabs/skeleton";
  import type {ExpandedClub} from "$lib/model/ExpandedResponse";
  import ClubDetailCard from "$lib/components/diamondplanner/club/ClubDetailCard.svelte";

  interface Props {
    clubs: ExpandedClub[]
  }

  let {clubs}: Props = $props()

  const drawerStore = getDrawerStore()
  let clubAddEditSettings: DrawerSettings = $derived({
    id: "club-form",
    position: "right",
    width: "w-[100%] sm:w-[80%] lg:w-[70%] xl:w-[50%]",
    meta: {
      club: null,
      admins: null,
    },
  })

  function openDrawer(club?: ExpandedClub) {
    clubAddEditSettings.meta.club = club
    clubAddEditSettings.meta.admins = club?.expand?.admins
    drawerStore.open(clubAddEditSettings)
  }
</script>

{#each clubs as club}
    <ClubDetailCard {club}/>
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