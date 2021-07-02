package mbook

func calculateWidth(position uint, height uint) uint {
	lastPosition := position + height
	i := uint(0)
	for ; lastPosition > 0; i++ {
		lastPosition >>= 8
	}
	return i
}
