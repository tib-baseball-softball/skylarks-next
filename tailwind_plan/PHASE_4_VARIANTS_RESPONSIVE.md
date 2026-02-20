# ~~Phase 4: Variant & Responsive Migration (shame.css)~~

Address complex classes in `ui/src/css/dp/shame.css` involving media queries and interaction states.

## ~~4.1 State Migration (Hover/Active)~~

- [x] Migrate `hover:preset-tonal-*` classes to component-scoped `:hover` styles.
- [x] Migrate `active:preset-filled-*` classes to component-scoped `:active` styles.
- [x] Use native CSS pseudo-classes in `<style>` blocks.

## ~~4.2 Dark Mode Migration~~

- [x] Identify all `dark:*` usages.
- [x] Replace with `@media (prefers-color-scheme: dark)` in scoped blocks.

## ~~4.3 Responsive Breakpoints~~

- [x] Identify all `sm:*`, `md:*`, `lg:*`, `xl:*` usages.
- [x] Replace with standard media queries:
    - [x] `sm`: `@media (min-width: 40rem)` (640px)
    - [x] `md`: `@media (min-width: 48rem)` (768px)
    - [x] `lg`: `@media (min-width: 64rem)` (1024px)
    - [x] `xl`: `@media (min-width: 80rem)` (1280px)
- [x] **Mobile-First Approach**: Ensure base styles are mobile-first and use `min-width` for overrides.

## 4.4 Shame.css Cleanup

- [x] As each variant is moved to a component, delete the corresponding class from `ui/src/css/dp/shame.css`.
- [x] Aim for 0 lines in `shame.css`.
