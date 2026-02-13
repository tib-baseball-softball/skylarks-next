# Refined Refactoring Plan: Tailwind CSS Migration

This plan details the steps to completely remove leftover Tailwind utility classes from
`ui/src/css/dp/theme/tailwind_merged.css` and `ui/src/css/dp/shame.css` while transitioning to scoped plain CSS in
Svelte components and modern CSS practices.

#### Phase 1: Audit and Preparation

[See Detailed Checklist: PHASE_1_AUDIT.md](./PHASE_1_AUDIT.md)

* **1.1. Class Mapping Inventory**: Extract all unique class selectors from `tailwind_merged.css` and `shame.css`. Map
  them to their native CSS equivalents using the project's established variables (e.g., `--spacing`, `--color-*`).
    * *Example*: `.p-4` -> `padding: calc(var(--spacing) * 4);`
* **1.2. Usage Audit**: Run a global search for each extracted class in `.svelte` files within `ui/src/lib` and
  `ui/src/routes` to identify "hotspot" components with heavy Tailwind usage.
* **1.3. Utility Safeguard**: Verify that no classes from `ui/src/css/dp/utility/tw_utils.css` are accidentally targeted
  for migration, as these are designated to be kept.

#### Phase 2: Leaf Component Migration (Simple Cases)

[See Detailed Checklist: PHASE_2_LEAF_COMPONENTS.md](./PHASE_2_LEAF_COMPONENTS.md)

Target components in `ui/src/lib/dp/components` that use static utility classes and do not pass them as props.

* **2.1. Spacing & Layout**: Replace `p-*`, `m-*`, `gap-*`, `flex`, `grid`, `items-*`, and `justify-*` with scoped CSS.
* **2.2. Typography & Colors**: Replace `text-*`, `font-*`, and `bg-*` classes.
* **2.3. Workflow**:
    1. Move classes from HTML to a `<style>` block.
    2. Use semantic class names (e.g., `.container`, `.title`, `.wrapper`).
    3. Remove the migrated classes from `tailwind_merged.css`.

#### Phase 3: Route & Layout Migration

[See Detailed Checklist: PHASE_3_ROUTES_LAYOUTS.md](./PHASE_3_ROUTES_LAYOUTS.md)

Apply the same logic as Phase 2 to the components and HTML structures in `ui/src/routes`.

* **3.1. Page Layouts**: Refactor high-level layout containers in `+layout.svelte` and `+page.svelte` files.
* **3.2. Scoped Styling**: Ensure page-specific styles are kept within the route's Svelte file to prevent global CSS
  bloat.

#### Phase 4: Variant & Responsive Migration (Shame.css)

[See Detailed Checklist: PHASE_4_VARIANTS_RESPONSIVE.md](./PHASE_4_VARIANTS_RESPONSIVE.md)

Address the complex classes in `ui/src/css/dp/shame.css` involving media queries and states.

* **4.1. Hover/Active States**: Replace classes like `hover:preset-tonal-primary` with native `:hover` and `:active`
  pseudo-classes in the component's `<style>` block.
* **4.2. Dark Mode**: Replace `dark:*` classes with `@media (prefers-color-scheme: dark)` or the project's specific dark
  mode selector (e.g., `[data-theme='dark']`) within the scoped styles.
* **4.3. Responsive Breakpoints**: Replace `md:*`, `lg:*` with standard `@media` queries using the project's breakpoint
  variables (if available) or standard values (e.g., `64rem`).

#### Phase 5: Complex Prop-based Migration

[See Detailed Checklist: PHASE_5_COMPLEX_PROPS.md](./PHASE_5_COMPLEX_PROPS.md)

Handle components that receive Tailwind classes via props (e.g., `class`, `panelPadding`, `buttonClasses`).

* **5.1. Simple Prop Mapping**: If a prop only passes a few variations of Tailwind classes, replace the prop with a
  semantic one (e.g., `size="small" | "large"`) and handle styling internally.
* **5.2. Marking Complex Logic**: As per the core plan, if a component's inner logic depends heavily on Tailwind class
  manipulation that cannot be easily refactored without touching logic, mark it with
  `TODO: tailwind-migration [description]` and skip for now.
* **5.3. CSS Variable Injection**: For highly configurable components, consider using CSS custom properties (e.g.,
  `--component-padding`) passed via the `style` attribute instead of utility classes.

#### Phase 6: Final Cleanup and Deletion

[See Detailed Checklist: PHASE_6_CLEANUP.md](./PHASE_6_CLEANUP.md)

* **6.1. Unused Class Pruning**: Periodically run an audit to find classes in `tailwind_merged.css` and `shame.css` that
  have zero occurrences in the codebase.
* **6.2. File Removal**: Once `tailwind_merged.css` and `shame.css` are empty:
    1. Delete the files.
    2. Remove their `@import` statements in `ui/src/css/dp/app.css` and `ui/src/css/dp/index.css`.
* **6.3. Visual Verification**: Conduct a final review of the UI to ensure consistency and that no styles were lost
  during the migration.

#### Important Constraints Reminder

* **DO NOT** touch any CSS files other than `tailwind_merged.css` and `shame.css`.
* **DO NOT** touch backend files.
* **DO NOT** change TypeScript files or non-styling logic in Svelte files.
* **ALWAYS** use Svelte 5 best practices (e.g., snippets for complex props) and modern CSS (e.g., nesting, `:where()`).
