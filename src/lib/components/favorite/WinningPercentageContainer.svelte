<script lang="ts">
import type { TableRow } from "bsm.js"
import { Crown, Hash } from "lucide-svelte"
import WinningPercentageGauge from "$lib/components/favorite/WinningPercentageGauge.svelte"

interface Props {
  tableRow: TableRow
}

let { tableRow }: Props = $props()
</script>

<section class="@container">
  <div class="grid grid-cols-1 @lg:grid-cols-3 gap-4 justify-items-center">
    <dl>
      <dt>Rank</dt>
      <dd class="flex gap-3 text-2xl font-extrabold justify-center text-details-element">
        {#if tableRow.rank === "1."}
          <Crown class="stroke-primary-500"/>
        {:else }
          <Hash/>
        {/if}

        {tableRow.rank}
      </dd>
    </dl>

    <dl>
      <dt>Wins/Losses</dt>
      <dd class=" text-2xl font-extrabold text-details-element">{tableRow.wins_count} - {tableRow.losses_count}</dd>
    </dl>

    <dl aria-hidden="true">
      <dt>Percentage</dt>
      <dd>
        <WinningPercentageGauge {tableRow}/>
      </dd>
    </dl>

  </div>
</section>

<style>
    dt {
        font-weight: var(--font-weight-light);
    }

    dl {
        display: flex;
        justify-content: space-between;
        align-items: center;
        width: 60%;
    }

    @container (width >= 32rem) {
        dl {
            flex-direction: column;
            gap: 0.5rem;
        }
        .text-details-element {
            padding: 3rem 0;
        }
    }
</style>