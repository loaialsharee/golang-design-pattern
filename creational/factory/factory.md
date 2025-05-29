# Abstract Factory

In simple terms: this design pattern helps you create "families" of related objects without specifying their exact classes.

## 🤔 How to think of it?

You go to visit a furniture store and you ask that you want a set of modern style (a chair, a sofa, and a table).

The store now will give you:

- a "modern" chair
- a "modern" sofa
- a "modern" table

The morale story is that you don't pick each item by yourself, you just say _the style_, and the store will give you the matching items!

## ⏱️ When do you need this pattern?

| Scenario                                         | Example                                                                                                                                                                                                                       |
| ------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1. Need to createfamilies of related objects     | building a UI library which supports light & dark theme which helps you switch themes without mixing styles.                                                                                                                  |
| 2. Want to switch product families at runtime    | Your app supports multiple database engines (MySQL, PostgreSQL, MongoDB) and each engine has a connection class, query class, and a result parser. This pattern will help to plug in the whole set for the selected database. |
| 3. Want to enforce consistency across components | In a game, you want to have different "worlds" and each workd has a character, a weapon, and an enemy. This shall ensure all components belong to the same world style.                                                       |

**Rule of Thumb**: Use this design pattern when you need a set of related things, and you want to be able to "swap" the entire set easily.

## 💻 Code Example in Go:

```
package main

import "fmt"

// --- Product Interfaces ---

type Chair interface {
	Sit()
}

type Sofa interface {
	Relax()
}

// --- Modern Products ---

type ModernChair struct{}
func (m ModernChair) Sit() {
	fmt.Println("Sitting on a modern chair")
}

type ModernSofa struct{}
func (m ModernSofa) Relax() {
	fmt.Println("Relaxing on a modern sofa")
}

// --- Victorian Products ---

type VictorianChair struct{}
func (v VictorianChair) Sit() {
	fmt.Println("Sitting on a Victorian chair")
}

type VictorianSofa struct{}
func (v VictorianSofa) Relax() {
	fmt.Println("Relaxing on a Victorian sofa")
}

// --- Abstract Factory Interface ---

type FurnitureFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
}

// --- Modern Factory ---

type ModernFactory struct{}
func (m ModernFactory) CreateChair() Chair {
	return ModernChair{}
}
func (m ModernFactory) CreateSofa() Sofa {
	return ModernSofa{}
}

// --- Victorian Factory ---

type VictorianFactory struct{}
func (v VictorianFactory) CreateChair() Chair {
	return VictorianChair{}
}
func (v VictorianFactory) CreateSofa() Sofa {
	return VictorianSofa{}
}

// --- Client Code ---

func FurnishRoom(factory FurnitureFactory) {
	chair := factory.CreateChair()
	sofa := factory.CreateSofa()

	chair.Sit()
	sofa.Relax()
}

// --- Main ---

func main() {
	fmt.Println("Modern Room:")
	FurnishRoom(ModernFactory{})

	fmt.Println("\nVictorian Room:")
	FurnishRoom(VictorianFactory{})
}
```
