# TASK 1: Add Yinsen to Alpha-Omega Kanban Board

**NAME**: Add Yinsen to kanban dropdown selection
**PRIORITY**: HIGH
**DATE**: 2025-01-06
**TYPE**: task

## Description
Add Yinsen as a selectable project in the Alpha-Omega dashboard dropdown menu. This will allow users to view and manage Yinsen tasks through the kanban board interface.

## Acceptance Criteria
- [ ] Yinsen appears in the project dropdown menu
- [ ] Selecting Yinsen loads its tasks in the kanban board
- [ ] Tasks are properly displayed in their respective columns (Queue, Dev, QA, Done)
- [ ] Task counts are correctly updated for each column

## Technical Implementation
1. Update the Go backend to recognize Yinsen directory (with trailing space)
2. Modify project discovery logic to include Yinsen as a special project
3. Ensure task loading works with the Yinsen directory structure
4. Test locally before creating PR

## Status
Currently implementing the feature in branch: feature/add-yinsen-to-dropdown