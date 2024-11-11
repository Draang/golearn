package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		want := 62.83185307179586476925286766559
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}
func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.15926535897932384626433832795
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
