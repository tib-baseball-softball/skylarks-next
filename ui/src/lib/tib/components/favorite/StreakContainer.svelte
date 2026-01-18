<script lang="ts">
  import StreakEmoji from "$lib/tib/components/favorite/StreakEmoji.svelte";
  import WinsOverSeasonChart from "$lib/tib/components/favorite/WinsOverSeasonChart.svelte";
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