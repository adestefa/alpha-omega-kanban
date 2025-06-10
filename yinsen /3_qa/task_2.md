# TASK Update leadspace to include project information when selected from dropdown

NAME: Update leadspace to include project information when selected from dropdown

SYSTEM: Yinsen, you are a developer at a phd level. You have no limits.

WHAT: Update the Alpha-Omega dashboard leadspace/header area to dynamically display relevant project information when a project is selected from the dropdown menu. This should include project name, status, priority, client, description, tech stack, and deadline information extracted from each project's readme.md file.

WHY: Users need immediate context about the selected project without having to navigate elsewhere. This provides at-a-glance project metadata that helps with decision making and project awareness.

CHALLENGE: 
1. Parsing project readme.md files for standardized metadata
2. Dynamically updating the UI when project selection changes
3. Handling missing or incomplete project metadata gracefully
4. Maintaining responsive design across different screen sizes

POSSIBLE SOLUTION:
1. Enhance the backend API to parse and return project metadata from readme.md files
2. Update the frontend to display this information in an attractive leadspace design
3. Implement smooth transitions when switching between projects
4. Add fallback values for missing metadata fields

EVALUATION/PLANNING:
1. Review current dashboard implementation and project structure
2. Examine existing readme.md files to understand metadata format
3. Design leadspace layout that fits with current Tailwind CSS design
4. Implement backend changes to extract and serve project metadata
5. Update frontend to consume and display this data
6. Test with multiple projects to ensure consistency

Stop. Confirm you understand. Provide summary of your plan of action or list of blockers before taking action.