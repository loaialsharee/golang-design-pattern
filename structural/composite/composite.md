# Composite

In simple terms: **Composite** is a structural design pattern that lets you **build tree-like structures** where **individual items** and **groups of items** are handled **uniformly**.

## 🤔 How to think of it?

You want to view a folder in your computer.

Once you open a folder, it will open all its contents.

## ⏱️ When do you need this pattern?

| Scenario | Example |
| -------- | ------- |
| 1\. You Have a **Tree-Like Structure** | If your data or system naturally forms a **hierarchy** (like a tree), where:<br>\* Some elements are **individual items**\* Some elements are **containers of items**<br>e.g. <br>\- File systems \(files and folders\)<br>\- Organization charts \(employees and managers\) |
| 2\. You Want to **Treat All Elements the Same** | You need to perform **the same operations** on:<br>\* A **single object** (`Leaf`)\* A **group of objects** (`Composite`)With Composite:<br>\* You don't care if it's a `File` or a `Folder`\* You just call `Open()`, `Render()`, `Print()`, etc. |
| 3\. You Want to **Build Structures Recursively** | If you want a structure where:<br>\* A `Folder` can contain `Files` or other `Folders`\* A `Manager` can manage `Employees` or other `Managers`This **recursive relationship** is ideal for Composite. |

**Rule of Thumb**: Use **Composite** when you want **uniform treatment of part-whole hierarchies** — like treating a **single object** and a **group of objects** the **same way**.

## 💻 Code Example in Go:

```
package main

import "fmt"

// ========================
// Component Interface
// ========================
type Employee interface {
	Display(indent string)
}

// ========================
// Leaf - Regular Employee
// ========================
type Developer struct {
	name string
	role string
}

func (d *Developer) Display(indent string) {
	fmt.Println(indent + "- Developer:", d.name, "("+d.role+")")
}

// ========================
// Composite - Manager
// ========================
type Manager struct {
	name      string
	role      string
	team      []Employee
}

func (m *Manager) Add(e Employee) {
	m.team = append(m.team, e)
}

func (m *Manager) Display(indent string) {
	fmt.Println(indent + "+ Manager:", m.name, "("+m.role+")")
	for _, member := range m.team {
		member.Display(indent + "  ")
	}
}

// ========================
// Main
// ========================
func main() {
	// Leaf nodes
	dev1 := &Developer{name: "Alice", role: "Frontend"}
	dev2 := &Developer{name: "Bob", role: "Backend"}
	dev3 := &Developer{name: "Charlie", role: "DevOps"}

	// Composite nodes
	manager1 := &Manager{name: "Diana", role: "Tech Lead"}
	manager2 := &Manager{name: "Evan", role: "Engineering Manager"}

	// Build team hierarchy
	manager1.Add(dev1)
	manager1.Add(dev2)

	manager2.Add(manager1)
	manager2.Add(dev3)

	// Display full org chart
	manager2.Display("")
}
```

- - -

## 🏃 How to run the codebase?

* Ensure your current directory is `~/structural/composite/code`
* Run `go run .`
