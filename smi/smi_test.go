package smi

import "testing"

func TestPerimeter(t *testing.T) {

	rect := Rectangle{10, 10}

	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10, 10}, 100.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}

	}

}
