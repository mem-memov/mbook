package mbook

// page contains a column of lines of same width
type page struct {
	position uint   // position of the first line in the whole book
	height   uint   // height of page in
	width    uint   // page line size in bytes
	size     uint   // page size in bytes
	column   []byte // page column with lines of fixed size
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
		height:   height,
		width:    width,
		size:     size,
		column:   column,
	}
}

func (p *page) Write(position uint, value uint) {
	offset := (position - p.position) * p.width + p.width - 1 // last byte goes first, as it contains the little end
	beforeShift := value
	for i := uint(0); i < p.width; i++ {
		if beforeShift == 0 {
			p.column[offset-i] = byte(0)
			continue
		}
		afterShift := (beforeShift >> 8) << 8
		p.column[offset-i] = byte(beforeShift - afterShift)
		beforeShift = afterShift >> 8
	}
}

func (p *page) Read(position uint) uint {
	offset := (position - p.position) * p.width
	value := uint(0)
	for i := uint(0); i < p.width; i++ {
		byteValue := uint(p.column[offset+i])
		value = (value << 8) + byteValue
	}
	return value
}
