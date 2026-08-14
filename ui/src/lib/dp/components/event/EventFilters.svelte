<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/state";
  import type { EventType } from "$lib/dp/types/ExpandedResponse";
  import type { TabSetOption } from "../formElements/TabsRadioGroup.svelte";
  import TabsRadioGroup from "../formElements/TabsRadioGroup.svelte";

  type SortingValues = "asc" | "desc" | string;
  type TypeValues = EventType | "any" | string;

  let showEvents = $state(page.url.searchParams.get("timeframe") ?? "next");

  let sorting: SortingValues = $state(
    page.url.searchParams.get("sort") ?? "asc",
  );

  let showTypes: TypeValues = $state(
    page.url.searchParams.get("type") ?? "any",
  );

  const queryString = $derived(
    `?timeframe=${showEvents}&sort=${sorting}&type=${showTypes}`,
  );

  const reloadWithQuery = () => {
    goto(queryString, {
      noScroll: true,
      keepFocus: true,
    });
  };

  const timeframeOptions: TabSetOption<string>[] = [
    {
      label: "Next",
      value: "next",
    },
    {
      label: "Past",
      value: "past",
    },
  ];

  const sortOptions: TabSetOption<string>[] = [
    {
      label: "Ascending",
      value: "asc",
    },
    {
      label: "Descending",
      value: "desc",
    },
  ];

  const typeOptions: TabSetOption<string>[] = [
    {
      label: "All",
      value: "any",
    },
    {
      label: "Game",
      value: "game",
    },
    {
      label: "Practice",
      value: "practice",
    },
    {
      label: "Other",
      value: "misc",
    },
  ];
</script>

<!--
@component reusable event filters - uses URL query string to save 
preferences and programmatically reloads the current page on state changes.

Usage on a page expects you to read the query parameters in your +page.ts.
See EventService.ts for an encapsulated solution that reads those parameters back.
-->

<div class="filters-bar preset-outlined-card">
  <label class="filter-label">
    <span>Timeframe</span>

    <TabsRadioGroup
      bind:value={showEvents}
      onValueChange={reloadWithQuery}
      options={timeframeOptions}
      label="Timeframe"
      hideLabel={true}
      name="timeframe"
      listClass="event-segment-container"
    ></TabsRadioGroup>
  </label>

  <label class="filter-label">
    <span>Sort</span>

    <TabsRadioGroup
      bind:value={sorting}
      onValueChange={reloadWithQuery}
      options={sortOptions}
      label="Sort"
      hideLabel={true}
      name="sort"
      listClass="event-segment-container"
    ></TabsRadioGroup>
  </label>

  <label class="filter-label">
    <span>Type</span>

    <TabsRadioGroup
      bind:value={showTypes}
      onValueChange={reloadWithQuery}
      options={typeOptions}
      label="Type"
      hideLabel={true}
      name="type"
      listClass="event-segment-container type-tabs"
    ></TabsRadioGroup>
  </label>
</div>

<style>
  .filters-bar {
    display: flex;
    flex-wrap: wrap;
    gap: calc(var(--spacing) * 4);
    justify-content: space-between;
    padding-inline: calc(var(--spacing) * 4);
    padding-block: calc(var(--spacing) * 3);
    border-radius: var(--radius-base);
    font-size: var(--text-sm);

    @media (min-width: 64rem) {
      font-size: var(--text-base);
    }
  }

  .filter-label {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 2);
    justify-content: space-between;
    flex-grow: 1;

    @media (min-width: 48rem) {
      flex-grow: 0;
    }

    @media (min-width: 80rem) {
      justify-content: flex-start;
    }
  }
</style>
