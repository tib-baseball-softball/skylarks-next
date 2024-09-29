<script lang="ts">
    import TeaserCard from "$lib/components/teaser/TeaserCard.svelte";
    import CurrentMatchesOverviewCard from "$lib/components/match/CurrentMatchesOverviewCard.svelte";

    interface Props {
        data: any;
    }

    let { data }: Props = $props();

    let page = $derived(data.pageObject)
//    $: contentObjects = data.contentObjects ?? []
</script>

<div class="w-full space-y-5 mt-2 mb-5">
    <section id="Hero" class="hero w-screen h-[20em] rounded-lg">
            <div class="hero-overlay bg-opacity-60"></div>
            <div class="hero-content">
                <div class="max-w-md text-white variant-glass-surface p-3 rounded-xl">
                    <h1 class="h1 mb-2">Berlin Skylarks</h1>
                    <p class="">{page?.meta?.subtitle ?? "Baseball & Softball im Herzen Berlins"}</p>
                </div>
            </div>
    </section>

    <section id="Scores">
        <CurrentMatchesOverviewCard
                matchesCurrent="{data.streamed.matchesCurrent}"
                matchesPrevious="{data.streamed.matchesPrevious}"
                matchesNext="{data.streamed.matchesNext}"
        />
    </section>

<!--    <section id="News">
        <h2 class="h2 mb-3">Aktuelle News</h2>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 lg:gap-6">
            {#each contentObjects as co}
                {#if co.type === "pagelist_sub"}
                    {#each co.content.pages as page}
                        <TeaserCard header="{page.title}" teaserText="Lorem ipsum dolor sit amet"></TeaserCard>
                    {/each}
                {/if}
            {/each}
        </div>
    </section>-->

<!--    <section class="">
        <a class="btn variant-filled" href="/skylarks-zwischenzeugnis-2022">Testseite mit Inhalten</a>
        <a class="btn variant-filled" href="/skylarks-holen-den-titel">Spielbericht</a>
        <p>Inhalt</p>
    </section>-->
</div>

<style>
    .hero {
        display: grid;
        width: 100%;
        place-items: start;
        background-size: cover;
        background-position: center;
        background-image: url("/header_hero.jpg");
    }

    .hero-content {
        z-index: 0;
        display: flex;
        align-items: center;
        justify-content: center;
        max-width: 80rem;
        gap: 1rem;
        padding: 1rem;
    }
</style>