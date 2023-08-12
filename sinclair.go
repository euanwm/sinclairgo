package sinclairgo

import (
	"golang.org/x/exp/constraints"
	"math"
)

// Coefficients is a struct that holds the coefficients for the sinclair formula.
type Coefficients struct {
	aCoefficient float64
	bCoefficient float64
}

// CalcSinclair Calculates the sinclair of a result passed to it. We are using ONLY the Senior coefficient because
// the Masters coefficient is absolute nonsense. You'll see there's a lot of switching between float types.
// It's frustrating but it serves a purpose.
func CalcSinclair[T constraints.Float](bodyweight T, liftedTotal T, coeffs Coefficients) (sinclairScore float64) {
	if float64(bodyweight) <= coeffs.bCoefficient {
		var X = math.Log10(float64(bodyweight) / coeffs.bCoefficient)
		var expX = math.Pow(X, 2)
		var coEffExp = coeffs.aCoefficient * expX
		var expSum = math.Pow(10, coEffExp)
		sinclairScore = float64(liftedTotal) * expSum
	} else {
		sinclairScore = float64(liftedTotal)
	}
	return sinclairScore
}
