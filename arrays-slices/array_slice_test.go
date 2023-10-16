package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of slice", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 4, 5})
		expected := 15

		if got != expected {
			t.Errorf("expeced '%d' we got '%d'", expected, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checksums := func(t testing.TB, got, expected []int) {
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expeced %v want %v", expected, got)
		}
	}

	t.Run("make sum of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checksums(t, expected, got)
	})

	t.Run("test with empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}

		checksums(t, expected, got)
	})
}
