package structs

import "testing"

/**
table driven tests: Given a table of test cases, the actual test simply iterates through all table entries and for each entry performs the necessary tests. The test code is written once and amortized over all table entries, so it makes sense to write a careful test with good error messages.
*/
func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 40.0
		checkPerimeter(t, rectangle, want)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		want := 62.83185307179586476925286766559
		checkPerimeter(t, circle, want)
	})

}
func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{10.0}, 314.15926535897932384626433832795},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
