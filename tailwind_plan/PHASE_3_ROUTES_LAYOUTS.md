# Phase 3: Route & Layout Migration

Refactor high-level layout containers and page-specific structures in `ui/src/routes`.

## 3.1 Global Layouts

- [x] Migrate `ui/src/routes/+layout.svelte` (Call 10)
- [x] Identify common layout patterns (e.g., sidebar + main content) and move to scoped CSS or a dedicated partial if
  shared.
- [x] Replace `h-screen`, `w-full`, `overflow-hidden` on main layout wrappers.

## 3.2 Page-Specific Styles

- [x] Audit/Migrate `ui/src/routes/(tib)/favorite/+page.svelte` (Call 10)
- [x] Audit/Migrate `ui/src/routes/(dp)/(protected)/account/announcements/[announcement]/+page.svelte` (Call 10)
- [x] Audit/Migrate `ui/src/routes/(dp)/(protected)/account/team/[id]/+page.svelte` (Call 10)
- [ ] Audit more `+page.svelte` for utility heavy sections.
- [x] Move "one-off" page styles into the page's `<style>` block.
- [x] Ensure that responsive layout changes (e.g., `flex-col md:flex-row`) are handled via media queries in the scoped
  block.

## 3.3 Folder-Specific Audits

- [x] **(tib) Group**: Partial audit - Favorite page migrated.
- [x] **(app) Group**: Partial audit - Team detail and Announcement detail migrated.
- [ ] **Admin/Internal**: Audit any internal or admin routes.

## 3.4 Verification

- [ ] Check page responsiveness on mobile, tablet, and desktop breakpoints.
- [ ] Ensure layout stability (no shifts or scrolls introduced).
