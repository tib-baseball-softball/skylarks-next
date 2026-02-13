# Phase 2: Leaf Component Migration

Target small, reusable components in `ui/src/lib/dp/components` that primarily use static utility classes.

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

- [/] List specific components migrated:
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
