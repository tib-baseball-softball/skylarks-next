<script lang="ts">
  import type {MatchBoxscore} from "bsm.js";
  import Linescore from "$lib/tib/components/boxscore/Linescore.svelte";
  import MatchBoxscoreOffensiveTable from "$lib/tib/components/boxscore/MatchBoxscoreOffensiveTable.svelte";
  import MatchBoxscoreAdditionalStatsSection
    from "$lib/tib/components/boxscore/MatchBoxscoreAdditionalStatsSection.svelte";
  import MatchBoxscorePitchingTable from "$lib/tib/components/boxscore/MatchBoxscorePitchingTable.svelte";

  interface Props {
    boxscore: MatchBoxscore;
  }

  let {boxscore}: Props = $props();
  const awayTeamName = boxscore.linescore.away.league_entry.team?.name ?? "Away Team";
  const homeTeamName = boxscore.linescore.home.league_entry.team?.name ?? "Home Team";
</script>

<section class="mt-2 mb-6">
  <h2 class="h3 my-2">Linescore</h2>
  <Linescore linescore={boxscore.linescore}/>
</section>

<h2 class="h3 my-8">Offensive</h2>

<!--ROAD TEAM-->

<MatchBoxscoreOffensiveTable
        matchStats={boxscore.offensive_away}
        teamName={awayTeamName}
></MatchBoxscoreOffensiveTable>

<MatchBoxscoreAdditionalStatsSection
        stats={boxscore.additional_away}
></MatchBoxscoreAdditionalStatsSection>

<hr class="my-8"/>

<!--HOME TEAM-->

<MatchBoxscoreOffensiveTable
        matchStats={boxscore.offensive_home}
        teamName={homeTeamName}
></MatchBoxscoreOffensiveTable>

<MatchBoxscoreAdditionalStatsSection
        stats={boxscore.additional_home}
></MatchBoxscoreAdditionalStatsSection>

<hr class="my-8"/>

<section class="flex flex-col gap-4 mb-3">
  <h2 class="h3 my-2">Pitching</h2>

  <MatchBoxscorePitchingTable
          matchStats={boxscore.pitching_away}
          teamName={awayTeamName}
  ></MatchBoxscorePitchingTable>

  <div class="mt-4">
    <MatchBoxscorePitchingTable
            matchStats={boxscore.pitching_home}
            teamName={homeTeamName}
    ></MatchBoxscorePitchingTable>
  </div>
</section>
