<script lang="ts">
  import {onDestroy, onMount, type Snippet} from 'svelte';
  import {browser} from '$app/environment';

  interface Props {
    options: ApexCharts.ApexOptions;
    beforeContent?: Snippet;
  }

  let {options, beforeContent}: Props = $props();

  let chartElement: HTMLElement;
  let chart: ApexCharts;

  onMount(async () => {
    if (browser) {
      const ApexCharts = (await import('apexcharts')).default;
      chart = new ApexCharts(chartElement, options);
      await chart.render();
    }
  });

  onDestroy(() => {
    if (chart) {
      chart.destroy();
    }
  });
</script>

<div class="chart-wrapper">
  {@render (beforeContent?.())}
  <div bind:this={chartElement} class="chart-container"></div>
</div>