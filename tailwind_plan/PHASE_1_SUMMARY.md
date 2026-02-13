# Phase 1 Summary: Audit and Preparation

## 1.1 Class Mapping Inventory

- Extracted unique utility classes from `tailwind_merged.css` (saved in `tailwind_classes.txt`).
- Extracted variant classes from `shame.css` (saved in `shame_classes.txt`).
- Mapping established in `tailwind_plan/MAPPING.md`.

## 1.2 Usage Audit

### Top 5 Component Hotspots (by utility class count)

1. `ui/src/lib/dp/components/player/PlayerDataCard.svelte` (40)
2. `ui/src/lib/dp/components/event/EventParticipantsOverviewSection.svelte` (37)
3. `ui/src/lib/dp/components/event/EventPageAdminSection.svelte` (32)
4. `ui/src/lib/dp/components/team/TeamAdminSection.svelte` (30)
5. `ui/src/lib/dp/components/forms/TeamGamesModal.svelte` (30)

### Top 3 Route Hotspots

1. `ui/src/routes/(dp)/(protected)/account/announcements/[announcement]/+page.svelte` (38)
2. `ui/src/routes/(tib)/favorite/+page.svelte` (28)
3. `ui/src/routes/(dp)/(protected)/account/team/[id]/+page.svelte` (22)

### Components with Dynamic/Logic-based Classes

- `CommentRow.svelte` (conditional preset classes)
- `CommentSection.svelte` (responsive margin and flex-direction logic)
- `ParticipationForm.svelte` (conditional state-based background/text classes)
- `sheet-footer.svelte` / `sheet-header.svelte` etc. (combining internal utilities with passed `className` prop)

## 1.3 Utility Safeguard

The following classes from `ui/src/css/dp/utility/tw_utils.css` are marked as **KEEP**:

- `sr-only`
- `transform`
- `shadow`, `shadow-md`, `shadow-lg`, `shadow-xl`, `shadow-2xl`
- `ring`
- `blur`
- `filter`
- `transition`

Cross-reference check: These classes are NOT present in `tailwind_merged.css`.
