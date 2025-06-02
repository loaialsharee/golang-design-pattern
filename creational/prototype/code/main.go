package main

import "fmt"

func main() {
	pasta := &Ingredient{name: "Pasta"}
	tomatoSauce := &Ingredient{name: "TomatoSauce"}
	meatballs := &Ingredient{name: "Meatballs"}

	meal1 := &Meal{
		children: []Inode{pasta},
		name:     "Meal1",
	}

	meal2 := &Meal{
		children: []Inode{meal1, tomatoSauce, meatballs},
		name:     "Meal2",
	}
	fmt.Println("\nPrinting Ingredients for Meal2")
	meal2.print("  ")

	cloneMeal := meal2.clone()
	fmt.Println("\nPrinting Ingredients for cloned Meals")
	cloneMeal.print("  ")
}
