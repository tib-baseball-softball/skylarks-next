<script lang="ts">
  import type {GameReport} from "$lib/model/GameReport.ts";
  import GameReportHeader from "$lib/components/gameReport/GameReportHeader.svelte";

  interface Props {
    report: GameReport;
    classes?: string;
  }

  let {report, classes = ""}: Props = $props();
</script>

<!-- MARK: Language Tag -->
<div class="flex justify-center">
  <div class="{classes}" lang="de">
    <h2 class="h2 mb-3">{report.title}</h2>
    <article>
      <GameReportHeader report={report}/>

      <section class="prose">
        {#if report.header_image}
          <img
                  class=""
                  loading="lazy"
                  alt={report.header_image.alt}
                  src={report.header_image.url}
          />
        {/if}
        <div class="mt-3">
          {@html report.introduction}
        </div>
      </section>

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
        <section class="card variant-soft dark:border dark:border-surface-500 rounded-token shadow-xl p-4 my-4">
          <h3 class="h3">Ausblick</h3>

          {@html report.preview}
        </section>
      {/if}

      {#if report.gallery}
        <section id="imageCarousel" class="">
          <h3 class="h3">Bildergalerie</h3>

        </section>
      {/if}
    </article>
  </div>
</div>

<style>
    .prose {
        max-width: unset;
    }
</style>
