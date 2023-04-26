package collections

// filterFunc is a function that takes an item of type T and returns a bool.
type filterFunc[T any] func(T) bool

// Filter is a function that takes a slice of items and a filterFunc and returns a slice of items.
func Filter[T any](items []T, ff filterFunc[T]) []T {
	filteredItems := make([]T, 0)
	for _, item := range items {
		if ff(item) {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}
