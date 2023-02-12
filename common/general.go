package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func sum(num int) int {
	var result int
	for i := 0; i < num; i++ {
		result += i
	}
	return result
}

type Gender int
type Grade int

const (
	Male Gender = iota
	Female
)

const (
	One = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Eleven
	Twelve
)

type Person struct {
	Name   string
	Gender Gender
	Age    uint
}

type Student struct {
	Person
	Scores []int
	Grade  Grade
}

func (p Person) printPersonInfo() {
	fmt.Printf("Person: name=%s, gender=%d, age=%d\n", p.Name, p.Gender, p.Age)
}

func (s Student) printPersonInfo() {
	fmt.Printf("Student: name=%s, scores=%v, grade=%d\n", s.Name, s.Scores, s.Grade)
}

func main() {
	r1 := add(6, 7)
	r2 := subtract(6, 7)
	r3 := sum(10)
	fmt.Printf("func add: %d\n", r1)
	fmt.Printf("func sub: %d\n", r2)
	fmt.Printf("func sum: %d\n", r3)
	p1 := Person{
		Name:   "John",
		Gender: Male,
		Age:    30,
	}
	p2 := Person{
		Name:   "Jane",
		Gender: Female,
		Age:    40,
	}
	s1 := Student{
		Person: Person{
			Name:   "Tommy",
			Gender: Male,
			Age:    12,
		},
		Scores: []int{90, 90, 100},
		Grade:  Four,
	}
	p1.printPersonInfo()
	p2.printPersonInfo()
	s1.printPersonInfo()
}
