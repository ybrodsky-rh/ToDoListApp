# 📝 To-Do List App

A simple, efficient command-line to-do list manager built with **Go**. Manage your tasks directly from the terminal with an intuitive interface.
This project was based on the following YouTube tutorial: "How to Build A CLI Todo App in Go" by Coding with Patrik. Link: https://www.youtube.com/watch?v=g16Zf0KQEWI 
It is intended for learning purposes.

## ✨ Features

- ✅ **Add tasks** - Quickly add new to-do items with priority levels
- 🗑️ **Delete tasks** - Remove completed or unwanted tasks
- ✏️ **Edit tasks** - Update task titles anytime
- ✔️ **Toggle completion** - Mark tasks as done/undone with timestamps
- 🎯 **Priority levels** - Categorize tasks as Low, Medium, or High priority
- 🏷️ **Tags/Categories** - Organize tasks with custom tags
- 📅 **Due dates** - Set deadlines for your tasks with reminders
- 🔍 **Search functionality** - Find tasks by title
- 🏷️ **Filter by tags** - View tasks by category
- 📊 **Statistics** - See task completion progress and priority breakdown
- 🔄 **Sorting** - Sort tasks by priority or creation date
- 💾 **Persistent storage** - All tasks saved to JSON file
- 📊 **Beautiful CLI table** - View all tasks in a colored, formatted table
- 🎨 **Color-coded output** - Priorities and completion status shown with colors

## 🚀 Getting Started

### Prerequisites
- Go 1.18 or higher installed on your machine

### Installation

1. Clone the repository:
```bash
git clone https://github.com/ybrodsky-rh/ToDoListApp.git
cd ToDoListApp
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build
```

## 📖 Usage

### Add a Task (Basic)
```bash
./ToDoListApp -add "Buy groceries"
./ToDoListApp -add "Finish project report"
```

### Add a Task (With Priority, Tags & Due Date)
```bash
./ToDoListApp -add "Learn Go" -priority "High" -tags "learning,golang" -due "2026-01-20"
./ToDoListApp -add "Buy groceries" -priority "Low" -tags "shopping,home"
./ToDoListApp -add "Project deadline" -priority "High" -tags "work" -due "2026-01-25"
```

### List All Tasks
```bash
./ToDoListApp -list
```
Shows all tasks in a colored table with priorities, due dates, and task statistics.

### Search Tasks by Title
```bash
./ToDoListApp -search "Go"
```
Finds all tasks containing the search term (case-insensitive).

### Filter Tasks by Tag
```bash
./ToDoListApp -tag "work"
./ToDoListApp -tag "home"
```

### View Statistics
```bash
./ToDoListApp -stats
```
Displays: Total tasks, Completed, Pending, and High Priority count.

### Sort Tasks
```bash
./ToDoListApp -sort "priority"  # Sort by priority (High → Medium → Low)
./ToDoListApp -sort "date"      # Sort by creation date (oldest first)
```

### Toggle Task Completion
```bash
./ToDoListApp -toggle 0
```

### Delete a Task
```bash
./ToDoListApp -del 0
```

### Edit a Task
```bash
./ToDoListApp -edit "0:Updated task title"
```

## 📂 Project Structure

```
ToDoListApp/
├── main.go          # Entry point of the application
├── todo.go          # Todo struct and operations (add, delete, toggle, etc.)
├── command.go       # Command-line flag parsing
├── storage.go       # JSON file I/O operations
├── go.mod           # Go module definition
├── go.sum           # Dependency checksums
├── todos.json       # Persistent storage file (auto-generated)
└── README.md        # This file
```

## 🔧 How It Works

1. **Load** - Reads existing todos from `todos.json`
2. **Parse** - Parses command-line flags (`-add`, `-del`, `-toggle`, `-edit`, `-list`, `-search`, `-tag`, `-sort`, `-stats`, etc.)
3. **Execute** - Performs the requested operation
4. **Save** - Persists changes back to `todos.json`

## 📋 Available CLI Flags

