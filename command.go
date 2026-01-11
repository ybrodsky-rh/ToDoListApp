package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type CmdFlags struct {
	Add       string
	Del       int
	Edit      string
	Toggle    int
	List      bool
	Priority  string
	Tags      string
	DueDate   string
	Search    string
	FilterTag string
	Sort      string
	Stats     bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Priority, "priority", "Medium", "Set priority: Low, Medium, High")
	flag.StringVar(&cf.Tags, "tags", "", "Add tags (comma-separated)")
	flag.StringVar(&cf.DueDate, "due", "", "Due date (YYYY-MM-DD)")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.StringVar(&cf.Search, "search", "", "Search todos by title")
	flag.StringVar(&cf.FilterTag, "tag", "", "Filter todos by tag")
	flag.StringVar(&cf.Sort, "sort", "", "Sort todos: priority, date")
	flag.BoolVar(&cf.Stats, "stats", false, "Show task statistics")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Stats:
		todos.printStats()
	case cf.Search != "":
		results := todos.search(cf.Search)
		if len(results) == 0 {
			fmt.Printf("No todos found matching: %s\n", cf.Search)
		} else {
			fmt.Printf("Search results for: %s\n\n", cf.Search)
			results.print()
		}
	case cf.FilterTag != "":
		results := todos.filterByTag(cf.FilterTag)
		if len(results) == 0 {
			fmt.Printf("No todos found with tag: %s\n", cf.FilterTag)
		} else {
			fmt.Printf("Todos tagged with: %s\n\n", cf.FilterTag)
			results.print()
		}
	case cf.Sort != "":
		if cf.Sort == "priority" {
			todos.sortByPriority()
			fmt.Println("Sorted by priority")
		} else if cf.Sort == "date" {
			todos.sortByDate()
			fmt.Println("Sorted by date")
		} else {
			fmt.Println("Invalid sort option. Use 'priority' or 'date'")
		}
	case cf.Add != "":
		priority := Priority(cf.Priority)
		var tags []string
		if cf.Tags != "" {
			tags = strings.Split(cf.Tags, ",")
			for i := range tags {
				tags[i] = strings.TrimSpace(tags[i])
			}
		}

		var dueDate *time.Time
		if cf.DueDate != "" {
			parsedDate, err := time.Parse("2006-01-02", cf.DueDate)
			if err != nil {
				fmt.Println("Error: invalid due date format. Use YYYY-MM-DD")
				os.Exit(1)
			}
			dueDate = &parsedDate
		}

		todos.add(cf.Add, priority, tags, dueDate)
		fmt.Printf("✅ Added: %s (Priority: %s)\n", cf.Add, priority)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid command")
	}
}
