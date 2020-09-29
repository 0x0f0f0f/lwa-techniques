package lin

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// compute the orthogonal complement of a vector subspace of R^n
func Complement(n int, spanset []*mat.VecDense) []*mat.VecDense {
	dim := len(spanset)
	// the orthogonal complement of {} is R^n
	if dim == 0 {
		return UnitaryVecs(n, n)
	}

	a := mat.NewDense(len(spanset), n, nil)

	for i, vec := range spanset {
		a.SetRow(i, vec.RawVector().Data)
	}

	fmt.Printf("%.5g\n", mat.Formatted(a, mat.Squeeze()))

	basis, _ := NullspaceSVD(a)

	return basis
}
