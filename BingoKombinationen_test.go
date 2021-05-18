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
		{"zero",args{0},0},{"oneone",args{2},1}, {"multiple",args{12},2 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := countOnes(tt.args.n); gotI != tt.wantI {
				t.Errorf("countOnes() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}
