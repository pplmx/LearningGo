package collections

import (
	"reflect"
	"sort"
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
			mf:    func(i int) int { return i * 2 },
		},
		want: []int{2, 4, 6, 8, 10},
	}
	t2 := &testCase[string]{
		name: "Upper strings",
		args: args[string]{
			items: []string{"a", "b", "c", "d", "e"},
			mf:    func(s string) string { return strings.ToUpper(s) },
		},
		want: []string{"A", "B", "C", "D", "E"},
	}
	t.Run(t1.name, func(t *testing.T) {
		if got := Map(t1.args.items, t1.args.mf); !reflect.DeepEqual(got, t1.want) {
			t.Errorf("Map() = %v, want %v", got, t1.want)
		}

		// Because the t1.args.items has been modified by Map function, so we need to reset it.
		t1.args.items = []int{1, 2, 3, 4, 5}
		got := ConcurrentMap(t1.args.items, t1.args.mf)
		sort.Ints(got)
		sort.Ints(t1.want)
		if !reflect.DeepEqual(got, t1.want) {
			t.Errorf("ConcurrentMap() = %v, want %v", got, t1.want)
		}
	})
	t.Run(t2.name, func(t *testing.T) {
		if got := Map(t2.args.items, t2.args.mf); !reflect.DeepEqual(got, t2.want) {
			t.Errorf("Map() = %v, want %v", got, t2.want)
		}

		// reset t2.args.items
		t2.args.items = []string{"a", "b", "c", "d", "e"}
		got := ConcurrentMap(t2.args.items, t2.args.mf)
		sort.Strings(got)
		sort.Strings(t2.want)
		if !reflect.DeepEqual(got, t2.want) {
			t.Errorf("ConcurrentMap() = %v, want %v", got, t2.want)
		}
	})
}

// BenchmarkMap benchmarks the Map function.
func BenchmarkMap(b *testing.B) {
	items := make([]int, 1000-0000-0000) // Adjust the size based on your needs
	mf := func(x int) int {
		return x * 2
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Map(items, mf)
	}
}

// BenchmarkConcurrentMap benchmarks the ConcurrentMap function.
// FIXME: ConcurrentMap is slower than Map, which is not expected.
func BenchmarkConcurrentMap(b *testing.B) {
	items := make([]int, 1000-0000-0000) // Adjust the size based on your needs
	mf := func(x int) int {
		return x * 2
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = ConcurrentMap(items, mf)
	}
}
