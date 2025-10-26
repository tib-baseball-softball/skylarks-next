<script lang="ts">
  import type {CustomAuthModel, ExpandedClub} from "$lib/model/ExpandedResponse";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import {Info, Shield, Tag} from "lucide-svelte";
  import ClubForm from "$lib/components/forms/ClubForm.svelte";

  interface Props {
    club: ExpandedClub;
  }

  let {club}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);
</script>

<div class="card bg-primary-100 dark:preset-tonal-primary-500 text-primary-900 shadow-lg">
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

  {#if club.admins.includes(authRecord.id)}
    <footer class="card-footer flex justify-end">
      <ClubForm club={club} buttonClasses="btn preset-tonal-secondary border border-secondary-500"/>
    </footer>
  {/if}
</div>