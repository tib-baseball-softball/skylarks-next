<script lang="ts">
  import ApexChart from "$lib/components/utility/charts/ApexChart.svelte";
  import {ColorUtility} from "$lib/tib/service/ColorUtility.ts";
  import type {StreakDataEntry} from "$lib/tib/types/HomeDataset.ts";

  interface Props {
    streakData: StreakDataEntry[];
  }

  const {streakData}: Props = $props();

  const options: ApexCharts.ApexOptions = $derived({
    chart: {
      height: 350,
      type: "line",
      stacked: false,
      zoom: {
        enabled: false,
      },
    },
    markers: {
      size: 4,
    },
    dataLabels: {
      enabled: false,
    },
    colors: [ColorUtility.SkylarksRed],
    series: [
      {
        name: "Wins Count",
        data: streakData.map((dataPoint) => dataPoint.wins_count),
      },
    ],
    stroke: {
      width: [3, 3],
    },
    plotOptions: {
      bar: {
        columnWidth: "20%",
      },
    },
    xaxis: {
      categories: streakData.map((dataPoint) => Number(dataPoint.game)),
      title: {
        text: "Gamedays",
        style: {
          color: ColorUtility.getDynamicColorNavySand(),
        },
      },
    },
    yaxis: [
      {
        axisTicks: {
          show: true,
        },
        axisBorder: {
          show: true,
          color: "white",
        },
        labels: {
          style: {
            colors: "white",
          },
        },
        title: {
          text: "Wins Count",
          style: {
            color: ColorUtility.getDynamicColorNavySand(),
          },
        },
      },
    ],
    tooltip: {
      shared: false,
      intersect: true,
      x: {
        show: false,
      },
    },
    legend: {
      horizontalAlign: "left",
      offsetX: 40,
    },
  });
</script>

<ApexChart {options}>
  {#snippet beforeContent()}
    <h5 class="h5">Season Progression</h5>
  {/snippet}
</ApexChart>
