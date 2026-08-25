<script lang="ts">
  import { onMount } from "svelte";
  import { MaerkchenEditor } from "maerkchen";

  interface Props {
    value: string;
    label?: string;
    formElementName?: string;
    required?: boolean;
  }

  let editor: MaerkchenEditor;

  let {
    value = $bindable(),
    label,
    formElementName,
    required = false,
  }: Props = $props();

  onMount(() => {
    editor.addEventListener("change", () => {
      value = editor.markdownText;
    })
  });
</script>

<maerkchen-editor
  bind:this={editor}
  {label}
  {formElementName}
  {required}
  markdownText={value}
></maerkchen-editor>

<style>
  maerkchen-editor {
    --color-border: light-dark(var(--color-surface-50), #334155);
  }
</style>
