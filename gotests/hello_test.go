package gotests

import "testing"

func TestHelloEmptyArg(t *testing.T) {
	emptyResult := Hello("")

	if emptyResult != "Hello Dude" {
		t.Errorf("Hello(\"\")  failed , expected: %v, got: % v", "Hello Dude", emptyResult)
	}
}

func TestHelloValidArg(t *testing.T) {
	result := Hello("Akhand")
	if result != "Hello Akhand" {
		t.Errorf("Hello (\"Akhand\") failed, expected: %v , got: %v", "Hello Akhand", result)
	}
}
