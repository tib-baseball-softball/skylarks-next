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
    onValueChange?: () => void;
  }

  let {
    label,
    name,
    options,
    value = $bindable(),
    classes = "",
    listClass = "",
    triggerClass = "tabs-trigger btn",
    required = false,
    hideLabel = false,
    onValueChange,
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

<fieldset class={["fieldset", listClass, classes]}>
  {#if label}
    <legend class={[hideLabel && "sr-only", "legend"]} data-required={required}>
      {label}
    </legend>
  {/if}
  <div class="input tabs-list">
    {#each opts as opt}
      <label class={[triggerClass, value === opt.value && "preset-filled"]}>
        <input
          onchange={onValueChange}
          type="radio"
          class="sr-only"
          {name}
          bind:group={value}
          value={opt.value}
        />
        {opt.label}
      </label>
    {/each}
  </div>
</fieldset>

<style>
  /** Caution - this should be a mixin, see forms.css */
  [data-required="true"]:after {
    content: "*";
    color: light-dark(var(--color-primary-500), var(--color-primary-300));
    margin-inline-start: calc(var(--spacing) * 0.5);
  }

  .tabs-list {
    margin-block-start: var(--spacing);
  }
</style>
