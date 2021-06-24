package mbook

type Book struct {
	height uint
	pages []page
}

func NewStorage(height uint) (*Book, error) {
	return &Book {
		height: height,
		pages: make([]page, 0),
	}, nil
}

func (b *Book) Write(position uint, value uint) error {
	return nil
}

func (b *Book) Read(position uint) (uint, error) {
	return 0, nil
}
