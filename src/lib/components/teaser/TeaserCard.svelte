<script lang="ts">
    export let header: String
    export let teaserText: String
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