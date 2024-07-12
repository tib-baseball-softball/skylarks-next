<script lang="ts">
    export let content: string
    export let classes = "variant-filled-primary"
</script>

<span class="text-4xl font-bold {classes}">
    {content}
</span>

<style lang="postcss">
    span {
        display: block;
        height: 4rem;
        width: 4rem;
        line-height: 4rem;
        border-radius: 50%;
        text-align: center;
    }
</style>