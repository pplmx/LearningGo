package tips

import "testing"

func TestVariadicFunc(t *testing.T) {
	type args struct {
		args []Option
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "TestVariadicFunc_empty", args: args{[]Option{}}, want: 0},
		{name: "TestVariadicFunc_1", args: args{[]Option{WithName("name")}}, want: 1},
		{name: "TestVariadicFunc_2", args: args{[]Option{WithName("name"), WithNum(1)}}, want: 2},
		{name: "TestVariadicFunc_3", args: args{[]Option{WithName("name"), WithNum(1), WithIsOk(true)}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VariadicFunc(tt.args.args...); got != tt.want {
				t.Errorf("VariadicFunc() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("TestVariadicFunc_4", func(t *testing.T) {
		if got := VariadicFunc(nil); got != 0 {
			t.Errorf("VariadicFunc() = %v, want %v", got, 0)
		}
	})
}
