# Prototype

In simple terms: this design pattern makes a copy of an object without knowing exactly what it is.

## 🤔 How to think of it?

You have a game or a toy!

You just tell that game/toy, "Clone yourself!"

That object knows how to clone itself, and gives you an exact copy

## ⏱️ When do you need this pattern?

| Scenario | Example |
| -------- | ------- |
| 1\. When Object Creation Is Expensive | Creating a game enemy with AI, animation, and stats — better to clone a base model than recreate all that. |
| 2\. When You Need Many Similar Objects | A form template that you fill in multiple times — just copy the template each time. |
| 3\. When You Want to Avoid Subclassing or Complex `new` Logic | Sometimes you’d need lots of `if/else` or subclassing just to decide how to make something.<br>With Prototype, each object **knows how to copy itself**, so no big `switch` logic or factory needed. |
| 4\. When You Need Deep Copies | If your object has **nested objects** inside (like folders with files), Prototype is great for deep cloning the whole structure. |
| 5\. When System Needs to Be Flexible and Extensible | If you're designing a plugin system or something where new object types will be added later, Prototype lets each type handle its own cloning — no need to change existing code. |

**Rule of Thumb**: Use the Prototype pattern when creating an object from scratch is expensive, repetitive, or complex — and you already have a similar object you can copy.

## 💻 Code Example in Go:

```
package main

import (
	"fmt"
)

// Prototype interface
type UIElement interface {
	Clone() UIElement
	Render()
}

// Concrete prototype: Button
type Button struct {
	Label  string
	Color  string
	Action string
}

func (b *Button) Clone() UIElement {
	// Return a new copy
	return &Button{
		Label:  b.Label + " (Copy)",
		Color:  b.Color,
		Action: b.Action,
	}
}

func (b *Button) Render() {
	fmt.Printf("Rendering Button: [%s] with color %s (action: %s)\n", b.Label, b.Color, b.Action)
}

func main() {
	// Original button
	saveButton := &Button{
		Label:  "Save",
		Color:  "Green",
		Action: "SaveDocument",
	}

	// Render original
	saveButton.Render()

	// Clone it
	clonedButton := saveButton.Clone()
	clonedButton.Render()

	// Clone again
	anotherClone := clonedButton.Clone()
	anotherClone.Render()
}
```

- - -

## 🏃 How to run the codebase?

* Ensure your current directory is `~/creational/prototype/code`
* Run `go run .`
