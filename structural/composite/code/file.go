package main

import "fmt"

type File struct {
	name string
}

func (f *File) scan() {
	fmt.Printf("Scanning for malwares in file %s\n", f.name)
}

func (f *File) getName() string {
	return f.name
}
