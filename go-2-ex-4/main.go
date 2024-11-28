package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}

	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Name     string
		Students []Student
	}

	// TODO: declare a map of modules being attended by multiple classes
	modules := map[int][]Class{
		346: {
			{Name: "Class A", Students: []Student{
				{FirstName: "Alice", LastName: "Smith"},
				{FirstName: "Bob", LastName: "Johnson"},
				{FirstName: "Charlie", LastName: "Brown"},
			}},
			{Name: "Class B", Students: []Student{
				{FirstName: "David", LastName: "Wilson"},
				{FirstName: "Eve", LastName: "Taylor"},
				{FirstName: "Frank", LastName: "Anderson"},
			}},
		},
		104: {
			{Name: "Class A", Students: []Student{
				{FirstName: "Alice", LastName: "Smith"},
				{FirstName: "Bob", LastName: "Johnson"},
				{FirstName: "Charlie", LastName: "Brown"},
			}},
		},
		117: {
			{Name: "Class B", Students: []Student{
				{FirstName: "David", LastName: "Wilson"},
				{FirstName: "Eve", LastName: "Taylor"},
				{FirstName: "Frank", LastName: "Anderson"},
			}},
		},
	}

	// TODO: output everything using fmt.Println()
	for moduleID, classes := range modules {
		fmt.Printf("Module %d:\n\n", moduleID)
		for _, class := range classes {
			fmt.Printf("  Class %s:\n\n", class.Name)
			for _, student := range class.Students {
				fmt.Printf("    - %s %s\n", student.FirstName, student.LastName)
			}
		}
	}
}
