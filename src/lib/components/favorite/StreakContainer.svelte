<script lang="ts">
  import type {HomeDataset} from "$lib/types/HomeDataset.ts";
  import StreakEmoji from "$lib/components/favorite/StreakEmoji.svelte";
  import WinsOverSeasonChart from "$lib/components/favorite/WinsOverSeasonChart.svelte";

  interface Props {
    dataset: HomeDataset
  }
  let {dataset}: Props = $props();
</script>

<section>
  <WinsOverSeasonChart streakData={dataset.streak_data}/>
</section>

<section>
  <StreakEmoji streak={dataset.table_row.streak}/>
</section>