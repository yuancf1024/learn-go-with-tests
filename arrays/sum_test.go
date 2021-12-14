package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	// t.Run("collection of 5 numbers", func (t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}

	// 	got := Sum(numbers)
	// 	want := 15

	// 	if got != want {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

// func TestSumAll(t *testing.T) {

// 	got := sumAll([]int{1, 2}, []int{0, 9})
// 	want := []int{3, 9}
// 	// want := "bob"

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }

func TestSumAllTails(t *testing.T) {

	// 我们的测试代码有一部分是重复的，我们可以把它放到另一个函数中复用。
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
		// checkSums(t, got, "dave")
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		// want := "bob"
		checkSums(t, got, want)
	})
}