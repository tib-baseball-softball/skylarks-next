<script lang="ts">
  import CurrentMatchesOverviewCard from "$lib/components/match/CurrentMatchesOverviewCard.svelte";
  import type {PageProps} from "./$types";

  let {data}: PageProps = $props();
</script>

<div class="w-full space-y-5 mt-2 mb-5">
  <section id="Hero" class="hero w-screen h-[20em] rounded-lg">
    <div class="hero-overlay bg-opacity-60"></div>
    <div class="hero-content">
      <div class="max-w-md text-white preset-glass-neutral p-3 rounded-xl">
        <h1 class="h1 mb-2">Berlin Skylarks</h1>
        <p>
          Baseball & Softball in the Heart of Berlin
        </p>
      </div>
    </div>
  </section>

  <section id="Scores">
    <CurrentMatchesOverviewCard
            matchesCurrent={data.streamed.matchesCurrent}
            matchesPrevious={data.streamed.matchesPrevious}
            matchesNext={data.streamed.matchesNext}
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
