<script lang="ts">
  import type {PageProps} from "./$types";
  import FavoriteTeamInfoCard from "$lib/components/favorite/FavoriteTeamInfoCard.svelte";
  import {preferences} from "$lib/globals.svelte.ts";
  import ClubTeamPicker from "$lib/components/utility/ClubTeamPicker.svelte";

  let {data}: PageProps = $props();
  const clubTeams = $derived(data.clubTeams ?? []);

  const favoriteTeam = $derived(clubTeams.find(clubTeam => clubTeam.id === preferences.current.favoriteTeamID));
</script>

<h1 class="h1">Favorite Team</h1>

<ClubTeamPicker clubTeams={clubTeams}/>

<section>
  <h2 class="h2 my-3">Quick Details</h2>
  {#if preferences.current.favoriteTeamID !== 0 && favoriteTeam}
    <FavoriteTeamInfoCard clubTeam={favoriteTeam}/>
  {:else }
    <p>You have not selected a favorite team yet.</p>
  {/if}
</section>

<section>
  <h2 class="h2 mt-5 mb-1">Leagues</h2>
  <p class="text-sm font-light">A single team can be associated to several leagues per season in BSM.</p>
</section>