package main

import "testing"

func Test_countOnes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		wantI int
	}{
		{"zero", args{0}, 0}, {"oneone", args{2}, 1}, {"multiple", args{12}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := countOnes(tt.args.n); gotI != tt.wantI {
				t.Errorf("countOnes() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func Test_isLine(t *testing.T) {
	type args struct {
		n        int
		gridsize int
	}
	var tests = []struct {
		name string
		args args
		want bool
	}{
		{"1emty", args{0, 1}, false},
		{"1full", args{1, 1}, true},
		{"2diag", args{0b1001, 2}, false},
		{"4line", args{0b1111 << 4, 4}, true},
		{"4nonsense", args{0b1111 << 3, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLine(tt.args.n, tt.args.gridsize); got != tt.want {
				t.Errorf("isLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isRow(t *testing.T) {
	type args struct {
		n        int
		gridsize int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1emty", args{0, 1}, false},
		{"1full", args{1, 1}, true},
		{"2diag", args{0b1001, 2}, false},
		{"2col", args{0b1010, 2}, true},
		{"2colshifted", args{0b0101, 2}, true},
		{"2row", args{0b1100, 2}, false},
		{"4line", args{0b1111 << 4, 4}, false},
		{"4nonsense", args{0b1111 << 3, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRow(tt.args.n, tt.args.gridsize); got != tt.want {
				t.Errorf("isRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDiagonal(t *testing.T) {
	type args struct {
		n        int
		gridsize int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1emty", args{0, 1}, false},
		{"1full", args{1, 1}, true},
		{"2diag", args{0b1001, 2}, true},
		{"2diagother", args{0b0110, 2}, true},
		{"2col", args{0b1010, 2}, false},
		{"2colshifted", args{0b0101, 2}, false},
		{"2row", args{0b1100, 2}, false},
		{"3diag", args{0b100010001, 3}, true},
		{"3diag", args{0b001010100, 3}, true},
		{"4line", args{0b1111 << 4, 4}, false},
		{"4nonsense", args{0b1111 << 3, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDiagonal(tt.args.n, tt.args.gridsize); got != tt.want {
				t.Errorf("isDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
