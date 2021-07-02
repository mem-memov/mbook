package mbook

import (
	"testing"
)

func TestBook(t *testing.T) {
	data := []struct {
		name   string
		height uint
	}{
		{"min height", 1},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			b, _ := NewBook(d.height)
			b.Write(0, 1)
			result := b.Read(0)
			if result != 1 {
				t.Errorf("expected %v, got %v", 1, result)
			}
		})
	}
}
