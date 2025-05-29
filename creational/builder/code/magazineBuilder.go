package main

type MagazineBuilder struct {
	pages int
	name  string
	price float64
}

func newMagazineBuilder() *MagazineBuilder {
	return &MagazineBuilder{}
}

func (b *MagazineBuilder) setPages() {
	b.pages = 25
}

func (b *MagazineBuilder) setName() {
	b.name = "Cosmopolitan"
}

func (b *MagazineBuilder) setPrice() {
	b.price = 18.99
}

func (b *MagazineBuilder) getBook() Book {
	return Book{
		pages: b.pages,
		name:  b.name,
		price: b.price,
	}
}
