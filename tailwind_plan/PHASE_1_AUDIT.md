# Phase 1: Audit and Preparation

This phase focuses on identifying the scope of the migration and preparing the mapping between Tailwind classes and
native CSS variables.

## 1.1 Class Mapping Inventory

- [ ] Extract all unique utility classes from `ui/src/css/dp/theme/tailwind_merged.css`.
- [ ] Extract all variant classes from `ui/src/css/dp/shame.css`.
- [ ] Create a reference mapping for common utility categories to CSS variables:
    - [ ] **Spacing**: `p-*`, `m-*`, `gap-*`, `space-*` -> `calc(var(--spacing) * n)`
    - [ ] **Colors**: `bg-*`, `text-*`, `border-*` -> `var(--color-*)`
    - [ ] **Typography**: `text-sm`, `text-lg`, etc. -> `var(--text-*)`
    - [ ] **Border Radius**: `rounded-*` -> `var(--radius-*)`
- [ ] Document any classes that do NOT have a direct variable mapping and decide on a replacement strategy.

## 1.2 Usage Audit

- [ ] Run a global search (e.g., `grep -r`) for Tailwind classes in `ui/src/lib/dp/components`.
- [ ] Run a global search for Tailwind classes in `ui/src/routes`.
- [ ] Identify "Hotspot" components:
    - [ ] List top 5 components with the highest utility class count.
    - [ ] List components that use Tailwind classes in logic (dynamic classes).

## 1.3 Utility Safeguard

- [ ] Review `ui/src/css/dp/utility/tw_utils.css`.
- [ ] Cross-reference with `tailwind_merged.css` to ensure no overlap in classes targeted for deletion.
- [ ] Explicitly mark `tw_utils.css` classes as "KEEP" in any internal documentation.
