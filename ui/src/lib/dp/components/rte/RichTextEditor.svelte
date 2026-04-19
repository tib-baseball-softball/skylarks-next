<script lang="ts">
  import {onMount} from "svelte";

  interface Props {
    content: string;
  }

  let {content = $bindable()}: Props = $props();

  let editorElement: HTMLElement;

  onMount(() => {

  });

</script>

<div bind:this={editorElement}></div>
