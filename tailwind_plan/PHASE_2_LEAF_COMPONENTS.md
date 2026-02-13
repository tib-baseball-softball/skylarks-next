# Phase 2: Leaf Component Migration

Target small, reusable components in `ui/src/lib/dp/components` and `ui/src/lib/tib/components/` that primarily use
static utility classes.

## 2.1 Component Identification

- [x] Identify components with zero prop-based styling logic.
- [x] Prioritize components used frequently (e.g., Icons, Badges, simple wrappers).

## 2.2 Migration Workflow (Repeat for each component)

- [ ] **Step 1: Extract Classes** - Identify all `class="..."` attributes in the template.
- [ ] **Step 2: Define Semantic Classes** - Create meaningful class names (e.g., `.root`, `.icon-wrapper`, `.label`)
  instead of `div class="flex items-center"`.
- [ ] **Step 3: Create `<style>` Block** - Add a `<style>` block at the bottom of the `.svelte` file if it doesn't
  exist.
- [ ] **Step 4: Map Utilities to CSS** - Convert Tailwind classes to native CSS properties:
    - [ ] `flex items-center` -> `display: flex; align-items: center;`
    - [ ] `p-4` -> `padding: calc(var(--spacing) * 4);`
    - [ ] `bg-surface-100` -> `background-color: var(--color-surface-100);`
- [ ] **Step 5: Apply Modern CSS** - Use nesting and `:where()` if applicable to keep specificity low.
- [ ] **Step 6: Update Template** - Replace utility strings with the new semantic classes.
- [ ] **Step 7: Verification** - Verify the component looks identical in the UI.

## 2.3 Progress Tracking

- after finishing, always document components migrated here:

- [x] List specific components migrated:
    - [x] `AnnouncementCard.svelte`
    - [x] `AnnouncementCoreContent.svelte`
    - [x] `Avatar.svelte`
    - [x] `Baseball.svelte`
    - [x] `ClubDetailCard.svelte`
    - [x] `DeleteButton.svelte`
    - [x] `EventAttireSection.svelte`
    - [x] `EventCoreInfo.svelte`
    - [x] `EventTypeBadge.svelte`
    - [x] `Footer.svelte`
    - [x] `GameResultIndicator.svelte`
    - [x] `GenericDetailRow.svelte`
    - [x] `MatchTeaserCard.svelte`
    - [x] `MatchTeaserCardRow.svelte`
    - [x] `Pagination.svelte`
    - [x] `PriorityBadge.svelte`
    - [x] `RowCount.svelte`
    - [x] `SeasonSelector.svelte`
    - [x] `icons/Shirt.svelte`
    - [x] `icons/Svelte.svelte`
    - [x] `icons/Pants.svelte`
    - [x] `icons/Cap.svelte`
    - [x] `utils/ProgressRing.svelte`
    - [x] `modal/sheet/sheet-title.svelte`
    - [x] `modal/sheet/sheet-description.svelte`
    - [x] `modal/sheet/sheet-header.svelte`
    - [x] `modal/sheet/sheet-footer.svelte`
    - [x] `modal/sheet/sheet-overlay.svelte`
    - [x] `modal/sheet/sheet-content.svelte`
    - [x] `player/PlayerNumberGraphic.svelte`
    - [x] `toast/ToastContainer.svelte`
    - [x] `datatable/RowsPerPage.svelte`
    - [x] `datatable/ThSort.svelte`
    - [x] `datatable/ThFilter.svelte`
    - [x] `datatable/Search.svelte`
    - [x] `event/match/MatchDetailLocationCard.svelte`
    - [x] `event/TimeSection.svelte`
    - [x] `formElements/Switch.svelte`
    - [x] `auth/LoginBadge.svelte`
    - [x] `meta/TopAppBarTrailing.svelte`
    - [x] `charts/ConicGradient.svelte`
    - [x] `event/IndividualParticipationEditButton.svelte`
    - [x] `team/TeamAddMemberButton.svelte`
    - [x] `stats/StatsBlockContent.svelte`
    - [x] `uniformset/UniformSetInfoCard.svelte`
    - [x] `player/PlayerRow.svelte`
    - [x] `team/TeamListTeaser.svelte`
    - [x] `team/TeamTeaserCard.svelte`
    - [x] `player/PlayerDataCard.svelte`
    - [x] `user/UserDataCard.svelte`
    - [x] `comments/CommentRow.svelte` (Call 8)
    - [x] `comments/CommentSection.svelte` (Call 8)
    - [x] `event/EventTeaser.svelte` (Call 8)
    - [x] `modal/Dialog.svelte` (Call 8)
    - [x] `user/PlayerProfileClubsSection.svelte` (Call 8)
    - [x] `user/PlayerDataProfileSection.svelte` (Call 8)
    - [x] `team/TeamMembersTableContent.svelte` (Call 8)

## 2.4 Summary - Call 8 Phase 2 Continuation

### Components Migrated (7 components)

1. **CommentRow.svelte** - Migrated flex layout classes, spacing utilities (mt-1, gap-2, gap-3), and typography
   classes (font-bold, text-wrap) to scoped CSS
2. **CommentSection.svelte** - Migrated list item layout classes (my-3, md:my-4, flex, gap-2, flex-row-reverse) and form
   spacing (mt-6)
3. **EventTeaser.svelte** - Migrated card layout classes (text-sm, h-full, flex, justify-end, gap-1), typography (
   font-bold, ms-1), and divider spacing (my-2)
4. **Dialog.svelte** - Migrated trigger and header layout classes (flex, gap-1, gap-5, items-center, mb-2) and
   typography (text-xl, font-semibold)
5. **PlayerProfileClubsSection.svelte** - Migrated empty state card layout (flex, flex-col, justify-between, p-4,
   space-y-2, items-center, gap-3, flex-wrap) and typography (font-semibold)
6. **PlayerDataProfileSection.svelte** - Migrated profile card layout (flex, flex-col, justify-between, p-4, font-light,
   mt-2) and responsive actions (gap-2, lg:gap-3)
7. **TeamMembersTableContent.svelte** - Migrated admin actions layout (flex, gap-1, lg:gap-2, justify-end, m-0.5) and
   empty state spacing (py-4)

### Key Patterns Used

- Converted spacing utilities to `calc(var(--spacing) * N)`
- Converted typography utilities to CSS variables (e.g., `var(--font-weight-bold)`, `var(--text-xl)`)
- Used CSS nesting for better organization
- Applied responsive breakpoints with `@media` queries
- Created semantic class names (e.g., `.header`, `.content`, `.actions`, `.title`) instead of utility classes

### Remaining Work for Phase 2

- Most remaining components with Tailwind classes are in the `forms/` directory
- These form components likely have more complex prop-based styling (covered in Phase 5)
- A few section/layout components remain (TeamAdminSection, EventParticipantsOverviewSection, etc.)
- Most simple leaf components have been successfully migrated
