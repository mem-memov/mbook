package mbook

import "fmt"

type page struct {
	position uint
	height uint
	width uint
	size uint
	column []byte
}

func newPage(position uint, height uint, width uint) (*page, error) {
	size := height * width

	var column []byte
	if size > 0 {
		column = make([]byte, size)
	}

	return &page{
		position: position,
		height: height,
		width: width,
		size: size,
		column: column,
	}, nil
}

func (p *page) Write(position uint, value uint) error {
	if position < p.position {
		return fmt.Errorf("position %v outside of page", position)
	}

	offset := (position - p.position) * p.width + p.width - 1
	beforeShift := value
	for i := uint(0); i < p.width; i++ {
		afterShift := (beforeShift >> 8) << 8
		p.column[offset - i] = byte(beforeShift - afterShift)
		beforeShift = afterShift >> 8
	}
	return nil
}

func (p *page) Read(position uint) (uint, error) {
	offset := (position - p.position) * p.width
	value := uint(0)
	for i := uint(0); i < p.width; i++ {
		byteValue := uint(p.column[offset + i])
		value = (value << 8) + byteValue
	}
	return value, nil
}
