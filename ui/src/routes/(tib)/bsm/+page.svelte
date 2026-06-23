<script lang="ts">

</script>

<svelte:head>
  <title>BSM Data</title>
  <meta
    content="BSM data page, displaying the BSM data for the club."
    name="description"
  />
</svelte:head>


<iframe title="Skylarks BSM Data" src="https://bsm.berlinskylarks.de/"></iframe>

<style>
  iframe {
    width: 100%;
    height: 100vh;
  }

  :global {
    .root-layout[data-page="/(tib)/bsm"] {
      .main-wrapper {
        width: 98%;
      }
    }
  }
</style>
