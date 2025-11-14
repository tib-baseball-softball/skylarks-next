<script lang="ts">
  import StreakEmoji from "$lib/components/favorite/StreakEmoji.svelte";
  import WinsOverSeasonChart from "$lib/components/favorite/WinsOverSeasonChart.svelte";
  import type {HomeDataset} from "$lib/tib/types/HomeDataset.ts";

  interface Props {
    dataset: HomeDataset;
  }

  const {dataset}: Props = $props();
</script>

<section>
  <WinsOverSeasonChart streakData={dataset?.streak_data ?? []}/>
</section>

<section class="mt-6!">
  <StreakEmoji streak={dataset.table_row.streak}/>
</section>