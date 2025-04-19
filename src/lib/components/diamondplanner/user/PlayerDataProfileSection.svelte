<script lang="ts">
  import PlayerDataCard from "$lib/components/player/PlayerDataCard.svelte";
  import PlayerHeaderSection from "$lib/components/player/PlayerHeaderSection.svelte";
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
  import type {Player} from "$lib/model/Player";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import {Link} from "lucide-svelte";
  import {convertIntKeyToHandedness} from "$lib/types/Handedness.js";
  import {positionKeysToEnumStringValues} from "$lib/types/BaseballPosition.js";
  import PlayerDataForm from "$lib/components/forms/PlayerDataForm.svelte";

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  // suboptimal adapter until data is properly fetched from PocketBase
  //@ts-expect-error
  let playerObject: Player = $derived({
    firstname: authRecord.first_name,
    lastname: authRecord.last_name,
    fullname: authRecord.first_name + " " + authRecord.last_name,
    birthday: 0,
    admission: new Date(authRecord.created).toLocaleDateString(),
    number: authRecord.number,
    throwing: convertIntKeyToHandedness(Number(authRecord.throws)),
    batting: convertIntKeyToHandedness(Number(authRecord.bats)),
    bsm_id: authRecord.bsm_id,
    positions: positionKeysToEnumStringValues(authRecord.position),
    teams: [],
    coach: "",
    media: [
      {
        url: client.files.getURL(authRecord, authRecord.avatar),
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
      You have the option to display your player data in your public
      profile on the Skylarks website and apps.
    </p>
    <p class="mt-2">
      Please note that all statistics and basic info such as your name and
      team affiliation is considered public information when participating
      in official DBV and BSVBB competitions.
    </p>
  </section>

  <footer class="card-footer">
    <div class="flex flex-col gap-2 lg:gap-3">
      <PlayerDataForm buttonClasses="btn variant-ghost-primary"></PlayerDataForm>

      {#if authRecord.bsm_id}
        <a
                href="/players/{authRecord?.bsm_id}"
                class="btn variant-ghost-secondary"
        >
          <Link/>
          <span>Go to Profile Page</span>
        </a>
      {/if}
    </div>
  </footer>
</div>

<PlayerHeaderSection player={playerObject}/>

<PlayerDataCard player={playerObject} showTeams={false}/>
