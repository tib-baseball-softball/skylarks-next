# Phase 1: Audit and Preparation

This phase focuses on identifying the scope of the migration and preparing the mapping between Tailwind classes and
native CSS variables.

## 1.1 Class Mapping Inventory

- [x] Extract all unique utility classes from `ui/src/css/dp/theme/tailwind_merged.css`.
- [x] Extract all variant classes from `ui/src/css/dp/shame.css`.
- [x] Create a reference mapping for common utility categories to CSS variables:
    - [x] **Spacing**: `p-*`, `m-*`, `gap-*`, `space-*` -> `calc(var(--spacing) * n)`
    - [x] **Colors**: `bg-*`, `text-*`, `border-*` -> `var(--color-*)`
    - [x] **Typography**: `text-sm`, `text-lg`, etc. -> `var(--text-*)`
    - [x] **Border Radius**: `rounded-*` -> `var(--radius-*)`
- [x] Document any classes that do NOT have a direct variable mapping and decide on a replacement strategy.

## 1.2 Usage Audit

- [x] Run a global search (e.g., `grep -r`) for Tailwind classes in `ui/src/lib/dp/components`.
- [x] Run a global search for Tailwind classes in `ui/src/routes`.
- [x] Identify "Hotspot" components:
    - [x] List top 5 components with the highest utility class count.
    - [x] List components that use Tailwind classes in logic (dynamic classes).

## 1.3 Utility Safeguard

- [x] Review `ui/src/css/dp/utility/tw_utils.css`.
- [x] Review `ui/src/css/dp/theme/presets.css` and all files in `ui/src/css/dp/partials/`.
- [x] Cross-reference with `tailwind_merged.css` to ensure no overlap in classes targeted for deletion.
- [x] Explicitly mark all classes found outside `tailwind_merged.css` and `shame.css` as "KEEP" (especially `preset-*`
  classes).
