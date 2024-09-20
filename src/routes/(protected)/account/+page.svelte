<script lang="ts">
  import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";
  import { UserOutline, EnvelopeOutline } from "flowbite-svelte-icons";
  import { authModel } from "$lib/pocketbase/Auth";
  import type { CustomAuthModel } from "$lib/model/ExpandedResponse.js";
  import { Avatar } from "@skeletonlabs/skeleton";
  import { client } from "$lib/pocketbase";

  const model = $authModel as CustomAuthModel;

  let { data } = $props();
  console.log(model);
</script>

<h1 class="h1">My Dashboard</h1>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3 mb-3">
  <div class="card variant-glass-surface lg:col-span-2 shadow-lg">
    <header class="card-header">
      <h2 class="h4 font-semibold">User Data</h2>
    </header>

    <section class="p-4 grid grid-cols-5 gap-2 lg:gap-3 items-center">
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
      <button class="btn variant-filled-primary">Edit</button>
    </footer>
  </div>

  <div class="card variant-glass-primary shadow-lg">
    <header class="card-header">(header)</header>
    <section class="p-4">(content)</section>
    <footer class="card-footer">(footer)</footer>
  </div>

  <div class="card">
    <header class="card-header">(header)</header>
    <section class="p-4">(content)</section>
    <footer class="card-footer">(footer)</footer>
  </div>
  <div class="card">
    <header class="card-header">(header)</header>
    <section class="p-4">(content)</section>
    <footer class="card-footer">(footer)</footer>
  </div>
  <div class="card">
    <header class="card-header">(header)</header>
    <section class="p-4">(content)</section>
    <footer class="card-footer">(footer)</footer>
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
