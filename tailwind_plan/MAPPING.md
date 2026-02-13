# Tailwind to Native CSS Mapping

This document provides a reference for mapping Tailwind utility classes to native CSS using project-specific variables.

## Spacing (`--spacing: 0.25rem`)

| Tailwind      | Native CSS Equivalent                                    |
|:--------------|:---------------------------------------------------------|
| `p-{n}`       | `padding: calc(var(--spacing) * {n});`                   |
| `pt-{n}`      | `padding-top: calc(var(--spacing) * {n});`               |
| `pb-{n}`      | `padding-bottom: calc(var(--spacing) * {n});`            |
| `pl-{n}`      | `padding-left: calc(var(--spacing) * {n});`              |
| `pr-{n}`      | `padding-right: calc(var(--spacing) * {n});`             |
| `px-{n}`      | `padding-inline: calc(var(--spacing) * {n});`            |
| `py-{n}`      | `padding-block: calc(var(--spacing) * {n});`             |
| `m-{n}`       | `margin: calc(var(--spacing) * {n});`                    |
| `mt-{n}`      | `margin-top: calc(var(--spacing) * {n});`                |
| `mb-{n}`      | `margin-bottom: calc(var(--spacing) * {n});`             |
| `ml-{n}`      | `margin-left: calc(var(--spacing) * {n});`               |
| `mr-{n}`      | `margin-right: calc(var(--spacing) * {n});`              |
| `mx-{n}`      | `margin-inline: calc(var(--spacing) * {n});`             |
| `my-{n}`      | `margin-block: calc(var(--spacing) * {n});`              |
| `gap-{n}`     | `gap: calc(var(--spacing) * {n});`                       |
| `space-x-{n}` | `& > * + * { margin-left: calc(var(--spacing) * {n}); }` |
| `space-y-{n}` | `& > * + * { margin-top: calc(var(--spacing) * {n}); }`  |

## Colors

| Tailwind         | Native CSS Equivalent                     |
|:-----------------|:------------------------------------------|
| `bg-{color}`     | `background-color: var(--color-{color});` |
| `text-{color}`   | `color: var(--color-{color});`            |
| `border-{color}` | `border-color: var(--color-{color});`     |

*Note: For `tonal` presets, use the specific semantic variables found in `shame.css` or `dp-theme.css`.*

## Typography

| Tailwind      | Native CSS Equivalent                                         |
|:--------------|:--------------------------------------------------------------|
| `text-xs`     | `font-size: var(--text-xs);`                                  |
| `text-sm`     | `font-size: var(--text-sm);`                                  |
| `text-base`   | `font-size: var(--text-base);`                                |
| `text-lg`     | `font-size: var(--text-lg);`                                  |
| `text-xl`     | `font-size: var(--text-xl);`                                  |
| `text-2xl`    | `font-size: var(--text-2xl);`                                 |
| `font-bold`   | `font-weight: bold;` (or check for `var(--font-weight-bold)`) |
| `font-medium` | `font-weight: 500;`                                           |

## Border Radius

| Tailwind            | Native CSS Equivalent                     |
|:--------------------|:------------------------------------------|
| `rounded`           | `border-radius: var(--radius-base);`      |
| `rounded-sm`        | `border-radius: var(--radius-sm);`        |
| `rounded-lg`        | `border-radius: var(--radius-lg);`        |
| `rounded-xl`        | `border-radius: var(--radius-xl);`        |
| `rounded-container` | `border-radius: var(--radius-container);` |

## Layout

| Tailwind         | Native CSS Equivalent      |
|:-----------------|:---------------------------|
| `flex`           | `display: flex;`           |
| `grid`           | `display: grid;`           |
| `hidden`         | `display: none;`           |
| `items-center`   | `align-items: center;`     |
| `justify-center` | `justify-content: center;` |
| `flex-col`       | `flex-direction: column;`  |

## KEEP (Designated Utilities & Presets)

The following should NOT be migrated to scoped CSS and must be preserved:

### From `tw_utils.css`

- `sr-only`, `transform`, `ring`, `blur`, `filter`, `transition`
- `shadow`, `shadow-md`, `shadow-lg`, `shadow-xl`, `shadow-2xl`

### From `presets.css`

- **ALL** `preset-*` classes (e.g., `preset-tonal-*`, `preset-filled-*`, `preset-outlined-*`).

### From Other Project CSS

- All classes in `ui/src/css/dp/partials/*.css` (e.g., `.btn`, `.card`, `.badge`, `.stats`).
- All classes in `ui/src/css/dp/theme/*.css` (except `tailwind_merged.css` and `shame.css`).
