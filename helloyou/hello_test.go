package main

import (
	"testing"
)

func TestHelo(t *testing.T) {
	got := Hello("Alfa")
	want := "Hello, Alfa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}