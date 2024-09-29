<script lang="ts">
    interface Props {
        header: String;
        teaserText: String;
    }

    let { header, teaserText }: Props = $props();
</script>

<div class="card">
    <header class="card-header">
       <h3> {header} </h3>
    </header>
    <section class="p-4">{teaserText}</section>
    <footer class="card-footer">
        LINK
    </footer>
</div>