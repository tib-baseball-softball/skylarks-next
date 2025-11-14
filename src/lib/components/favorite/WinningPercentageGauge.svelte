<script lang="ts">
  import type {TableRow} from "bsm.js";
  import ApexChart from "$lib/components/utility/charts/ApexChart.svelte";
  import {ColorUtility} from "$lib/tib/service/ColorUtility.ts";

  interface Props {
    tableRow: TableRow;
  }

  const {tableRow}: Props = $props();
  const percentage = $derived(Math.round(parseFloat(tableRow.quota) * 100));

  const options: ApexCharts.ApexOptions = $derived({
    chart: {
      height: 280,
      type: "radialBar",
    },

    series: [percentage],
    colors: [ColorUtility.SkylarksSand],
    plotOptions: {
      radialBar: {
        hollow: {
          margin: 15,
          size: "70%",
        },

        dataLabels: {
          name: {
            show: true,
            fontSize: "16px",
          },
          value: {
            fontSize: "16px",
            show: true,
          },
        },
      },
    },

    stroke: {
      lineCap: "round",
    },
    labels: ["Pct."],
  });
</script>

<ApexChart {options}/>