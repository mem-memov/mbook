package mbook

type Keeper interface {
	Write(position uint, value uint) error
	Read(position uint) (uint, error)
}
