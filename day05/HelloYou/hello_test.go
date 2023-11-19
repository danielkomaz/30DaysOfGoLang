package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Daniel", "")
		want := "Hello, Daniel"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Daniel", "Spanish")
		want := "Hola, Daniel"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Daniel", "French")
		want := "Bonjour, Daniel"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Austrian", func(t *testing.T) {
		got := Hello("Daniel", "Austrian")
		want := "Servus, Daniel"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Prefix Spanish", func(t *testing.T) {
		got := greetingPrefix("Spanish")
		want := "Hola, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("Prefix French", func(t *testing.T) {
		got := greetingPrefix("French")
		want := "Bonjour, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("Prefix Austrian", func(t *testing.T) {
		got := greetingPrefix("Austrian")
		want := "Servus, "
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
