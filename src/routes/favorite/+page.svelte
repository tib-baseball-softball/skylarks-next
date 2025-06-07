<script lang="ts">
  import type {PageProps} from "./$types";
  import FavoriteTeamInfoCard from "$lib/components/favorite/FavoriteTeamInfoCard.svelte";
  import {preferences} from "$lib/globals.svelte.ts";
  import ClubTeamPicker from "$lib/components/utility/ClubTeamPicker.svelte";

  let {data}: PageProps = $props();
  const clubTeams = $derived(data.clubTeams ?? []);

  const favoriteTeam = $derived(clubTeams.filter(clubTeam => clubTeam.id === preferences.current.favoriteTeamID).at(0));
</script>

<h1 class="h1">Favorite Team</h1>

<ClubTeamPicker clubTeams={clubTeams}/>

<section>
  {#if preferences.current.favoriteTeamID !== 0 && favoriteTeam}
    <FavoriteTeamInfoCard clubTeam={favoriteTeam}/>
  {:else }
    möp
  {/if}
</section>
