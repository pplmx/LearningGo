package tests

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test_add1", args: args{1, 2}, want: 3},
		{name: "Test_add2", args: args{-1, 2}, want: 1},
		{name: "Test_add3", args: args{-1, -2}, want: -3},
		{name: "Test_add4", args: args{1999, 2001}, want: 4000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divide(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{name: "Test_divide1", args: args{1, 2}, want: 0.5},
		{name: "Test_divide2", args: args{-1, -2}, want: 0.5},
		{name: "Test_divide_expect_error", args: args{-1, 0}, want: "ZeroDivisionError", wantErr: true},
		//{name: "Test_divide_failed_case", args: args{-1, -2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.args.a, tt.args.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("divide() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err.Error() != "ZeroDivisionError" {
					t.Errorf("divide() got = %v, want %v", err.Error(), tt.want)
					return
				}
				return
			}
			if got != tt.want {
				t.Errorf("divide() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiply(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test_multiply1", args: args{1, 2}, want: 2},
		{name: "Test_multiply2", args: args{-1, 2}, want: -2},
		{name: "Test_multiply3", args: args{-1, -2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtract(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test_subtract1", args: args{1, 2}, want: -1},
		{name: "Test_subtract2", args: args{-1, 2}, want: -3},
		{name: "Test_subtract3", args: args{-1, -2}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
