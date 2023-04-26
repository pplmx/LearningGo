package collections

// reduceFunc is a function that takes two items of the same type and returns a single item of that type.
type reduceFunc[T any] func(T, T) T

// Reduce is a function that takes a slice of items and a reduceFunc and returns a single item.
func Reduce[T any](items []T, rf reduceFunc[T]) T {
	if len(items) == 0 {
		return *new(T)
	}
	result := items[0]
	for i := 1; i < len(items); i++ {
		result = rf(result, items[i])
	}
	return result
}
