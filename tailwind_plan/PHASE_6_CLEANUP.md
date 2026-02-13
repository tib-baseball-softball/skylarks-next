# Phase 6: Final Cleanup and Deletion

Final steps to remove leftover artifacts and ensure UI consistency.

## 6.1 Unused Class Audit

- [ ] Run a final search for ALL classes remaining in `tailwind_merged.css`.
- [ ] If 0 occurrences exist in `ui/src`, safely delete them.

## 6.2 File Removal

- [ ] Empty and delete `ui/src/css/dp/theme/tailwind_merged.css`.
- [ ] Empty and delete `ui/src/css/dp/shame.css`.
- [ ] Remove imports from:
    - [ ] `ui/src/css/dp/app.css`
    - [ ] `ui/src/css/dp/index.css`

## 6.3 Documentation Update

- [ ] Update any developer documentation (README, etc.) to reflect that Tailwind utilities (except for designated ones
  in `tw_utils.css`) are no longer used.
