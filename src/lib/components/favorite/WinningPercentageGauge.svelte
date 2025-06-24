<script lang="ts">
  import type {TableRow} from "bsm.js";
  import ApexChart from "$lib/components/utility/charts/ApexChart.svelte";
  import {ColorUtility} from "$lib/service/ColorUtility.ts";

  interface Props {
    tableRow: TableRow;
  }

  let {tableRow}: Props = $props();
  const percentage = $derived(Math.round(parseFloat(tableRow.quota) * 100));

  const options: ApexCharts.ApexOptions = $derived({
    chart: {
      height: 280,
      type: "radialBar"
    },

    series: [percentage],
    colors: [ColorUtility.SkylarksSand],
    plotOptions: {
      radialBar: {
        hollow: {
          margin: 15,
          size: "70%"
        },

        dataLabels: {
          name: {
            show: true,
            fontSize: "16px"
          },
          value: {
            fontSize: "16px",
            show: true
          }
        }
      }
    },

    stroke: {
      lineCap: "round",
    },
    labels: ["Pct."]
  });
</script>

<ApexChart {options}/>