package sinclairgo

import (
	"golang.org/x/exp/constraints"
	"math"
)

// Coefficients is a struct that holds the coefficients for the sinclair formula.
type Coefficients struct {
	ACoefficient float64
	BCoefficient float64
}

// CalcSinclair is a function that calculates the sinclair score for a given bodyweight, lifted total, and coefficients.
// We recommend using the provided constants for the coefficients within the coeffyears.go file.
func CalcSinclair[T constraints.Float](bodyweight T, liftedTotal T, coeffs Coefficients) (sinclairScore float64) {
	if float64(bodyweight) <= coeffs.BCoefficient {
		var X = math.Log10(float64(bodyweight) / coeffs.BCoefficient)
		var expX = math.Pow(X, 2)
		var coEffExp = coeffs.ACoefficient * expX
		var expSum = math.Pow(10, coEffExp)
		sinclairScore = float64(liftedTotal) * expSum
	} else {
		sinclairScore = float64(liftedTotal)
	}
	return sinclairScore
}
