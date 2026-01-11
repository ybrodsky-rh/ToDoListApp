# 📝 To-Do List App

A simple, efficient command-line to-do list manager built with **Go**. Manage your tasks directly from the terminal with an intuitive interface.

## ✨ Features

- ✅ **Add tasks** - Quickly add new to-do items
- 🗑️ **Delete tasks** - Remove completed or unwanted tasks
- ✏️ **Edit tasks** - Update task titles anytime
- ✔️ **Toggle completion** - Mark tasks as done/undone with timestamps
- 💾 **Persistent storage** - All tasks saved to JSON file
- 📊 **Beautiful CLI table** - View all tasks in a formatted table

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

### Add a Task
```bash
./ToDoListApp -add "Buy groceries"
./ToDoListApp -add "Finish project report"
```

### List All Tasks
```bash
./ToDoListApp -list
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
2. **Parse** - Parses command-line flags (`-add`, `-del`, `-toggle`, `-edit`, `-list`)
3. **Execute** - Performs the requested operation
4. **Save** - Persists changes back to `todos.json`

## 💡 Key Go Concepts Used

- **Pointers** (`*`) - Modify original data in methods
- **Methods** - Functions attached to types using receivers
- **Generics** (`[T any]`) - Flexible Storage type for any data
- **Error Handling** - Proper error returns and validation
- **JSON marshaling** - Serialize/deserialize data
- **Slices** - Dynamic arrays for managing todos

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
# Add some tasks
./ToDoListApp -add "Learn Go"
./ToDoListApp -add "Build a project"
./ToDoListApp -add "Deploy to GitHub"

# View all tasks
./ToDoListApp -list

# Mark first task as done
./ToDoListApp -toggle 0

# Edit a task
./ToDoListApp -edit "1:Build 3 projects"

# Delete a task
./ToDoListApp -del 2

# View updated list
./ToDoListApp -list
```

## 🎓 Learning Resource

This project is great for learning:
- Go fundamentals (structs, methods, receivers)
- Command-line argument parsing
- File I/O and JSON handling
- Generic types (Go 1.18+)
- Error handling patterns

## 📄 License

This project is open source and available under the MIT License.

## 🤝 Contributing

Feel free to fork this repository and submit pull requests for any improvements!

## 📧 Contact

For questions or suggestions, please open an issue on GitHub.

---

**Happy task managing! 🚀**
