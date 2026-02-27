<script lang="ts">
  interface Option<T extends string> {
    label?: string;
    value: T;
  }

  interface Props<T extends string> {
    label?: string;
    name: string;
    options: Option<T>[] | T[];
    value: T;
    classes?: string;
    listClass?: string;
    triggerClass?: string;
    required?: boolean;
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
  }: Props<string> = $props();

  function toOptions(arr: Option<string>[] | string[]): Option<string>[] {
    return (arr as any[]).map((o) =>
      typeof o === "string" ? {value: o, label: o.charAt(0).toUpperCase() + o.slice(1)} : o
    );
  }

  const opts: Option<string>[] = $derived(toOptions(options));
</script>

{#if label}
  <span class="block" data-required="{required}">{label}</span>
{/if}
<fieldset class={listClass + (classes)}>
  {#each opts as opt}
    <label class={[triggerClass, value === opt.value && "preset-filled"]}>
      <input
        type="radio"
        class="hidden-radio"
        name={name}
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

  .hidden-radio {
    display: none;
  }
</style>
