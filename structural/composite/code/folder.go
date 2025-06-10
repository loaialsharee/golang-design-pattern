package main

import "fmt"

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) scan() {
	fmt.Printf("Scanning for malwares in folder %s\n", f.name)
	for _, composite := range f.components {
		composite.scan()
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}
