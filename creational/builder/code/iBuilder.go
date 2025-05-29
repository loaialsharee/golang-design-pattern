package main

type IBuilder interface {
	setPages()
	setName()
	setPrice()
	getBook() Book
}

func getBuilder(builderType string) IBuilder {
	if builderType == "magazine" {
		return newMagazineBuilder()
	}

	if builderType == "novel" {
		return newNovelBuilder()
	}
	return nil
}
