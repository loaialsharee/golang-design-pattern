package main

import "fmt"

type Ingredient struct {
	name string
}

func (f *Ingredient) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *Ingredient) clone() Inode {
	return &Ingredient{name: f.name + "_clone"}
}
