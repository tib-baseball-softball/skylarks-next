<script lang="ts">
  import {onMount} from "svelte";
  import flatpickr from "flatpickr";
  import "flatpickr/dist/flatpickr.css";
  import type {Options} from "flatpickr/dist/types/options";
  import {browser} from "$app/environment";

  if (browser && window.matchMedia("(prefers-color-scheme: dark)").matches) {
    import("flatpickr/dist/themes/dark.css");
  }

  interface Props {
    value?: string | Date;
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

<input bind:this={datepicker} class="input" type="text" {value}/>

<!-- svelte-ignore css_unused_selector -->
<style lang="postcss">
  input {
    color: unset;
  }
</style>
