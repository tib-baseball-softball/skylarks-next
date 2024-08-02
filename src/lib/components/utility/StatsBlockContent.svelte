<script lang="ts">
    import type {SingleStatElement} from "$lib/types/SingleStatElement";

    interface props {
        block: SingleStatElement
    }

    let {block}: props = $props()
</script>

<div class="stat">
    <div class="stat-title">{block.title}</div>
    <div class="stat-value">{block.value}</div>
    <div class="stat-desc">{block.desc}</div>
</div>