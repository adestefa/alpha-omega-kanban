# Changelog

All notable changes to the Alpha-Omega Dashboard project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.2.0] - 2025-06-11

### Added
- **TASK:5** - Full task details display in modal popup
  - Task content now shown with markdown formatting
  - Metadata grid showing status, priority, type, and dates
  - Scrollable content area for long task descriptions
  - Improved modal layout and sizing
- **TASK:6** - Green "Run Claude" button on task cards
  - Quick access to execute Claude Code without opening modal
  - Play icon for visual clarity
  - Shows command preview when clicked
  - Proper event handling to prevent conflicts

### Changed
- Task struct now includes Content field for full markdown
- Modal window increased to max-w-4xl for better content display
- Task cards now have slightly more spacing to accommodate button

## [1.1.0] - 2025-06-11

### Added
- **TASK:3** - Multi-directory project discovery
  - Scans clients/, engagements/, and platform/ directories
  - Supports platform projects with different naming scheme
  - Discovered 5 yinsen projects across all directories

### Changed
- ProjectScanner now uses BasePaths array instead of single ClientsPath
- Server port changed from 8080 to 9999 to avoid conflicts
- Start script shows all monitored directories

## [1.0.0] - 2025-06-10

### Added
- **TASK:1** - Alpha-Omega (Yinsen) project in dropdown menu
- **TASK:2** - Project leadspace with metadata display
- Core kanban board functionality
- Real-time WebSocket connections
- Project discovery and task parsing
- Queue-based task management (1_queue → 2_dev → 3_qa → 4_done)

### Initial Features
- Lightweight Go backend with zero dependencies
- HTML/Tailwind/JavaScript frontend
- Project dropdown selector
- Kanban columns with task counts
- Task cards with priority indicators
- Modal for task details
- WebSocket for future real-time updates