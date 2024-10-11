package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	//.2f two decimal points
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	//table driven tests - ref: go.dev/wiki/TableDrivenTests
	//anonymous struct
	//to run a specific test: go test -run TestArea/Rectangle
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, test_unit := range areaTests {
		got := test_unit.shape.Area()
		// the #v will return the struct parameters alongside its values - makes debugging easier
		if got != test_unit.want {
			t.Errorf("%#v got %g want %g", test_unit.shape, got, test_unit.want)
		}
	}
}
