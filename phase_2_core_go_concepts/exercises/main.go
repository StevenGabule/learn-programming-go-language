package main

import (
	"errors"
	"fmt"
	"time"
)

// Custom errors
var (
	ErrTaskNotFound = errors.New("task not found")
	ErrInvalidTask  = errors.New("invalid task")
)

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

func (p Priority) String() string {
	return [...]string{"Low", "Medium", "High"}[p]
}

type Task struct {
	ID          int
	Title       string
	Description string
	Priority    Priority
	Completed   bool
	CreatedAt   time.Time
}

type TaskManager struct {
	tasks  map[int]*Task
	nextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

func (tm *TaskManager) Add(title, description string, priority Priority) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("%w: title cannot be empty", ErrInvalidTask)
	}

	task := &Task{
		ID:          tm.nextID,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   time.Now(),
	}

	tm.tasks[task.ID] = task
	tm.nextID++
	return task, nil
}

func (tm *TaskManager) Get(id int) (*Task, error) {
	task, exists := tm.tasks[id]
	if !exists {
		return nil, fmt.Errorf("task %d: %w", id, ErrTaskNotFound)
	}

	return task, nil
}

func (tm *TaskManager) Complete(id int) error {
	task, err := tm.Get(id)
	if err != nil {
		return err
	}

	task.Completed = true
	return nil
}

func (tm *TaskManager) Delete(id int) error {
	if _, exists := tm.tasks[id]; !exists {
		return fmt.Errorf("delete task %d: %w", id, ErrTaskNotFound)
	}

	delete(tm.tasks, id)
	return nil
}

func (tm *TaskManager) List() []*Task {
	result := make([]*Task, 0, len(tm.tasks))
	for _, task := range tm.tasks {
		result = append(result, task)
	}

	return result
}

func (t *Task) String() string {
	status := "[ ]"
	if t.Completed {
		status = "[✓]"
	}
	return fmt.Sprintf("%s #%d [%s] %s", status, t.ID, t.Priority, t.Title)
}

func main() {
	tm := NewTaskManager()

	// Add Tasks
	task1, _ := tm.Add("Learn GO Basic", "Complete phase 1 and 2", High)
	task2, _ := tm.Add("Build a project", "Create a REST API", Medium)
	tm.Add("Read Effective Go", "", Low)

	fmt.Println("All Tasks:")
	for _, t := range tm.List() {
		fmt.Println(t)
	}

	// Complete a task
	tm.Complete(task1.ID)
	fmt.Println("\nAfter completing task 1: ")
	fmt.Println(task1)

	// Error handling
	_, err := tm.Get(999)
	if errors.Is(err, ErrTaskNotFound) {
		fmt.Println("\nError:", err)
	}

	// Try adding invalid task
	_, err = tm.Add("", "no title", High)
	if errors.Is(err, ErrInvalidTask) {
		fmt.Println("Validation error:", err)
	}

	// Delete
	tm.Delete(task2.ID)
	fmt.Println("\nAfter deleting task 2:")
	for _, t := range tm.List() {
		fmt.Println(t)
	}
}
