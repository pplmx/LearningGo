package main

import "fmt"

type Obj[T any] struct {
	Value T
}

type Option[T any] func(*Obj[T])

func WithValue[T any](value T) Option[T] {
	return func(c *Obj[T]) {
		c.Value = value
	}
}

func NewObj[T any](options ...Option[T]) *Obj[T] {
	c := &Obj[T]{}
	for _, opt := range options {
		opt(c)
	}
	return c
}

func main() {
	c1 := NewObj(WithValue(42), WithValue(1))
	fmt.Println(c1.Value)
}
