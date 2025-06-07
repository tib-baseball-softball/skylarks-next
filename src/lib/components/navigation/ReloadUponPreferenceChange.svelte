<script lang="ts">
  import {preferences} from "$lib/globals.svelte.ts";

  interface Props {
    callback: () => void;
  }

  let {callback}: Props = $props();

  $effect(() => {
    // we want to track the preferences and react when they change, so they must be mentioned here,
    // but we don't need to do anything with them
    preferences;
    console.log(`global app preferences changed - reload`);
    callback();
  });
</script>