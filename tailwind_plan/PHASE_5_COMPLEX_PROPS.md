# Phase 5: Complex Prop-based Migration

Handle components that receive Tailwind classes via props (e.g., `class`, `panelPadding`, `buttonClasses`).

## 5.1 Simple Prop Mapping

- [ ] Identify props that only take a few Tailwind variations.
- [ ] Replace with semantic props (e.g., `padding="small" | "large"`).
- [ ] Implement styling internally using a map or switch in the `<style>` block:
    ```css
    .root[data-padding="small"] { padding: calc(var(--spacing) * 1); }
    ```

## 5.2 CSS Variable Injection

- [ ] For highly configurable components, use CSS custom properties passed via the `style` attribute.
- [ ] Example: `style="--component-padding: {paddingValue}"`.

## 5.3 Modern Svelte 5 Patterns

- [ ] Use `snippets` for passing complex HTML/styles if applicable.
- [ ] Utilize `class` prop correctly in Svelte 5 components (binding or passing down).

## 5.4 Marking Complex Logic (Deferred)

- [ ] If a component's logic is too coupled with Tailwind strings, mark it with:
  `TODO: tailwind-migration [description]`
- [ ] Document why it's complex (e.g., "Uses dynamic string concatenation for 20+ utility variants").
