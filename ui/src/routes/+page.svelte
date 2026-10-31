<svelte:head>
  <title>Diamond Planner</title>
</svelte:head>

<div class="hero-wrapper">
  <section class="hero" id="Hero">
    <div class="hero-overlay"></div>
    <div class="hero-content">
      <div class="hero-text preset-glass-neutral">
        <h1 class="h1">Diamond Planner</h1>
        <p>Berlin Skylarks Team App</p>
      </div>
    </div>
  </section>

  <section class="prose">
    <p>
      The core functionality of this app (getting up-to-date information about
      games, standings and public club information) is available without logging
      in.
    </p>
    <p class="not-prose">
      <a
        class="btn preset-tonal-primary border-primary"
        href="/bsm"
        title="BSM Data">Go to game data</a
      >
    </p>
    <p>
      To access internal team information and participate in team activities,
      please
      <a href="/login"> log in to your account. </a>
    </p>
  </section>
</div>

<div class="cards">
  <article class="card card-hover preset-outlined-card shadow">
    <header>
      <h2 class="h4">Help & Documentation</h2>
    </header>
    <p>
      Find usage information for this app and frequently asked questions on the Skylarks team
      website.
    </p>
    <footer>
      <a
        class="btn preset-tonal-primary border-primary"
        href="https://www.tib-baseball.de/diamond-planner-dokumentation"
        target="_blank"
        title="Read the official documentation on how to use this app, opens in new tab"
        >Get started</a
      >
    </footer>
  </article>

  <article class="card card-hover preset-outlined-card shadow">
    <header>
      <h2 class="h4">Install as web app</h2>
    </header>
    <p>
      Find out how you can install this app to your home screen on touch devices
      and to the dock on select desktop browsers.
    </p>
    <footer>
      <a
        class="btn preset-tonal-tertiary border-tertiary"
        href="https://www.tib-baseball.de/diamond-planner-dokumentation/installation-als-web-app"
        target="_blank"
        title="Installation instructions as a progressive web app"
        >Installation instructions</a
      >
    </footer>
  </article>
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
    margin-block: calc(var(--spacing) * 10);
  }

  .cards {
    display: grid;
    grid-template-columns: 1fr;
    gap: var(--card-padding);

    @media (min-width: 32rem) {
      grid-template-columns: 1fr 1fr;
    }

    .card {
      padding: var(--card-padding);
    }

    p {
      padding-block: calc(var(--spacing) * 4);
    }
  }
</style>
