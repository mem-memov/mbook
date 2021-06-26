package mbook

import "testing"

func TestBook(t *testing.T) {
	data := []struct{
		name string
	} {
		{"min height"},
	}

	for _, d := range data {
		t.Run(d.name, func (t *testing.T) {

		})
	}
}