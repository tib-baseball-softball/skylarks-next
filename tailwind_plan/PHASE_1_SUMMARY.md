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

The following classes are explicitly marked as **KEEP**:

### Designated Utilities (`ui/src/css/dp/utility/tw_utils.css`)

- `sr-only`, `transform`, `ring`, `blur`, `filter`, `transition`
- `shadow`, `shadow-md`, `shadow-lg`, `shadow-xl`, `shadow-2xl`

### Theme Presets (`ui/src/css/dp/theme/presets.css`)

- **ALL** `preset-*` classes (e.g., `preset-tonal-primary`, `preset-filled-secondary-500`, etc.)

### Component & Global Styles

- **ALL** classes defined in `ui/src/css/dp/partials/*.css` (e.g., `.btn`, `.card`, `.badge`, `.stats`, `.stat`, etc.)
- **ALL** classes defined in `ui/src/css/dp/theme/*.css` **EXCEPT** `tailwind_merged.css` (e.g., `.h1`, `.anchor`,
  `.base-font-*`, etc.)

**Fundamental Rule**: If a class exists in any CSS file other than `tailwind_merged.css` or `shame.css`, it MUST be kept
and NOT targeted for migration.

Cross-reference check: These classes are NOT present in `tailwind_merged.css`.
