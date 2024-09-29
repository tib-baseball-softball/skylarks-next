<script lang="ts">
    import type {Typo3ContentObject} from "$lib/model/Typo3ContentObject";

    interface Props {
        co: Typo3ContentObject;
    }

    let { co }: Props = $props();
</script>

<section>
    <hgroup>
        <h2 class="h{co.content.headerLayout}">
            {co.content.header}
        </h2>

        <h3 class="h{co.content.headerLayout + 1}">
            {co.content.subheader}
        </h3>
    </hgroup>

    <p>
        {@html co.content.bodytext}
    </p>
</section>