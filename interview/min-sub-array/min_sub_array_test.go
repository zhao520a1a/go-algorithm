package min_sub_array

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		s   int
		num []int
	}
	tests := []struct {
		name       string
		args       args
		wantMinLen int
	}{
		{
			"1",
			args{7, []int{2, 3, 1, 2, 4, 3}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMinLen := minSubArrayLen(tt.args.s, tt.args.num); gotMinLen != tt.wantMinLen {
				t.Errorf("minSubArrayLen() = %v, want %v", gotMinLen, tt.wantMinLen)
			}

			if gotMinLen := minSubArrayLenBinary(tt.args.s, tt.args.num); gotMinLen != tt.wantMinLen {
				t.Errorf("minSubArrayLenBinary() = %v, want %v", gotMinLen, tt.wantMinLen)
			}
		})
	}
}
