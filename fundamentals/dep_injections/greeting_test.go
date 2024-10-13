package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	//buffer type implements Writer interface - using in fmt under the hood
	//because it has the method Write(p []byte) (n int, err error)
	buffer := bytes.Buffer{}
	Greet(&buffer, "Larissa")
	got := buffer.String()
	want := "Hello, Larissa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
