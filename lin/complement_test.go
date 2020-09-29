package lin

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestComplement(t *testing.T) {
	w1 := mat.NewVecDense(3, []float64{1, 7, 2})
	w2 := mat.NewVecDense(3, []float64{-2, 3, 1})
	//assert.Equal()

	wperp := Complement(3, []*mat.VecDense{w1, w2})

	for _, vec := range wperp {
		fmt.Printf("%.20g\n", mat.Formatted(vec, mat.Squeeze()))
	}

}
