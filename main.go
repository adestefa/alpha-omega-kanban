package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// Project represents a discovered project
type Project struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Status      string `json:"status"`
	Description string `json:"description"`
	HasYinsen   bool   `json:"has_yinsen"`
	Priority    string `json:"priority"`
	Client      string `json:"client"`
	Deadline    string `json:"deadline"`
	TaskCount   int    `json:"task_count"`
}

// Task represents a task in the kanban system
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	Type        string    `json:"type"`
	CreatedDate time.Time `json:"created_date"`
	Project     string    `json:"project"`
	FilePath    string    `json:"file_path"`
}

// ProjectScanner handles project discovery and parsing
type ProjectScanner struct {
	ClientsPath string
}

// DiscoverProjects scans for all proj-* directories and special projects
func (ps *ProjectScanner) DiscoverProjects() ([]Project, error) {
	var projects []Project

	// First, add Yinsen as a special project in the alpha-omega directory
	yinsenPath := filepath.Join(ps.ClientsPath, "alpha-omega", "yinsen ")
	if _, err := os.Stat(yinsenPath); err == nil {
		project := Project{
			Name:      "Alpha-Omega",
			Path:      yinsenPath,
			HasYinsen: true,
		}

		// Try to read readme for metadata
		ps.parseProjectMetadata(&project, yinsenPath)

		// Determine status based on task activity
		project.Status = ps.determineProjectStatus(yinsenPath)
		
		// Count tasks
		project.TaskCount = ps.countTotalTasks(yinsenPath)
		
		projects = append(projects, project)
	}

	// Then scan for proj-* directories
	entries, err := os.ReadDir(ps.ClientsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read clients directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "proj-") {
			projectPath := filepath.Join(ps.ClientsPath, entry.Name())
			project := Project{
				Name: entry.Name(),
				Path: projectPath,
			}

			// Check if yinsen directory exists
			yinsenPath := filepath.Join(projectPath, "yinsen")
			if _, err := os.Stat(yinsenPath); err == nil {
				project.HasYinsen = true

				// Try to read readme for metadata
				ps.parseProjectMetadata(&project, yinsenPath)

				// Determine status based on task activity
				project.Status = ps.determineProjectStatus(yinsenPath)
				
				// Count tasks
				project.TaskCount = ps.countTotalTasks(yinsenPath)
			}

			projects = append(projects, project)
		}
	}

	return projects, nil
}

// GetProjectTasks retrieves all tasks for a specific project
func (ps *ProjectScanner) GetProjectTasks(projectName string) ([]Task, error) {
	var tasks []Task

	var yinsenPath string
	if projectName == "Alpha-Omega" {
		// Special case for Alpha-Omega project with Yinsen directory (with trailing space)
		yinsenPath = filepath.Join(ps.ClientsPath, "alpha-omega", "yinsen ")
	} else {
		yinsenPath = filepath.Join(ps.ClientsPath, projectName, "yinsen")
	}

	// Define kanban directories and their status
	kanbanDirs := map[string]string{
		"1_queue": "queue",
		"2_dev":   "dev",
		"3_qa":    "qa",
		"4_done":  "done",
	}

	for dir, status := range kanbanDirs {
		dirPath := filepath.Join(yinsenPath, dir)
		if _, err := os.Stat(dirPath); err != nil {
			continue // Skip if directory doesn't exist
		}

		err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() || !strings.HasSuffix(d.Name(), ".md") {
				return nil
			}

			task, err := ps.parseTaskFile(path, status, projectName)
			if err != nil {
				log.Printf("Error parsing task file %s: %v", path, err)
				return nil
			}

			tasks = append(tasks, task)
			return nil
		})

		if err != nil {
			log.Printf("Error walking directory %s: %v", dirPath, err)
		}
	}

	// Sort tasks by ID
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	return tasks, nil
}

// parseTaskFile extracts task information from markdown file
func (ps *ProjectScanner) parseTaskFile(filePath, status, projectName string) (Task, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return Task{}, err
	}

	lines := strings.Split(string(content), "\n")
	task := Task{
		Status:   status,
		Project:  projectName,
		FilePath: filePath,
		Type:     "task",
	}

	// Extract task ID from filename
	filename := filepath.Base(filePath)
	if strings.HasPrefix(filename, "task_") {
		task.ID = strings.TrimSuffix(strings.TrimPrefix(filename, "task_"), ".md")
	} else if strings.HasPrefix(filename, "defect_") {
		task.ID = strings.TrimSuffix(strings.TrimPrefix(filename, "defect_"), ".md")
		task.Type = "defect"
	}

	// Parse content for metadata
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "# TASK") {
			task.Title = line
		} else if strings.HasPrefix(line, "**NAME**:") {
			task.Title = strings.TrimSpace(strings.TrimPrefix(line, "**NAME**:"))
		} else if strings.HasPrefix(line, "**PRIORITY**:") {
			task.Priority = strings.TrimSpace(strings.TrimPrefix(line, "**PRIORITY**:"))
		} else if strings.HasPrefix(line, "**DATE**:") {
			dateStr := strings.TrimSpace(strings.TrimPrefix(line, "**DATE**:"))
			if parsed, err := time.Parse("2006-01-02", dateStr); err == nil {
				task.CreatedDate = parsed
			}
		}
	}

	if task.Title == "" {
		task.Title = fmt.Sprintf("Task %s", task.ID)
	}

	if task.Priority == "" {
		task.Priority = "MEDIUM"
	}

	return task, nil
}

