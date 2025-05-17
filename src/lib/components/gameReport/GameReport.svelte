<script lang="ts">
  import type {GameReport} from "$lib/model/GameReport.ts";
  import GameReportHeader from "$lib/components/gameReport/GameReportHeader.svelte";
  import TypoImage from "$lib/components/media/TypoImage.svelte";

  interface Props {
    report: GameReport;
    classes?: string;
  }

  let {report, classes = ""}: Props = $props();
</script>

<div class="flex justify-center">
  <!-- MARK: Language Tag (as long as we do not have English content) -->
  <div class="{classes}" lang="de">
    <h2 class="h2 mb-3">{report.title}</h2>
    <article>
      <GameReportHeader report={report}/>

      <section class="prose">
        {#if report.header_image}
          {#each report.header_image as image}
            <TypoImage media={image}/>
          {/each}
        {/if}

        <div class="mt-3">
          {@html report.introduction}
        </div>
      </section>

      <hr class="my-6!">

      <section class="mt-4 prose">
        <h3 class="h3">Bericht Spiel 1</h3>

        {@html report.report_first}
      </section>

      {#if report.report_second}
        <section class="my-3 prose">
          <h3 class="h3">Bericht Spiel 2</h3>

          {@html report.report_second}
        </section>
      {/if}

      {#if report.preview}
        <section class="card preset-tonal dark:border dark:border-surface-500 rounded-base shadow-xl p-4 my-4">
          <h3 class="h3">Ausblick</h3>

          {@html report.preview}
        </section>
      {/if}

      {#if report.gallery}
        <section class="my-5">
          <h3 class="h3 mb-3">Bildergalerie</h3>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-3">

            {#each report.gallery as image}
              <TypoImage media={image}/>
            {/each}

          </div>
        </section>
      {/if}
    </article>
  </div>
</div>

<style>
    .prose :global {
        max-width: unset;
        text-align: justify;

        p {
            margin: 1rem 0;
        }
    }
</style>
