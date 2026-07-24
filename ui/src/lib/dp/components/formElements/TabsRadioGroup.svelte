<script lang="ts">
  export interface TabSetOption<T extends string> {
    label?: string;
    value: T;
  }

  interface Props<T extends string> {
    label?: string;
    name: string;
    options: TabSetOption<T>[] | T[];
    value: T;
    classes?: string;
    listClass?: string;
    triggerClass?: string;
    required?: boolean;
    hideLabel?: boolean;
  }

  let {
    label,
    name,
    options,
    value = $bindable(),
    classes = "",
    listClass = "tabs-list input",
    triggerClass = "tabs-trigger btn",
    required = false,
    hideLabel = false,
  }: Props<string> = $props();

  function toOptions(
    arr: TabSetOption<string>[] | string[],
  ): TabSetOption<string>[] {
    return (arr as any[]).map((o) =>
      typeof o === "string"
        ? { value: o, label: o.charAt(0).toUpperCase() + o.slice(1) }
        : o,
    );
  }

  const opts: TabSetOption<string>[] = $derived(toOptions(options));
</script>

<fieldset class={listClass + classes}>
  {#if label}
    <legend class={["block", hideLabel && "sr-only"]} data-required={required}>
      {label}
    </legend>
  {/if}
  {#each opts as opt}
    <label class={[triggerClass, value === opt.value && "preset-filled"]}>
      <input
        type="radio"
        class="sr-only"
        {name}
        bind:group={value}
        value={opt.value}
      />
      {opt.label}
    </label>
  {/each}
</fieldset>

<style>
  /** Caution - this should be a mixin, see forms.css */
  [data-required="true"]:after {
    content: "*";
    color: light-dark(var(--color-primary-500), var(--color-primary-300));
    margin-inline-start: calc(var(--spacing) * 0.5);
  }
</style>
