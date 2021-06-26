package mbook

import "fmt"

type Book struct {
	height uint
	pages []page
}

func NewBook(height uint) (*Book, error) {
	if height < 1 {
		return nil, fmt.Errorf("zero book height not allowed")
	}

	book := &Book {
		height: height,
		pages: make([]page, 0),
	}

	book.addPage(0)

	return book, nil
}

func (b *Book) Write(position uint, value uint) {
	pageNumber := b.appendPages(position)
	b.pages[pageNumber].Write(position, value)
}

func (b *Book) Read(position uint) uint {
	pageNumber := b.getPageNumber(position)

	if pageNumber > b.countPages() {
		return 0
	}

	return b.pages[pageNumber].Read(position)
}

func (b *Book) appendPages(position uint) uint {
	pageNumber := b.getPageNumber(position)

	if pageNumber > b.countPages() {
		for i := b.countPages() + 1; i < pageNumber; i++ {
			b.addPage(i * b.height)
		}
	}

	return pageNumber
}

func (b *Book) addPage(position uint) {
	page := newPage(position, b.height)
	b.pages = append(b.pages, *page)
}

func  (b *Book) countPages() uint {
	return uint(len(b.pages))
}

func (b *Book) getPageNumber(position uint) uint {
	pageNumber := position / b.height
	rest := position % b.height
	if rest > 0 {
		pageNumber++
	}
	return pageNumber
}
