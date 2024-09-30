<script lang="ts">
  import PlayerDataCard from "$lib/components/player/PlayerDataCard.svelte";
  import PlayerHeaderSection from "$lib/components/player/PlayerHeaderSection.svelte";
  import type { CustomAuthModel } from "$lib/model/ExpandedResponse";
  import type { Player } from "$lib/model/Player";
  import { client } from "$lib/pocketbase";
  import { authModel } from "$lib/pocketbase/Auth";
  import { EditOutline, LinkOutline } from "flowbite-svelte-icons";

  const model = $authModel as CustomAuthModel;

  // suboptimal adapter until data is properly fetched from PocketBase
  let playerObject: Player = $derived({
    firstname: model.first_name,
    lastname: model.last_name,
    fullname: model.first_name + " " + model.last_name,
    birthday: 0,
    admission: model.created,
    number: model.number,
    throwing: model.throws,
    batting: model.bats,
    bsm_id: model.bsm_id,
    positions: model.position,
    teams: [],
    coach: "",
    media: [
      {
        url: client.files.getUrl(model, model.avatar),
        alt: "Player Profile Avatar",
      },
    ],
  });
</script>

<div class="card variant-glass-primary shadow-lg flex flex-col justify-between">
  <header class="card-header">
    <h2 class="h4 font-semibold">Public Profile</h2>
  </header>

  <section class="p-4 font-light">
    <p>
      You have the option to display your player data in your public profile on
      the Skylarks website and apps.
    </p>
    <p class="mt-2">
      Please note that all statistics and basic info such as your name and team affiliation is considered
      public information when participating in official DBV and BSVBB
      competitions.
    </p>
  </section>

  <footer class="card-footer">
    <div class="flex flex-col gap-2 lg:gap-3">
      <button class="btn variant-ghost-primary">
        <EditOutline />
        <span>Edit Player Data</span>
      </button>

      {#if model.bsm_id}
        <a href="/players/{model?.bsm_id}" class="btn variant-ghost-secondary">
          <LinkOutline />
          <span>Go to Profile Page</span>
        </a>
      {/if}
    </div>
  </footer>
</div>

<PlayerHeaderSection player={playerObject} />

<PlayerDataCard player={playerObject} showTeams={false} />
