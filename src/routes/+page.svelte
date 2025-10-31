<script lang="ts">
  import CurrentMatchesOverviewCard from "$lib/components/match/CurrentMatchesOverviewCard.svelte";
  import type {PageProps} from "./$types";

  let {data}: PageProps = $props();
</script>

<div class="w-full space-y-5 mt-2 mb-5">
  <section class="hero w-screen h-[20em] rounded-lg" id="Hero">
    <div class="hero-overlay bg-opacity-60"></div>
    <div class="hero-content">
      <div class="max-w-md text-white preset-glass-neutral p-3 rounded-xl">
        <h1 class="h1 mb-2">Diamond Planner</h1>
        <p>
          Berlin Skylarks Team App
        </p>
      </div>
    </div>
  </section>

  <section class="prose dark:prose-invert">
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
            matchesCurrent={data.streamed.matchesCurrent}
            matchesNext={data.streamed.matchesNext}
            matchesPrevious={data.streamed.matchesPrevious}
    />
  </section>

</div>

<style>
    .hero {
        display: grid;
        width: 100%;
        place-items: start;
        background-size: cover;
        background-position: center;
        background-image: url("/header_hero.jpg");
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
</style>
