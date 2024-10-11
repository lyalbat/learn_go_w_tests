package main

import "testing"


func TestAddition(t *testing.T) {
	t.Run("equal 4 when adding 2 + 2", func(t *testing.T){
		sum := Add(2,2)
		expected := 4
	
		assertIncorrectResult(t,sum,expected)
	})
}

func assertIncorrectResult(t testing.TB, sum, expected int){
	t.Helper()
	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}