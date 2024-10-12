package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("apple")
		want := ErrNotFound
		if err == nil {
			t.Fatal(ErrNotFound)
		}
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, definition)
	})
	t.Run("add a word already defined", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new definition")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update word", func(t *testing.T) {
		word := "test"
		previousDefinition := "this is just a test"
		dictionary := Dictionary{word: previousDefinition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		if err != nil {
			t.Fatal("should find word:", err)
		}
		got, _ := dictionary.Search("test")
		assertStrings(t, got, newDefinition)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, definition string) {
	t.Helper()

	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}
