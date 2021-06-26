package mbook

type page struct {
	position uint
	height uint
	width uint
	size uint
	column []byte
}

func newPage(position uint, height uint) *page {
	width := calculateWidth(position, height)
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
	}
}

func (p *page) Write(position uint, value uint) {
	offset := (position - p.position) * p.width + p.width - 1
	beforeShift := value
	for i := uint(0); i < p.width; i++ {
		afterShift := (beforeShift >> 8) << 8
		p.column[offset - i] = byte(beforeShift - afterShift)
		beforeShift = afterShift >> 8
	}
}

func (p *page) Read(position uint) uint {
	offset := (position - p.position) * p.width
	value := uint(0)
	for i := uint(0); i < p.width; i++ {
		byteValue := uint(p.column[offset + i])
		value = (value << 8) + byteValue
	}
	return value
}

func calculateWidth(position uint, height uint) uint {
	lastPosition := position + height
	i := uint(1)
	for ; lastPosition > 256; i++ {
		lastPosition <<= 8
	}
	return i
}
