package main

import "testing"

import "./basicgrammer/packages"

func TestHello(t *testing.T) {

	t.Run("Simple Test", func(t *testing.T) {
		got := Hello()
		want := "Hello, 世界"
		assertCorrectMessage(t, got, want)
	})
}

func TestSandbox(t *testing.T) {

	t.Run("Simple Test", func(t *testing.T) {
		Sandbox()
		//Packages()
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
