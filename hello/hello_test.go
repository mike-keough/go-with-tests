package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Mike", "")
		want := "Hello, Mike"

		assertMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Mike", "Spanish")
		want := "Hola, Mike"
		assertMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Mike", "French")
		want := "Bonjour, Mike"

		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
