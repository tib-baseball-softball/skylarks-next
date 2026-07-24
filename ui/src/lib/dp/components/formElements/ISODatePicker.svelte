<script lang="ts">
  import { DateTimeUtility } from "$lib/dp/service/DateTimeUtility";
  import { onMount } from "svelte";
  import type { HTMLInputAttributes } from "svelte/elements";

  interface Props {
    value: string;
    type?: "datetime-local" | "date";
    inputProps?: HTMLInputAttributes;
    required?: boolean;
  }

  let {
    value = $bindable(),
    type = "datetime-local",
    required = false,
    ...inputProps
  }: Props = $props();

  let datepicker: HTMLInputElement;

  onMount(() => {
    switch (type) {
      case "datetime-local":
        datepicker.value = DateTimeUtility.formatForDatetimeLocal(value);
        break;

      case "date":
        datepicker.value = DateTimeUtility.formatForDate(value);
        break;
    }

    datepicker.addEventListener("change", () => {
      value = new Date(datepicker.value).toISOString();
    });
  });
</script>

<!--
@component
Thin wrapper over native datetime elements to ensure correct date format.

Server sends RFC3399: "2026-07-24 10:00:00.000Z"
See [documentation](https://pocketbase.io/docs/collections/#datefield)
-->

<input {...inputProps} bind:this={datepicker} {type} class="input" {required} />
