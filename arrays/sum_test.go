package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Coleccion de 5 numeros", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := Sum(nums)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	})
	t.Run("Coleccion de 2 numeros", func(t *testing.T) {
		nums := []int{1, 0}
		got := Sum(nums)
		want := 1
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	})

}
func TestSumAll(t *testing.T) {
	listNums := [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	got := SumAll(listNums)
	want := []int{6, 6, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given, %v", got, want, listNums)
	}
}
