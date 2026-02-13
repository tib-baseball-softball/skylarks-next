# Phase 4: Variant & Responsive Migration (shame.css)

Address complex classes in `ui/src/css/dp/shame.css` involving media queries and interaction states.

## 4.1 State Migration (Hover/Active)

- [ ] Migrate `hover:preset-tonal-*` classes to component-scoped `:hover` styles.
- [ ] Migrate `active:preset-filled-*` classes to component-scoped `:active` styles.
- [ ] Use native CSS pseudo-classes in `<style>` blocks.

## 4.2 Dark Mode Migration

- [ ] Identify all `dark:*` usages.
- [ ] Replace with `@media (prefers-color-scheme: dark)` in scoped blocks.
- [ ] *Optional*: If using a theme toggle, use the appropriate selector (e.g.,
  `:global([data-theme='dark']) .my-class`).

## 4.3 Responsive Breakpoints

- [ ] Identify all `sm:*`, `md:*`, `lg:*`, `xl:*` usages.
- [ ] Replace with standard media queries:
    - [ ] `sm`: `@media (min-width: 40rem)` (640px)
    - [ ] `md`: `@media (min-width: 48rem)` (768px)
    - [ ] `lg`: `@media (min-width: 64rem)` (1024px)
    - [ ] `xl`: `@media (min-width: 80rem)` (1280px)
- [ ] **Mobile-First Approach**: Ensure base styles are mobile-first and use `min-width` for overrides.

## 4.4 Shame.css Cleanup

- [ ] As each variant is moved to a component, delete the corresponding class from `ui/src/css/dp/shame.css`.
- [ ] Aim for 0 lines in `shame.css`.
