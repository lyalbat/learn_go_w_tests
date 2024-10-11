package iteration

import "testing"

func TestIteration(t *testing.T){
	scream := ScreamGenerator("a",5)
	expectedScream := "aaaaa"
	if scream != expectedScream {
		t.Errorf("expected %q but got %q", expectedScream, scream)
	}
}

//to run use:  go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ScreamGenerator("a",5)
	}
}

