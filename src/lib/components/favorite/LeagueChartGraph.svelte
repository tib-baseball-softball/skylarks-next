<script lang="ts">
  import type {Table} from "bsm.js";
  import ApexChart from "$lib/components/utility/charts/ApexChart.svelte";
  import {ColorUtility} from "$lib/service/ColorUtility.ts";

  interface Props {
    table: Table;
  }

  let {table}: Props = $props();

  let options: ApexCharts.ApexOptions = {
    series: [{
      name: 'Wins',
      data: table.rows.map((row) => {
        return {
          x: row.short_team_name,
          y: row.wins_count,
          fillColor: row.short_team_name.includes("BEA") ? ColorUtility.SkylarksRed : ColorUtility.SkylarksNavy,
        }
      })
    }],
    chart: {
      type: 'bar',
      height: 350,
      toolbar: {
        show: true
      }
    },
    plotOptions: {
      bar: {
        horizontal: false,
        columnWidth: '55%',
      }
    },
    dataLabels: {
      enabled: false
    },
    stroke: {
      show: true,
      width: 2,
      colors: ['transparent']
    },
    yaxis: {
      title: {
        text: 'Total Wins'
      }
    },
    fill: {
      opacity: 1
    },
    tooltip: {
      y: {
        formatter: function (val) {
          return `${val}`;
        }
      }
    },
    colors: [ColorUtility.SkylarksNavy]
  };
</script>

<ApexChart {options}>
  {#snippet beforeContent()}
    <h5 class="h5">League Wins Visualisation</h5>
  {/snippet}
</ApexChart>