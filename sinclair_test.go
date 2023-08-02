package sinclairgo

import (
	"testing"
)

func TestCalcSinclair(t *testing.T) {
	type args struct {
		bodyweight  float64
		liftedTotal float64
		isFemale    bool
	}
	tests := []struct {
		name             string
		args             args
		expectedSinclair float64
	}{
		{
			name:             "NormalSinclairMale",
			args:             args{bodyweight: 81, liftedTotal: 235, isFemale: false},
			expectedSinclair: 285.66986,
		},
		{
			name:             "Over500SinclairMale",
			args:             args{bodyweight: 160, liftedTotal: 510, isFemale: false},
			expectedSinclair: 511.42737,
		},
		{
			name:             "NormalSinclairFemale",
			args:             args{bodyweight: 81, liftedTotal: 235, isFemale: true},
			expectedSinclair: 270.17587,
		},
		{
			name:             "Over500SinclairFemale",
			args:             args{bodyweight: 160, liftedTotal: 510, isFemale: true},
			expectedSinclair: 510,
		},
		{
			name:             "SuperHeavySinclairMale",
			args:             args{bodyweight: 200, liftedTotal: 400, isFemale: false},
			expectedSinclair: 400,
		},
		{
			name:             "SuperHeavySinclairFemale",
			args:             args{bodyweight: 200, liftedTotal: 400, isFemale: true},
			expectedSinclair: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := float32(CalcSinclair(tt.args.bodyweight, tt.args.liftedTotal, tt.args.isFemale)); got != float32(tt.expectedSinclair) {
				t.Errorf("CalcSinclair() = %v, want %v", got, tt.expectedSinclair)
			}
		})
	}
}
