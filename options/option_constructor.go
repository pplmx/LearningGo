package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Age   int
	Grade int
}

type UserOption func(*User)

func NewUser(opts ...UserOption) *User {
	u := &User{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}

func WithID(id int) UserOption {
	return func(u *User) {
		u.ID = id
	}
}

func WithName(name string) UserOption {
	return func(u *User) {
		u.Name = name
	}
}

func WithAge(age int) UserOption {
	return func(u *User) {
		u.Age = age
	}
}

func WithGrade(grade int) UserOption {
	return func(u *User) {
		u.Grade = grade
	}
}

func main() {
	u := NewUser(WithID(1), WithName("Tom"), WithAge(18), WithGrade(1))
	fmt.Printf("%+v", u)
}
