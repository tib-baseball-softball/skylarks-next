<script lang="ts">
  interface props {
    strokeColor: string;
  }

  let {strokeColor}: props = $props();
</script>

<svg class="icon" viewBox="0 0 512 512" xmlns="http://www.w3.org/2000/svg"
>
  <path
    d="M294.25 108.6l-60.57 17.59M317.13 156.11l-51.5 36.4M355.89 194.87l-36.4 51.5M403.4 217.75l-17.59 60.57"
    fill="none"
    stroke={strokeColor}
    stroke-linecap="round"
    stroke-linejoin="round"
    stroke-width="32"
  />
  <circle
    cx="256"
    cy="256"
    fill="none"
    r="192"
    stroke={strokeColor}
    stroke-linecap="round"
    stroke-linejoin="round"
    stroke-width="32"
  />
  <path
    d="M432.94 255.05a192 192 0 01-176.31-180.7M108.54 294.31l17.59-60.57M156.05 317.19l36.4-51.5M194.81 355.95l51.5-36.4M217.69 403.46l60.57-17.59M255 433.61A192 192 0 0074.29 256.69"
    fill="none"
    stroke={strokeColor}
    stroke-linecap="round"
    stroke-linejoin="round"
    stroke-width="32"
  />
</svg>

<style>
  .icon {
    width: calc(var(--spacing) * 10);
    height: calc(var(--spacing) * 10);
  }
</style>
