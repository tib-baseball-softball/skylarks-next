<script lang="ts">
  import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";
  import {
    UserOutline,
    EnvelopeOutline,
    ShieldOutline,
    TagOutline,
    InfoCircleOutline,
  } from "flowbite-svelte-icons";
  import { authModel } from "$lib/pocketbase/Auth";
  import type { CustomAuthModel } from "$lib/model/ExpandedResponse.js";
  import {
    Avatar,
    type ConicStop,
    ConicGradient,
  } from "@skeletonlabs/skeleton";
  import { client } from "$lib/pocketbase";
  import StatsBlockContent from "$lib/components/utility/StatsBlockContent.svelte";
  import type { SingleStatElement } from "$lib/types/SingleStatElement.js";

  const model = $authModel as CustomAuthModel;

  let { data } = $props();
  //console.log(model);

  const conicStops: ConicStop[] = [
    {
      label: "In",
      color: "rgb(var(--color-success-500))",
      start: 0,
      end: 33,
    },
    {
      label: "Maybe",
      color: "rgb(var(--color-warning-500))",
      start: 33,
      end: 66,
    },
    {
      label: "Out",
      color: "rgb(var(--color-error-500))",
      start: 66,
      end: 100,
    },
  ];

  const block: SingleStatElement = {
    title: "Games",
    value: "10 / 16",
    desc: "All possible games for 2024 season",
  };

  const block2: SingleStatElement = {
    title: "Practices",
    value: "23 / 45",
    desc: "All possible practices for your teams (2024 season)",
  };
</script>

<h1 class="h1 lg:mt-4">My Dashboard</h1>

<div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
  <div class="card variant-glass-surface lg:col-span-2 shadow-lg">
    <header class="card-header">
      <h2 class="h4 font-semibold">User Data</h2>
    </header>

    <section class="p-4 flex flex-col sm:flex-row gap-4 lg:gap-3 sm:items-center">
      <Avatar src={client.files.getUrl(model, model?.avatar)} />

      <div class="grid grid-cols-1 gap-2 col-span-4">
        <div class="flex items-center gap-3">
          <UserOutline size="lg" />
          <div>
            <p>{`${model?.first_name ?? ""} ${model?.last_name ?? ""}`}</p>
            <p class="text-sm font-light">Name</p>
          </div>
        </div>

        <div class="flex items-center gap-3">
          <EnvelopeOutline size="lg" />
          <div>
            <p>{model?.email}</p>
            <p class="text-sm font-light">email address</p>
          </div>
        </div>
      </div>
    </section>

    <footer class="card-footer flex justify-end">
      <button class="btn variant-ghost-primary">Edit</button>
    </footer>
  </div>

  {#each model.expand.club as club}
    <div class="card variant-glass-primary shadow-lg">
      <header class="card-header">
        <h2 class="h4 font-semibold">Club</h2>
      </header>

      <section class="p-4 space-y-2">
        <div class="flex items-center gap-3">
          <ShieldOutline />
          <div>
            <p>{club.name}</p>
            <p class="text-sm font-light">Name</p>
          </div>
        </div>

        <div class="flex items-center gap-3">
          <TagOutline />
          <div>
            <p>{club.acronym}</p>
            <p class="text-sm font-light">Acronym</p>
          </div>
        </div>

        <div class="flex items-center gap-3">
          <InfoCircleOutline />
          <div>
            <p>{club.bsm_id}</p>
            <p class="text-sm font-light">BSM-ID</p>
          </div>
        </div>
      </section>
    </div>
  {/each}

  <div class="card variant-glass-surface shadow-lg">
    <header class="card-header">
      <h2 class="h4 font-semibold">Attendance Stats</h2>
    </header>

    <section class="p-4">
      <ConicGradient legend stops={conicStops}></ConicGradient>
    </section>

    <footer class="mt-1 card-footer font-light text-sm">
      Events where you did not select anything are counted as "out".
    </footer>
  </div>

  <div class="card variant-glass-surface shadow-lg">
    <header class="card-header">
      <h2 class="h4 font-semibold">Totals</h2>
    </header>

    <section class="p-4 flex justify-center">
      <StatsBlockContent {block} classes="place-items-center gap-3" />
    </section>
  </div>

  <div class="card variant-glass-surface shadow-lg">
    <header class="card-header">
      <h2 class="h4 font-semibold">Totals</h2>
    </header>

    <section class="p-4 flex justify-center">
      <StatsBlockContent block={block2} classes="place-items-center gap-3" />
    </section>
  </div>
</div>

<h2 class="h2 mt-3">My Teams</h2>

{#await data.teams then teams}
  <div class="grid grid-cols-1 md:grid-cols-2 gap-3 mb-3">
    {#each teams as team}
      <a href="/account/team/{team.id}">
        <TeamTeaserCard {team} link={true} />
      </a>
    {/each}
  </div>
{/await}
