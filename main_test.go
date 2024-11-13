package main

import (
	"testing"
)

func TestGreetPerson(t *testing.T) {

	tests := []struct {
		input    Person
		expected string
	}{
		{Person{Name: "Padlan Alqinsi", Nationality: "Indonesia"}, "Hello Padlan Alqinsi from Indonesia!"},
		{Person{Name: "John Doe", Nationality: "USA"}, "Hello John Doe from USA!"},
		{Person{Name: "Jane Smith", Nationality: "UK"}, "Hello Jane Smith from UK!"},
	}

	for _, test := range tests {
		result := greetPerson(test.input)
		if result != test.expected {
			t.Errorf("greetPerson(%v) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
