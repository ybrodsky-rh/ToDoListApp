package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aquasecurity/table"
)

type Priority string

const (
	Low    Priority = "Low"
	Medium Priority = "Medium"
	High   Priority = "High"
)

const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
)

type Todo struct {
	Title       string
	Completed   bool
	Priority    Priority
	Tags        []string
	CreatedAt   time.Time
	CompletedAt *time.Time
	DueDate     *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string, priority Priority, tags []string, dueDate *time.Time) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		Priority:    priority,
		Tags:        tags,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
		DueDate:     dueDate,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	if len(*todos) == 0 {
		fmt.Println("📭 No tasks yet! Create one with -add flag")
		return
	}

	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Priority", "Completed", "Created At", "Due Date")
	for index, t := range *todos {
		completed := ColorRed + "❌" + ColorReset
		if t.Completed {
			completed = ColorGreen + "✅" + ColorReset
		}

		dueDate := ""
		if t.DueDate != nil {
			dueDate = t.DueDate.Format("2006-01-02")
		}

		priorityColor := ""
		switch t.Priority {
		case High:
			priorityColor = ColorRed + string(t.Priority) + ColorReset
		case Medium:
			priorityColor = ColorYellow + string(t.Priority) + ColorReset
		case Low:
			priorityColor = ColorGreen + string(t.Priority) + ColorReset
		}

		table.AddRow(strconv.Itoa(index), t.Title, priorityColor, completed, t.CreatedAt.Format("2006-01-02"), dueDate)
	}

	table.Render()
	todos.printStats()
	todos.checkDueToday()
}

// Search todos by title
func (todos *Todos) search(query string) Todos {
	var results Todos
	lowercaseQuery := strings.ToLower(query)
	for _, todo := range *todos {
		if strings.Contains(strings.ToLower(todo.Title), lowercaseQuery) {
			results = append(results, todo)
		}
	}
	return results
}

// Filter todos by tag
func (todos *Todos) filterByTag(tag string) Todos {
	var results Todos
	for _, todo := range *todos {
		for _, t := range todo.Tags {
			if t == tag {
				results = append(results, todo)
				break
			}
		}
	}
	return results
}

// Get statistics
func (todos *Todos) getStats() map[string]int {
	stats := map[string]int{
		"total":     len(*todos),
		"completed": 0,
		"pending":   0,
		"high":      0,
	}

	for _, todo := range *todos {
		if todo.Completed {
			stats["completed"]++
		} else {
			stats["pending"]++
		}
		if todo.Priority == High {
			stats["high"]++
		}
	}
	return stats
}

// Print statistics
func (todos *Todos) printStats() {
	stats := todos.getStats()
	fmt.Printf("\n📊 Task Statistics:\n")
	fmt.Printf("   Total: %d | Completed: %d | Pending: %d | High Priority: %d\n\n", 
		stats["total"], stats["completed"], stats["pending"], stats["high"])
}

// Check tasks due today
func (todos *Todos) checkDueToday() {
	today := time.Now().Format("2006-01-02")
	fmt.Printf("⏰ Due Today:\n")
	found := false

	for i, todo := range *todos {
		if todo.DueDate != nil && todo.DueDate.Format("2006-01-02") == today && !todo.Completed {
			fmt.Printf("   [%d] %s\n", i, todo.Title)
			found = true
		}
	}

	if !found {
		fmt.Println("   Nothing due today! 🎉")
	}
}

// Sort by priority
func (todos *Todos) sortByPriority() {
	priorityMap := map[Priority]int{
		High:   1,
		Medium: 2,
		Low:    3,
	}

	sort.Slice(*todos, func(i, j int) bool {
		return priorityMap[(*todos)[i].Priority] < priorityMap[(*todos)[j].Priority]
	})
}

// Sort by date
func (todos *Todos) sortByDate() {
	sort.Slice(*todos, func(i, j int) bool {
		return (*todos)[i].CreatedAt.Before((*todos)[j].CreatedAt)
	})
}

// Delete all completed todos
func (todos *Todos) deleteCompleted() error {
	var result Todos
	for _, todo := range *todos {
		if !todo.Completed {
			result = append(result, todo)
		}
	}
	*todos = result
	return nil
}
