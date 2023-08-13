package sinclairgo

import (
	"fmt"
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
				ACoefficient: AMale2017,
				BCoefficient: BMale2017,
			}},
			expectedSinclair: 285.66986,
		},
		{
			name: "Over500SinclairMale2009",
			args: args{bodyweight: 160, liftedTotal: 510, coefficients: Coefficients{
				ACoefficient: AMale2009,
				BCoefficient: BMale2009,
			}},
			expectedSinclair: 511.21796,
		},
		{
			name: "NormalSinclairFemale2009",
			args: args{bodyweight: 81, liftedTotal: 235, coefficients: Coefficients{
				ACoefficient: AFemale2009,
				BCoefficient: BFemale2009,
			}},
			expectedSinclair: 256.5641,
		},
		{
			name: "Over500SinclairFemale2017",
			args: args{bodyweight: 160, liftedTotal: 510, coefficients: Coefficients{
				ACoefficient: AFemale2017,
				BCoefficient: BFemale2017,
			}},
			expectedSinclair: 510,
		},
		{
			name: "SuperHeavySinclairMale2013",
			args: args{bodyweight: 200, liftedTotal: 400, coefficients: Coefficients{
				ACoefficient: AMale2013,
				BCoefficient: BMale2013,
			}},
			expectedSinclair: 400,
		},
		{
			name: "SuperHeavySinclairFemale2013",
			args: args{bodyweight: 200, liftedTotal: 400, coefficients: Coefficients{
				ACoefficient: AFemale2013,
				BCoefficient: BFemale2013,
			}},
			expectedSinclair: 400,
		},
		{
			name: "NormalSinclairMale2021",
			args: args{bodyweight: 81, liftedTotal: 235, coefficients: Coefficients{
				ACoefficient: AMale2021,
				BCoefficient: BMale2021,
			}},
			expectedSinclair: 298.24963,
		},
		{
			name: "NormalSinclairFemale2021",
			args: args{bodyweight: 81, liftedTotal: 235, coefficients: Coefficients{
				ACoefficient: AFemale2021,
				BCoefficient: BFemale2021,
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

func ExampleCalcSinclair() {
	var coeffs = Coefficients{ACoefficient: AMale2021, BCoefficient: BMale2021}
	var sinclairScore = CalcSinclair(100.0, 200.0, coeffs)
	fmt.Println(int(sinclairScore + 0.5))
	// Output: 229
}
