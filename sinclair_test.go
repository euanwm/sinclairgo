package sinclairgo

import (
	"testing"
)

func TestCalcSinclair(t *testing.T) {
	type args struct {
		bodyweight   float64
		liftedTotal  float64
		coefficients Coefficients
	}
	tests := []struct {
		name             string
		args             args
		expectedSinclair float64
	}{
		{
			name: "NormalSinclairMale2017",
			args: args{bodyweight: 81, liftedTotal: 235, coefficients: Coefficients{
				aCoefficient: AMale2017,
				bCoefficient: BMale2017,
			}},
			expectedSinclair: 285.66986,
		},
		{
			name: "Over500SinclairMale2009",
			args: args{bodyweight: 160, liftedTotal: 510, coefficients: Coefficients{
				aCoefficient: AMale2009,
				bCoefficient: BMale2009,
			}},
			expectedSinclair: 511.21796,
		},
		{
			name: "NormalSinclairFemale2009",
			args: args{bodyweight: 81, liftedTotal: 235, coefficients: Coefficients{
				aCoefficient: AFemale2009,
				bCoefficient: BFemale2009,
			}},
			expectedSinclair: 256.5641,
		},
		{
			name: "Over500SinclairFemale2017",
			args: args{bodyweight: 160, liftedTotal: 510, coefficients: Coefficients{
				aCoefficient: AFemale2017,
				bCoefficient: BFemale2017,
			}},
			expectedSinclair: 510,
		},
		{
			name: "SuperHeavySinclairMale2013",
			args: args{bodyweight: 200, liftedTotal: 400, coefficients: Coefficients{
				aCoefficient: AMale2013,
				bCoefficient: BMale2013,
			}},
			expectedSinclair: 400,
		},
		{
			name: "SuperHeavySinclairFemale2013",
			args: args{bodyweight: 200, liftedTotal: 400, coefficients: Coefficients{
				aCoefficient: AFemale2013,
				bCoefficient: BFemale2013,
			}},
			expectedSinclair: 400,
		},
		{
			name: "NormalSinclairMale2021",
			args: args{bodyweight: 81, liftedTotal: 235, coefficients: Coefficients{
				aCoefficient: AMale2021,
				bCoefficient: BMale2021,
			}},
			expectedSinclair: 298.24963,
		},
		{
			name: "NormalSinclairFemale2021",
			args: args{bodyweight: 81, liftedTotal: 235, coefficients: Coefficients{
				aCoefficient: AFemale2021,
				bCoefficient: BFemale2021,
			}},
			expectedSinclair: 270.42316,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := float32(CalcSinclair(tt.args.bodyweight, tt.args.liftedTotal, tt.args.coefficients)); got != float32(tt.expectedSinclair) {
				t.Errorf("CalcSinclair() = %v, want %v", got, tt.expectedSinclair)
			}
		})
	}
}
