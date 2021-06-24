package mbook

type page struct {
	position uint
	height uint
	width uint
	size uint
	column [][]byte
}

func newPage(position uint, height uint, width uint) (*page, error) {
	size := height * width
	return &page{
		position: position,
		height: height,
		width: width,
		size: size,
		column: make([][]byte, size),
	}, nil
}

func (p *page) Write(position uint, value uint) error {
	entry := make([]byte, p.width)
	beforeShift := value
	for i := uint(0); i < p.width; i++ {
		afterShift := (beforeShift >> 8) << 8
		entry[i] = byte(beforeShift - afterShift)
		beforeShift = afterShift >> 8
	}

	p.column[(position - p.position) * p.width] = entry
	return nil
}

func (p *page) Read(position uint) (uint, error) {
	return 0, nil
}
