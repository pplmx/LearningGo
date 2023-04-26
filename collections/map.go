package collections

// mapFunc is a function that takes an item of type T and returns an item of type T.
type mapFunc[T any] func(T) T

// Map is a function that takes a slice of items and a mapFunc and returns a slice of items.
func Map[T any](items []T, mf mapFunc[T]) []T {
	for i := range items {
		items[i] = mf(items[i])
	}
	return items
}
