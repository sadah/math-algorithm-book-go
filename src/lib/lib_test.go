package lib

import (
	"reflect"
	"testing"
)

func Test_isPrime(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"OK: 1", args{1}, true},
		{"OK: 2", args{2}, true},
		{"OK: 10", args{10}, false},
		{"OK: 13", args{13}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.n); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divisors(t *testing.T) {
	type args struct {
		N int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"OK: 1", args{1}, []int64{1}},
		{"OK: 2", args{2}, []int64{1, 2}},
		{"OK: 10", args{10}, []int64{1, 2, 5, 10}},
		{"OK: 13", args{13}, []int64{1, 13}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisors(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_primeFactorize(t *testing.T) {
	type args struct {
		N int64
	}
	tests := []struct {
		name string
		args args
		want map[int64]int64
	}{
		{"OK: 1", args{1}, map[int64]int64{}},
		{"OK: 2", args{2}, map[int64]int64{2: 1}},
		{"OK: 10", args{10}, map[int64]int64{2: 1, 5: 1}},
		{"OK: 13", args{13}, map[int64]int64{13: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := primeFactorize(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("primeFactorize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"OK: 1, 1", args{1, 1}, 1},
		{"OK: 10, 5", args{10, 5}, 5},
		{"OK: 3, 9", args{3, 9}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lcm(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"OK: 1, 1", args{1, 1}, 1},
		{"OK: 2, 5", args{2, 5}, 10},
		{"OK: 3, 9", args{3, 9}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcm(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combination(t *testing.T) {
	type args struct {
		n int64
		r int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"OK: 1, 1", args{1, 1}, 1},
		{"OK: 3, 2", args{3, 2}, 3},
		{"OK: 5, 3", args{5, 3}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combination(tt.args.n, tt.args.r); got != tt.want {
				t.Errorf("combination() = %v, want %v", got, tt.want)
			}
		})
	}
}
