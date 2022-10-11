package di

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}


	Greet(&buffer, "Rodrigo")

  fmt.Printf("address: %v", &buffer)
	got := buffer.String()
	want := "Rodrigo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
