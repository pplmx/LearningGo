package tests

import "testing"

func Test_Divide(t *testing.T) {
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
			got, err := Divide(tt.args.a, tt.args.b)
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

func Test_Multiply(t *testing.T) {
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
			if got := Multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Multiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(32232, 21214)
	}
}

func Benchmark_Divide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Divide(121345, 532312)
		if err != nil {
			return
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

func BenchmarkFibonacciWithoutIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciWithoutIterative(80)
	}
}

func BenchmarkFibonacciForVeryVeryBigNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciForVeryVeryBigNumber(80)
	}
}
