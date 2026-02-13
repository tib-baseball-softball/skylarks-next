# Phase 3: Route & Layout Migration

Refactor high-level layout containers and page-specific structures in `ui/src/routes`.

## 3.1 Global Layouts

- [ ] Migrate `ui/src/routes/+layout.svelte`.
- [ ] Identify common layout patterns (e.g., sidebar + main content) and move to scoped CSS or a dedicated partial if
  shared.
- [ ] Replace `h-screen`, `w-full`, `overflow-hidden` on main layout wrappers.

## 3.2 Page-Specific Styles

- [ ] Audit each `+page.svelte` for utility heavy sections.
- [ ] Move "one-off" page styles into the page's `<style>` block.
- [ ] Ensure that responsive layout changes (e.g., `flex-col md:flex-row`) are handled via media queries in the scoped
  block.

## 3.3 Folder-Specific Audits

- [ ] **(tib) Group**: Audit routes within `ui/src/routes/(tib)`.
- [ ] **(app) Group**: Audit routes within `ui/src/routes/(app)`.
- [ ] **Admin/Internal**: Audit any internal or admin routes.

## 3.4 Verification

- [ ] Check page responsiveness on mobile, tablet, and desktop breakpoints.
- [ ] Ensure layout stability (no shifts or scrolls introduced).
