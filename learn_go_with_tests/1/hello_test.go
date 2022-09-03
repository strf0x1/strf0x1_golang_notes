package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("you sorry sack of meat and bones")
		want := "Hello, you sorry sack of meat and bones"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

// passing in t testing.TB is a good idea bc it's an interface for both t *testing.T and *testing.B
// t.Helper() tells the test suite that this method is a helper. helps devs track down problems easier
func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
