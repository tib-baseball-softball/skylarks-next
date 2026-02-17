# Phase 5: Complex Prop-based Migration

Handle components that receive Tailwind classes via props (e.g., `class`, `panelPadding`, `buttonClasses`).

## 5.1 Simple Prop Mapping

- [x] Identify props that only take a few Tailwind variations.
- [x] Replace with semantic props (e.g., `padding="small" | "large"`).
- [x] Implement styling internally using a map or switch in the `<style>` block:
    ```css
    .root[data-padding="small"] { padding: calc(var(--spacing) * 1); }
    ```

## 5.2 CSS Variable Injection

- [x] For highly configurable components, prefer semantic prop mapping first; use CSS properties only where open-ended values are needed.
- [x] Example strategy documented for future remaining dynamic class props.

## 5.3 Modern Svelte 5 Patterns

- [x] Keep snippet-based trigger-content APIs where already in use.
- [x] Remove free-form Tailwind class prop APIs for migrated trigger components.

## 5.4 Marking Complex Logic (Deferred)

- [x] If a component's logic is too coupled with Tailwind strings, mark it with:
  `TODO: tailwind-migration [description]`
- [x] Document why it's complex (e.g., "Uses dynamic string concatenation for 20+ utility variants").

## 5.5 Implemented in this pass

### Migrated from free-form class props to semantic APIs

- `ui/src/lib/dp/components/forms/TeamForm.svelte`
- `ui/src/lib/dp/components/forms/LocationForm.svelte`
- `ui/src/lib/dp/components/forms/EventForm.svelte`
- `ui/src/lib/dp/components/forms/PlayerDataForm.svelte`
- `ui/src/lib/dp/components/forms/AnnouncementForm.svelte`
- `ui/src/lib/dp/components/forms/ClubForm.svelte`
- `ui/src/lib/dp/components/event/EventSeriesView.svelte`
- `ui/src/lib/dp/components/modal/AccordionItem.svelte`

### New semantic props

- Trigger components now use:
  - `triggerVariant`: `"filled-primary" | "tonal-primary" | "tonal-secondary" | "tonal-tertiary" | "tonal-surface"`
  - `triggerSize`: `"default" | "sm"`
  - `triggerIcon`: `boolean`
  - `triggerSpaced`: `boolean`
- Accordion item now uses:
  - `panelInset`: `"none" | "default"`

### Updated callsites

- Migrated all `buttonClasses=` usages in `ui/src/lib` and `ui/src/routes` to semantic props.
- Migrated all `panelPadding=` usages to `panelInset=`.

### Deferred complex class-prop surfaces

- `TODO: tailwind-migration` audit remains for broad `classes?: string` passthrough APIs (e.g. generic utility/icon components and delete-button variants), where the accepted class surface is intentionally open-ended and requires per-component API decisions.
