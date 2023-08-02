package sinclairgo

import (
	"golang.org/x/exp/constraints"
	"math"
)

// Coefficient numbers
const (
	aMale   = 0.751945030
	bMale   = 175.508
	aFemale = 0.783497476
	bFemale = 153.655
)

// CalcSinclair Calculates the sinclair of a result passed to it. We are using ONLY the Senior coefficient because
// the Masters coefficient is absolute nonsense. You'll see there's a lot of switching between float types.
// It's frustrating but it serves a purpose.
func CalcSinclair[T constraints.Float](bodyweight T, liftedTotal T, isFemale bool) (sinclairScore float64) {
	var coEffA = aMale
	var coEffB = bMale
	if isFemale {
		coEffA = aFemale
		coEffB = bFemale
	}
	if float64(bodyweight) <= coEffB {
		var X = math.Log10(float64(bodyweight) / coEffB)
		var expX = math.Pow(X, 2)
		var coEffExp = coEffA * expX
		var expSum = math.Pow(10, coEffExp)
		sinclairScore = float64(liftedTotal) * expSum
	} else {
		sinclairScore = float64(liftedTotal)
	}
	return sinclairScore
}
