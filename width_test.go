package mbook

import "testing"

func TestCalculateWidth(t *testing.T) {
	data := []struct {
		name     string
		position uint
		height   uint
		expected uint
	}{
		{"one line, one byte width, min value", 0, 1, 1},
		{"one line, one byte width, max value", 254, 1, 1},
		{"one line, two byte width, min value", 255, 1, 2},
		{"one line, two byte width, max value", 65534, 1, 2},
		{"one line, three byte width, min value", 65535, 1, 3},
		{"two lines, one byte width, min value", 0, 2, 1},
		{"two lines, one byte width, max value", 253, 2, 1},
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
