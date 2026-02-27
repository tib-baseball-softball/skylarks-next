<script lang="ts">
  import CurrentMatchesOverviewCard from "$lib/tib/components/match/CurrentMatchesOverviewCard.svelte";
  import type {PageProps} from "./$types";

  let {data}: PageProps = $props();
</script>

<div class="hero-wrapper">
  <section class="hero" id="Hero">
    <div class="hero-overlay"></div>
    <div class="hero-content">
      <div class="hero-text preset-glass-neutral">
        <h1 class="h1">Diamond Planner</h1>
        <p>
          Berlin Skylarks Team App
        </p>
      </div>
    </div>
  </section>

  <section class="prose">
    <p>
      The core functionality of this app (getting up-to-date information about games, standings and public club
      information) is available without logging in.
    </p>
    <p>To access internal team information and participate in team activities, please
      <a href="/login">
        log in to your account.
      </a>
    </p>
  </section>

  <section id="Scores">
    <CurrentMatchesOverviewCard
      matchesCurrent={data.matchesCurrent}
      matchesNext={data.matchesNext}
      matchesPrevious={data.matchesPrevious}
    />
  </section>

</div>

<style>
  .hero-wrapper {
    width: 100%;
    margin-block-start: calc(var(--spacing) * 2);
    margin-block-end: calc(var(--spacing) * 5);
  }

  .hero {
    display: grid;
    width: 100%;
    place-items: start;
    background-size: cover;
    background-position: center;
    background-image: url("/header_hero.jpg");
    border-radius: var(--radius-lg);
    height: 20em;
  }

  .hero-content {
    z-index: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    max-width: 80rem;
    gap: 1rem;
    padding: 1rem;
  }

  .hero-text {
    color: white;
    padding: calc(var(--spacing) * 3);
    border-radius: var(--radius-lg);
  }

  .prose {
    margin-block: calc(var(--spacing) * 6);
  }
</style>
