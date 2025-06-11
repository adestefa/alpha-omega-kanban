# TASK 6: Add Run Button to Task Cards

**NAME**: Update the card face to include a small green button to "run" claude
**PRIORITY**: HIGH
**DATE**: 2025-06-11
**TYPE**: task

## Description
Add a green "Run" button directly on each task card that allows quick execution of Claude Code for that specific task without opening the modal.

## Acceptance Criteria
- [ ] Green run button appears on each task card
- [ ] Button is visually distinct but not overwhelming
- [ ] Button has appropriate hover states
- [ ] Clicking button triggers Claude Code execution
- [ ] Button includes icon (play symbol)
- [ ] Works for all task types (tasks and defects)

## Technical Implementation
1. Update createTaskCard function to include run button
2. Add CSS styling for the button
3. Implement click handler for quick execution
4. Ensure button doesn't interfere with card click