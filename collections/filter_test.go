package collections

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args[T any] struct {
		items []T
		ff    filterFunc[T]
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}

	t1 := &testCase[int]{
		name: "Filter even numbers",
		args: args[int]{
			items: []int{1, 2, 3, 4, 5},
		},
		want: []int{2, 4},
	}

	t2 := &testCase[string]{
		name: "Filter strings with length > 3",
		args: args[string]{
			items: []string{"a", "ab", "abc", "abcd", "abcde"},
		},
		want: []string{"abcd", "abcde"},
	}

	t.Run(t1.name, func(t *testing.T) {
		if got := Filter(t1.args.items, func(i int) bool {
			return i%2 == 0
		}); !reflect.DeepEqual(got, t1.want) {
			t.Errorf("Filter() = %v, want %v", got, t1.want)
		}
	})

	t.Run(t2.name, func(t *testing.T) {
		if got := Filter(t2.args.items, func(s string) bool {
			return len(s) > 3
		}); !reflect.DeepEqual(got, t2.want) {
			t.Errorf("Filter() = %v, want %v", got, t2.want)
		}
	})

}
