package lin

import (
	"log"
	"math"
	"testing"
)

var tc = testCase{
	a: [][]float64{
		{1.00, 0.00, 0.00, 0.00, 0.00, 0.00},
		{1.00, 0.63, 0.39, 0.25, 0.16, 0.10},
		{1.00, 1.26, 1.58, 1.98, 2.49, 3.13},
		{1.00, 1.88, 3.55, 6.70, 12.62, 23.80},
		{1.00, 2.51, 6.32, 15.88, 39.90, 100.28},
		{1.00, 3.14, 9.87, 31.01, 97.41, 306.02}},
	b: []float64{-0.01, 0.61, 0.91, 0.99, 0.60, 0.02},
	x: []float64{-0.01, 1.602790394502114, -1.6132030599055613,
		1.2454941213714368, -0.4909897195846576, 0.065760696175232},
}

func TestGaussPartial(t *testing.T) {
	x, err := GaussPartialo(tc.a, tc.b)
	if err != nil {
		log.Fatal(err)
	}
	for i, xi := range x {
		if math.Abs(tc.x[i]-xi) > ε {
			log.Println("out of tolerance")
			t.Errorf("expected %.20g", tc.x)
		}
	}
}
