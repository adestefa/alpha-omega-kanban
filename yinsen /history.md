# Alpha-Omega Dashboard Project Plan

## Project Overview
Alpha-Omega is a lightweight, real-time Kanban dashboard that visualizes the status of multiple Yinsen-managed projects without duplicating the underlying file system structure. It leverages existing project directories and kanban queue systems as the single source of truth.

## Core Philosophy
- **Lightweight**: HTML, Tailwind CSS, plain JavaScript
- **Non-Invasive**: No modification to existing yinsen workflow
- **Real-Time**: Live updates showing tasks moving between queues
- **Scalable**: Easy addition of new projects
- **Extensible**: Favorites system for cross-project priority tracking

## Architecture Components

### 1. Project Discovery & Parsing
**Directory Structure Analysis:**
```
/Users/corelogic/satori-dev/clients/
├── proj-<name>/
│   ├── readme.md (project description & metadata)
    └── yinsen
        ├── 1_queue
        │   └── task_1.md
        ├── 2_dev
        ├── 3_qa
        ├── 4_done
        ├── artifacts
        ├── task_list.md
        ├── test_plan.md    
        ├── bootstrap.md
        ├── git_ops.sh
        ├── history.md
        └── readme.md
```

**Data Extraction Requirements:**
- Parse all `proj-*` directories under clients/
- Extract project metadata from readme.md files
- Read task_list.md for current assignments
- Scan kanban subdirectories for task/defect files
- Cross-reference tasks between queues and master list

### 2. Standardized Project Metadata
**readme.md Standard Format:**
```markdown
# Project Name
**Status:** Active/Paused/Complete
**Priority:** High/Medium/Low
**Client:** Client Name
**Description:** Brief project description
**Tech Stack:** Technologies used
**Deadline:** Target completion date
```

### 3. Dashboard Features

#### Core Kanban View
- **Project Selector**: Dropdown menu for project switching
- **Queue Columns**: backlog → ready → in-progress → review → done
- **Defects Lane**: Separate row/section for defect tracking
- **Task Cards**: Show task details with action buttons

#### Task Card Components
- Task title and brief description
- Current status and queue location
- Priority indicator
- Action buttons: Start, Stop, View Details
- Visual indicators for movement/updates

#### Favorites Dashboard
- Cross-project priority task aggregation
- Flagged high-priority items from all projects
- Quick access to critical tasks regardless of project

#### Real-Time Updates
- File system watchers for queue directories
- WebSocket connections for live updates
- Visual animations for task transitions
- Auto-refresh capabilities with manual refresh option

### 4. Technical Implementation

#### Backend (Golang)
```
dashboard/
├── main.go (HTTP server & WebSocket handler)
├── scanner/
│   ├── project_discovery.go
│   ├── task_parser.go
│   └── file_watcher.go
├── api/
│   ├── projects.go
│   ├── tasks.go
│   └── favorites.go
└── static/
    ├── index.html
    ├── styles.css (Tailwind)
    └── app.js
```

#### Frontend (HTML/Tailwind/JS)
- Single-page application
- Responsive design for multiple screen sizes
- Drag-and-drop functionality (future enhancement)
- Real-time WebSocket connection handling

#### API Endpoints
- `GET /api/projects` - List all discovered projects
- `GET /api/projects/{name}/tasks` - Get tasks for specific project
- `GET /api/favorites` - Get flagged priority tasks
- `POST /api/tasks/{id}/action` - Execute task actions (start/stop)
- `WebSocket /ws` - Real-time updates

### 5. File System Integration

#### Task File Format Detection
- Markdown files with standardized frontmatter
- YAML headers for metadata extraction
- Content parsing for task descriptions

#### Queue Monitoring
- Watch for file creation/deletion/moves
- Track timestamp changes for activity detection
- Maintain state consistency across refreshes

### 6. Development Phases

#### Phase 1: Foundation
- [ ] Project discovery and parsing logic
- [ ] Basic Golang HTTP server setup
- [ ] Static HTML dashboard skeleton
- [ ] Project dropdown and basic navigation

#### Phase 2: Core Kanban
- [ ] Task parsing and display
- [ ] Queue visualization
- [ ] Basic task card rendering
- [ ] Project switching functionality

#### Phase 3: Real-Time Features
- [ ] File system watchers
- [ ] WebSocket implementation
- [ ] Live task movement animations
- [ ] Auto-refresh mechanisms

#### Phase 4: Interactive Features
- [ ] Task action buttons (start/stop/details)
- [ ] Favorites system implementation
- [ ] Priority flagging
- [ ] Cross-project aggregation

#### Phase 5: Polish & Enhancement
- [ ] Visual improvements and animations
- [ ] Performance optimization
- [ ] Error handling and resilience
- [ ] Documentation and deployment

### 7. Benefits

#### For Development Teams
- **Unified View**: See all projects at a glance
- **Real-Time Awareness**: Know when tasks move and change
- **Priority Management**: Focus on critical items across projects
- **Non-Disruptive**: Works with existing yinsen workflow

#### For Project Management
- **Status Visibility**: Clear project health indicators
- **Resource Allocation**: See where work is concentrated
- **Bottleneck Identification**: Spot queue buildups
- **Cross-Project Coordination**: Manage dependencies and priorities

### 8. Technical Considerations

#### Performance
- Efficient file system scanning
- Minimal memory footprint
- Fast project switching
- Optimized real-time updates

#### Scalability
- Support for unlimited projects
- Efficient queue monitoring
- Lazy loading for large task sets
- Configurable refresh rates

#### Reliability
- Graceful handling of missing files
- Recovery from file system changes
- Robust error reporting
- Consistent state management

## Next Steps
1. Set up basic project structure
2. Implement project discovery scanner
3. Create minimal HTML dashboard
4. Test with existing yinsen projects
5. Iterate and enhance based on real usage

---

*This dashboard will transform how we monitor and manage multiple simultaneous development projects, creating a unified command center for the Satori development army.*