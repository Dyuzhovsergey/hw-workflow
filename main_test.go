package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld()
	expected := "Hello world"

	if result != expected {
		t.Errorf("Ожидалось: %s, получено: %s", expected, result)
	}
}
