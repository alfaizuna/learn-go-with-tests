package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("test spanish greeting", func(t *testing.T) {
		got := Hello("Atun", "Spanish")
		want := "Hola, Atun"
		
		if(got != want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
} 