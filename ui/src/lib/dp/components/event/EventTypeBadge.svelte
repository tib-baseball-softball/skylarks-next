<script lang="ts">
  interface props {
    type: string;
  }

  const {type}: props = $props();

  function getDisplayString(): string {
    switch (type) {
      case "game":
        return "Game";
      case "practice":
        return "Practice";
      default:
        return "Event";
    }
  }
</script>

<span class="badge" data-type={type}>
  {getDisplayString()}
</span>

<style>
  .badge {
    &[data-type="game"] {
      background-color: var(--color-primary-500);
      color: white;
    }

    &[data-type="practice"] {
      background-color: var(--color-secondary-500);
      color: white;
    }

    &[data-type="misc"] {
      background-color: var(--color-tertiary-500);
      color: var(--color-secondary-500)
    }
  }
</style>
