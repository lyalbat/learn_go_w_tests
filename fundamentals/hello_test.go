package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to a specific person when sent a name", func(t *testing.T){
		got := CustomHello("larissa")
		want := "hello larissa"
	
		assertIncorrectMessage(t,got,want)
	})
	t.Run("say hello world when an empty string is sent", func(t *testing.T){
		got := CustomHello("")
		want := "hello world"
	
		assertIncorrectMessage(t,got,want)
	})
}

func assertIncorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}