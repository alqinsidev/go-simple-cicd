package main

import "fmt"

type Person struct {
	Name        string
	Nationality string
}

func main() {

	person := Person{
		Name:        "Padlan Alqinsi",
		Nationality: "Indonesia",
	}

	fmt.Println(greetPerson(person))
}

func greetPerson(person Person) string {
	return fmt.Sprintf("Hello %s from %s!", person.Name, person.Nationality)
}
