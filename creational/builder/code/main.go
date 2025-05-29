package main

import "fmt"

func main() {
	magazineBuilder := getBuilder("magazine")
	novelBuilder := getBuilder("novel")

	director := newDirector(magazineBuilder)
	magazine := director.buildBook()

	fmt.Printf("Magazine Pages: %d\n", magazine.pages)
	fmt.Printf("Magazine Name: %s\n", magazine.name)
	fmt.Printf("Magazine Price: %.2f\n", magazine.price)

	director.setBuilder(novelBuilder)
	novel := director.buildBook()

	fmt.Printf("\nNovel Pages: %d\n", novel.pages)
	fmt.Printf("Novel Name: %s\n", novel.name)
	fmt.Printf("Novel Price: %.2f\n", novel.price)
}
