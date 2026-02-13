<script lang="ts">
  import {SquareArrowOutUpRight} from "lucide-svelte";
  import ClubDetailCard from "$lib/dp/components/club/ClubDetailCard.svelte";
  import ClubForm from "$lib/dp/components/forms/ClubForm.svelte";
  import type {ExpandedClub} from "$lib/dp/types/ExpandedResponse.ts";

  interface Props {
    clubs: ExpandedClub[];
  }

  let {clubs}: Props = $props();
</script>

{#each clubs as club}
  <ClubDetailCard {club}/>
{/each}

{#if !clubs}
  <div class="card preset-tonal-primary shadow-lg">
    <header class="card-header">
      <h2 class="h4 title">Club</h2>
    </header>

    <section class="content">
      <div class="message">
        <p>You are not a member of any clubs yet.</p>
      </div>
    </section>

    <footer class="card-footer">
      <div class="actions">

        <ClubForm club={null} buttonClasses="btn preset-tonal-primary border border-primary-500"/>

        <!--TODO: This doesn't do anything yet-->
        <button class="btn preset-tonal-primary border border-primary-500">
          <SquareArrowOutUpRight/>
          <span>Join a Club</span>
        </button>

      </div>
    </footer>
  </div>
{/if}

<style>
  .title {
    font-weight: var(--font-weight-semibold);
  }

  .content {
    padding: calc(var(--spacing) * 4);
  }

  .message {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 3);
  }

  .actions {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: calc(var(--spacing) * 3);
  }
</style>