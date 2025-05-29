package main

type NovelBuilder struct {
	pages int
	name  string
	price float64
}

func newNovelBuilder() *NovelBuilder {
	return &NovelBuilder{}
}

func (b *NovelBuilder) setPages() {
	b.pages = 450
}

func (b *NovelBuilder) setName() {
	b.name = "The Alchemist"
}

func (b *NovelBuilder) setPrice() {
	b.price = 49.99
}

func (b *NovelBuilder) getBook() Book {
	return Book{
		pages: b.pages,
		name:  b.name,
		price: b.price,
	}
}
