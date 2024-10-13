<script lang="ts">
  interface Props {
    identifier: string; // Skylarks Caps get assigned a unique identifier, for all others this is just a color
    classes: string
  }

  let {identifier, classes}: Props = $props();
</script>

{#if identifier === "B"}
    <img class="{classes}" src="/uniform/B_cap.svg" alt="Berlin Skylarks Home Cap">
{:else if identifier === "Bird"}
    <img class="{classes}" src="/uniform/Bird_cap.svg" alt="Berlin Skylarks Road Cap">
{:else}
    <div>MÖP</div>
{/if}