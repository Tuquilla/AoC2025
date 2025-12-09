package main

import "testing"

func TestWithin(t *testing.T) {

	outerArea := Area{x1: 0, y1: 0, x2: 4, y2: 4, area: 25}
	innerArea := Area{x1: 1, y1: 1, x2: 3, y2: 3, area: 9}

	isWithin := within(&outerArea, &innerArea)
	if isWithin == false {
		t.Errorf("Is withing is false, should be true")
	}
}

func TestWithin2(t *testing.T) {

	outerArea := Area{x1: 0, y1: 0, x2: 4, y2: 4, area: 25}
	innerArea := Area{x1: 0, y1: 0, x2: 4, y2: 4, area: 25}

	isWithin := within(&outerArea, &innerArea)
	if isWithin == false {
		t.Errorf("Is withing is false, should be true")
	}
}

func TestWithin3(t *testing.T) {

	outerArea := Area{x1: 0, y1: 0, x2: 4, y2: 4, area: 25}
	innerArea := Area{x1: 0, y1: 0, x2: 4, y2: 5, area: 25}

	isWithin := within(&outerArea, &innerArea)
	if isWithin == true {
		t.Errorf("Is withing is true, should be false")
	}
}
