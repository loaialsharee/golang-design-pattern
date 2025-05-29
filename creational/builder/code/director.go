package main

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildBook() Book {
	d.builder.setPages()
	d.builder.setName()
	d.builder.setPrice()
	return d.builder.getBook()
}