| Flag | Description | Example |
|------|-------------|---------|
| `-add` | Add a new task | `-add "Buy milk"` |
| `-priority` | Set priority level | `-priority "High"` (Low, Medium, High) |
| `-tags` | Add tags (comma-separated) | `-tags "work,urgent,learning"` |
| `-due` | Set due date | `-due "2026-01-20"` (YYYY-MM-DD) |
| `-list` | List all tasks | `-list` |
| `-search` | Search by title | `-search "golang"` |
| `-tag` | Filter by tag | `-tag "work"` |
| `-sort` | Sort tasks | `-sort "priority"` or `-sort "date"` |
| `-stats` | Show statistics | `-stats` |
| `-toggle` | Mark as complete/incomplete | `-toggle 0` |
| `-edit` | Edit task title | `-edit "0:New title"` |
| `-del` | Delete a task | `-del 0` |

## 💡 Key Go Concepts Used

- **Pointers** (`*`) - Modify original data in methods
- **Methods** - Functions attached to types using receivers
- **Generics** (`[T any]`) - Flexible Storage type for any data
- **Error Handling** - Proper error returns and validation
- **JSON marshaling** - Serialize/deserialize data
- **Slices** - Dynamic arrays for managing todos
- **Type constants** - Custom Priority type and color constants
- **Sorting** - Using `sort.Slice()` for custom sorting logic
- **String manipulation** - Parsing dates, tags, and search queries

## 📦 Dependencies

- [`github.com/aquasecurity/table`](https://github.com/aquasecurity/table) - Beautiful CLI table formatting

## 🛠️ Development

### Running Tests (if applicable)
```bash
go test ./...
```

### Modifying Dependencies
```bash
go mod tidy
```

## 📝 Example Workflow

```bash
# Add some tasks with different priorities
./ToDoListApp -add "Learn Go basics" -priority "High" -tags "learning,golang" -due "2026-01-20"
./ToDoListApp -add "Build a CLI project" -priority "High" -tags "work,golang"
./ToDoListApp -add "Buy groceries" -priority "Low" -tags "shopping,home"
./ToDoListApp -add "Code review" -priority "Medium" -tags "work"

# View all tasks with colors and statistics
./ToDoListApp -list

# Search for specific tasks
./ToDoListApp -search "Go"

# Filter tasks by tag
./ToDoListApp -tag "work"
./ToDoListApp -tag "shopping"

# Sort tasks by priority
./ToDoListApp -sort "priority"

# View quick statistics
./ToDoListApp -stats

# Mark first task as done
./ToDoListApp -toggle 0

# Edit a task
./ToDoListApp -edit "1:Completed my first project"

# Delete a task
./ToDoListApp -del 2

# View updated list
./ToDoListApp -list
```

### Output Example
When you run `-list`, you'll see:
```
┌───┬─────────────────────────┬──────────┬───────────┬────────────┬────────────┐
│ # │        Title            │ Priority │ Completed │ Created At │  Due Date  │
├───┼─────────────────────────┼──────────┼───────────┼────────────┼────────────┤
│ 0 │ Learn Go basics          │ High     │ ✅        │ 2026-01-11 │ 2026-01-20 │
│ 1 │ Build a CLI project      │ High     │ ❌        │ 2026-01-11 │            │
│ 2 │ Buy groceries            │ Low      │ ❌        │ 2026-01-11 │            │
│ 3 │ Code review              │ Medium   │ ❌        │ 2026-01-11 │            │
└───┴─────────────────────────┴──────────┴───────────┴────────────┴────────────┘

📊 Task Statistics:
   Total: 4 | Completed: 1 | Pending: 3 | High Priority: 2

⏰ Due Today:
   Nothing due today! 🎉
```

## 🎓 Learning Resource

This project is great for learning:
- Go fundamentals (structs, methods, receivers)
- Command-line argument parsing with the `flag` package
- File I/O and JSON handling (marshaling/unmarshaling)
- Generic types (Go 1.18+)
- Error handling patterns
- Custom type definitions and constants
- Sorting and searching algorithms
- Working with time/date handling
- Building practical CLI applications

## 📄 License

This project is open source and available under the MIT License.

## 🤝 Contributing

Feel free to fork this repository and submit pull requests for any improvements!

## 📧 Contact

For questions or suggestions, please open an issue on GitHub.

---

**Happy task managing! 🚀**
