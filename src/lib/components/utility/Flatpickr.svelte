<script lang="ts">
  import {onMount} from "svelte";
  import flatpickr from "flatpickr";
  import "flatpickr/dist/flatpickr.css";
  import type {Options} from "flatpickr/dist/types/options";

  interface Props {
    value: string; // flatpickr accepts timestamps, date objects and ISO strings, but we work with strings exclusively
    options: Options;
  }

  let {value = $bindable(), options = {}}: Props = $props();

  let datepicker: HTMLInputElement;

  onMount(() => {
    const fp = flatpickr(datepicker, {
      ...options,
      defaultDate: value ?? "",
      onChange: (selectedDates) => {
        value = selectedDates[0].toISOString();
      },
    });

    return () => {
      fp.destroy();
    };
  });
</script>

<input class="dark:bg-surface-700" type="text" bind:this={datepicker} {value}/>

<style lang="postcss">
    input {
        @apply input;
        color: unset;
    }
</style>
