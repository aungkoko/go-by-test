package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Aung Ko")
	want := "Hello, Aung Ko"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
