<script lang="ts">
  import {Info, Shield, Tag} from "lucide-svelte";
  import ClubForm from "$lib/dp/components/forms/ClubForm.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedClub} from "$lib/dp/types/ExpandedResponse.ts";

  interface Props {
    club: ExpandedClub;
  }

  const {club}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);
</script>

<div class="card shadow-lg">
  <header class="card-header">
    <h2 class="title h4">Club</h2>
  </header>

  <section class="content">
    <div class="item">
      <Shield/>
      <div>
        <p>{club.name}</p>
        <p class="light-text">Name</p>
      </div>
    </div>

    <div class="item">
      <Tag/>
      <div>
        <p>{club.acronym}</p>
        <p class="light-text">Acronym</p>
      </div>
    </div>

    <div class="item">
      <Info/>
      <div>
        <p>{club.bsm_id}</p>
        <p class="light-text">BSM-ID</p>
      </div>
    </div>
  </section>

  {#if club.admins.includes(authRecord.id)}
    <footer class="card-footer">
      <ClubForm club={club} buttonClasses="btn preset-tonal-secondary border border-secondary-500"/>
    </footer>
  {/if}
</div>

<style>
  .card {
    background-color: light-dark(var(--color-primary-100), var(--color-primary-950));
  }

  .title {
    font-weight: 600;
  }

  .content {
    padding: calc(var(--spacing) * 4);

    & > * + * {
      margin-top: calc(var(--spacing) * 2);
    }
  }

  .item {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 3);
  }

  .light-text {
    font-weight: var(--font-weight-light);
    font-size: var(--text-sm);
    line-height: var(--tw-leading, var(--text-sm--line-height));
  }

  .card-footer {
    display: flex;
    justify-content: end;
  }
</style>