// determineProjectStatus analyzes task distribution to determine project status
func (ps *ProjectScanner) determineProjectStatus(yinsenPath string) string {
	queueCount := ps.countTasksInDir(filepath.Join(yinsenPath, "1_queue"))
	devCount := ps.countTasksInDir(filepath.Join(yinsenPath, "2_dev"))
	qaCount := ps.countTasksInDir(filepath.Join(yinsenPath, "3_qa"))

	if devCount > 0 {
		return "Active"
	} else if queueCount > 0 {
		return "Pending"
	} else if qaCount > 0 {
		return "Testing"
	}

	return "Idle"
}

// countTasksInDir counts markdown files in a directory
func (ps *ProjectScanner) countTasksInDir(dirPath string) int {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return 0
	}

	count := 0
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			count++
		}
	}
	return count
}

// parseProjectMetadata extracts metadata from project readme files
func (ps *ProjectScanner) parseProjectMetadata(project *Project, yinsenPath string) {
	readmePath := filepath.Join(yinsenPath, "readme.md")
	content, err := os.ReadFile(readmePath)
	if err != nil {
		// Try parent directory readme
		parentPath := filepath.Dir(yinsenPath)
		readmePath = filepath.Join(parentPath, "readme.md")
		content, err = os.ReadFile(readmePath)
		if err != nil {
			return
		}
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Extract title as description
		if strings.HasPrefix(line, "# ") && project.Description == "" {
			project.Description = strings.TrimPrefix(line, "# ")
		}
		
		// Look for metadata in **Key:** Value format
		if strings.Contains(line, "**Status:**") {
			parts := strings.SplitN(line, "**Status:**", 2)
			if len(parts) > 1 {
				project.Status = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "**Priority:**") {
			parts := strings.SplitN(line, "**Priority:**", 2)
			if len(parts) > 1 {
				project.Priority = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "**Client:**") {
			parts := strings.SplitN(line, "**Client:**", 2)
			if len(parts) > 1 {
				project.Client = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "**Deadline:**") {
			parts := strings.SplitN(line, "**Deadline:**", 2)
			if len(parts) > 1 {
				project.Deadline = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "**Project:**") && project.Description == "" {
			parts := strings.SplitN(line, "**Project:**", 2)
			if len(parts) > 1 {
				project.Description = strings.TrimSpace(parts[1])
			}
		}
	}
	
	// Default values
	if project.Priority == "" {
		project.Priority = "Medium"
	}
	if project.Client == "" {
		project.Client = "Satori Tech"
	}
}

// countTotalTasks counts all tasks across kanban directories
func (ps *ProjectScanner) countTotalTasks(yinsenPath string) int {
	total := 0
	dirs := []string{"1_queue", "2_dev", "3_qa", "4_done"}
	
	for _, dir := range dirs {
		total += ps.countTasksInDir(filepath.Join(yinsenPath, dir))
	}
	
	return total
}

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for development
	},
}

func main() {
	scanner := &ProjectScanner{
		ClientsPath: "/Users/corelogic/satori-dev/clients",
	}

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// API endpoints
	http.HandleFunc("/api/projects", func(w http.ResponseWriter, r *http.Request) {
		projects, err := scanner.DiscoverProjects()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(projects)
	})

	http.HandleFunc("/api/projects/", func(w http.ResponseWriter, r *http.Request) {
		projectName := strings.TrimPrefix(r.URL.Path, "/api/projects/")
		projectName = strings.TrimSuffix(projectName, "/tasks")

		tasks, err := scanner.GetProjectTasks(projectName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	})

	// WebSocket endpoint for real-time updates
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}
		defer conn.Close()

		// Keep connection alive and send periodic updates
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// Send heartbeat or project updates
				if err := conn.WriteMessage(websocket.TextMessage, []byte("heartbeat")); err != nil {
					return
				}
			}
		}
	})

	// Main dashboard
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	fmt.Println("🚀 Alpha-Omega Dashboard starting on http://localhost:9999")
	fmt.Println("📊 Monitoring projects in /Users/corelogic/satori-dev/clients")

	log.Fatal(http.ListenAndServe(":9999", nil))
}
