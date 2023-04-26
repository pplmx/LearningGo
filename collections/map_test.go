package collections

import (
	"reflect"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	type args[T any] struct {
		items []T
		mf    mapFunc[T]
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	t1 := &testCase[int]{
		name: "Double numbers",
		args: args[int]{
			items: []int{1, 2, 3, 4, 5},
		},
		want: []int{2, 4, 6, 8, 10},
	}
	t2 := &testCase[string]{
		name: "Upper strings",
		args: args[string]{
			items: []string{"a", "b", "c", "d", "e"},
		},
		want: []string{"A", "B", "C", "D", "E"},
	}
	t.Run(t1.name, func(t *testing.T) {
		if got := Map(t1.args.items, func(i int) int {
			return i * 2
		}); !reflect.DeepEqual(got, t1.want) {
			t.Errorf("Map() = %v, want %v", got, t1.want)
		}
	})
	t.Run(t2.name, func(t *testing.T) {
		if got := Map(t2.args.items, func(s string) string {
			return strings.ToUpper(s)
		}); !reflect.DeepEqual(got, t2.want) {
			t.Errorf("Map() = %v, want %v", got, t2.want)
		}
	})
}
