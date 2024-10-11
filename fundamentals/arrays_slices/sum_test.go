package sum

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("should do the correct sum of a slice", func (t *testing.T){
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
	
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("Should get two slices and sum both (individually)", func(t *testing.T){
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		//reflect.DeepEqual checks if any two variables are the same
		//note: DeepEqual is NOT type safe, so it can be used to do wrong operations :
		//like comparing strings with slices or integers with strings. Important to be careful
		if !reflect.DeepEqual(got,want){
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Should return the sum of all tails (individually)", func(t *testing.T){
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 1})
		want := []int{5, 10}

		if !reflect.DeepEqual(got,want){
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}