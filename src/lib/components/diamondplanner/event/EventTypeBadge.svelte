<script lang="ts">
    interface props {
        type: string
    }

    const {type}: props = $props()

    function getDisplayString(): string {
        switch (type) {
            case "game":
                return "Spiel"
            case "practice":
                return "Training"
            default:
                return "Event"
        }
    }
</script>

<span
    class="badge"
    class:variant-filled-primary={type === "game"}
    class:variant-filled-secondary={type === "practice"}
    class:dark:border={type === "practice"}
    class:variant-filled-tertiary={type === "misc"}
>
{getDisplayString()}
</span>