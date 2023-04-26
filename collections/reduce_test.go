package collections

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	type args[T any] struct {
		items []T
		rf    reduceFunc[T]
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	t1 := &testCase[int]{
		name: "Sum numbers",
		args: args[int]{
			items: []int{1, 2, 3, 4, 5},
		},
		want: 15,
	}

	t2 := &testCase[string]{
		name: "Concatenate strings",
		args: args[string]{
			items: []string{"h", "e", "l", "l", "o"},
		},
		want: "hello",
	}

	// generate an empty list
	t3 := &testCase[int]{
		name: "Sum numbers",
		args: args[int]{
			items: []int{},
		},
		want: 0,
	}

	t.Run(t1.name, func(t *testing.T) {
		if got := Reduce(t1.args.items, func(a, b int) int {
			return a + b
		}); !reflect.DeepEqual(got, t1.want) {
			t.Errorf("Reduce() = %v, want %v", got, t1.want)
		}
	})
	t.Run(t2.name, func(t *testing.T) {
		if got := Reduce(t2.args.items, func(a, b string) string {
			return a + b
		}); !reflect.DeepEqual(got, t2.want) {
			t.Errorf("Reduce() = %v, want %v", got, t2.want)
		}
	})
	t.Run(t3.name, func(t *testing.T) {
		if got := Reduce(t3.args.items, func(a, b int) int {
			return a + b
		}); !reflect.DeepEqual(got, t3.want) {
			t.Errorf("Reduce() = %v, want %v", got, t3.want)
		}
	})
}
