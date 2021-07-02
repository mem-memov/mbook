package mbook

import "testing"

func TestCalculateWidth(t *testing.T) {
	data := []struct {
		name     string
		position uint
		height   uint
		expected uint
	}{
		{"minimum width, first position", 0, 1, 1},
		{"minimum width, last position", 255, 1, 1},
		//{"width one, last position", 256, 1, 2},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := calculateWidth(d.position, d.height)
			if result != d.expected {
				t.Errorf("expected %d, got %d", d.expected, result)
			}
		})
	}
}
