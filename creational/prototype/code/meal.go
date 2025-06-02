package main

import "fmt"

type Meal struct {
	children []Inode
	name     string
}

func (f *Meal) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Meal) clone() Inode {
	cloneMeal := &Meal{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneMeal.children = tempChildren
	return cloneMeal
}
