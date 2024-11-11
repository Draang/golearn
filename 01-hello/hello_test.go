package main

import "testing"

// 1- se inicializa modulo => go mod init hello => crea un modulo go.mod => aqui se puede documentar
// 2- run >>=> go test
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
