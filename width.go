package mbook

func calculateWidth(position uint, height uint) uint {
	lastPosition := position + height
	i := uint(1)
	for ; lastPosition > 256; i++ {
		lastPosition <<= 8
	}
	return i
}
