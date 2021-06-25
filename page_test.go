package mbook

import "testing"

func TestPage(t *testing.T) {
	data := []struct{
		name string
		position uint
		height uint
		width uint
	}{
		{"empty page", 0, 0, 0},
		{"one byte page one column", 0, 1, 1},
		{"one byte page one column shifted", 10, 1, 1},
		{"one byte page max columns", 0, 256, 1},
		{"one byte page max columns shifted", 10, 256-10, 1},
		{"two byte page one column", 0, 1, 2},
		{"two byte page one column shifted", 10, 1, 2},
		{"two byte page max columns", 0, 65536, 2},
		{"two byte page max columns shifted", 10, 65536-10, 2},
	}

	for _, d := range data {
		t.Run(d.name, func (t *testing.T) {
			p, _ := newPage(d.position, d.height, d.width)

			fillPage(p, 1, t)
			writeOneValue(p, 1, t)

			maxValue := 2 ^ (d.width * 8)
			fillPage(p, maxValue, t)
			writeOneValue(p, maxValue, t)
		})
	}
}

func writeOneValue(p *page, value uint, t *testing.T) {
	if p.height > 20 {
		writeOneInHead(p, value, 10, t)
		writeOneInTail(p, value, 10, t)
	} else {
		writeOneInHead(p, value, p.height, t)
	}
}

func fillPage(p *page, value uint, t *testing.T) {
	if p.height > 20 {
		fillHead(p, value, 10, t)
		fillTail(p, value, 10, t)
	} else {
		fillHead(p, value, p.height, t)
	}
}

func writeOneInHead(p *page, value uint, span uint, t *testing.T) {
	for h := uint(0); h < span; h++ {
		fillHead(p, 0, span, t)
		checkModification(p, p.position + h, value, t)
	}
}

func fillHead(p *page, value uint, span uint, t *testing.T) {
	for h := uint(0); h < span; h++ {
		checkModification(p, p.position + h, value, t)
	}
}

func writeOneInTail(p *page, value uint, span uint, t *testing.T) {
	for h := p.height - span; h < p.height; h++ {
		fillTail(p, 0, span, t)
		checkModification(p, p.position + h, value, t)
	}
}

func fillTail(p *page, value uint, span uint, t *testing.T) {
	for h := p.height - span; h < p.height; h++ {
		checkModification(p, p.position + h, value, t)
	}
}

func checkModification(p *page, position uint, value uint, t *testing.T) {
	_ = p.Write(position, value)
	result, _ := p.Read(position)
	if result != value {
		t.Errorf("Expected %v, got %v", value, result)
	}
}
