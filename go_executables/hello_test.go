package main

import "testing"

func TestHello(t *testing.T) {
	result := hello("Akhand")
	if result != "hello Akhand" {
		t.Errorf("hello(\"Akhand\") failed, expected %v , got %v", "hello Akhand", result)
	}
}
